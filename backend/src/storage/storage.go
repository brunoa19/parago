package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"shipa-gen/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = errors.New("record not found")
)

const (
	dbName            = "shipa-gen"
	configsCollection = "configs"
	eventsCollection  = "events"
)

type service struct {
	configs *mongo.Collection
	events  *mongo.Collection
}

func New(dbConn *mongo.Client) (*service, error) {
	database := dbConn.Database(dbName)

	return &service{
		configs: database.Collection(configsCollection),
		events:  database.Collection(eventsCollection),
	}, nil
}

func (s *service) FindByName(name string, userInfo *models.UserInfo) ([]*models.ResourceConfig, error) {
	ctx := context.TODO()

	var configs []*models.ResourceConfig
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"deleted", bson.D{{"$ne", true}}}},
				bson.D{{"name", bson.D{{"$eq", name}}}},
				bson.D{{"org", bson.D{{"$eq", userInfo.Org}}}},
				bson.D{{"user", bson.D{{"$eq", userInfo.User}}}},
			},
		},
	}

	cursor, err := s.configs.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &configs); err != nil {
		return nil, err
	}

	if configs == nil || len(configs) == 0 {
		return nil, ErrNotFound
	}

	return configs, nil
}

func (s *service) Search(search string, user *models.UserInfo) ([]*models.ResourceConfig, error) {
	ctx := context.TODO()

	var configs []*models.ResourceConfig
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"deleted", bson.D{{"$ne", true}}}},
				bson.D{{"name", bson.D{{"$regex", fmt.Sprintf(".*%s.*", search)}, {"$options", "si"}}}},
			},
		},
	}

	cursor, err := s.configs.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &configs); err != nil {
		return nil, err
	}

	var results []*models.ResourceConfig
	for _, config := range configs {
		if config.IsAvailableForUser(user) {
			results = append(results, config)
		}
	}

	if results == nil {
		return nil, ErrNotFound
	}

	return results, nil
}

func (s *service) List(userInfo *models.UserInfo, accessLevel models.AccessLevelType, paginator models.Paginator) ([]*models.ResourceConfig, error) {
	ctx := context.TODO()
	var configs []*models.ResourceConfig

	conditions := bson.A{
		bson.D{{"deleted", bson.D{{"$ne", true}}}},
	}
	accessLevelSupported := false
	// public access level condition
	if accessLevel == models.AccessLevelPublic {
		conditions = append(
			conditions,
			bson.D{{"accessLevel", bson.D{{"$eq", models.AccessLevelPublic}}}},
		)
		accessLevelSupported = true
	}
	// organisation access level condition
	if accessLevel == models.AccessLevelOrganization {
		conditions = append(
			conditions,
			bson.D{{"accessLevel", bson.D{{"$eq", models.AccessLevelOrganization}}}},
			bson.D{{"org", bson.D{{"$eq", userInfo.Org}}}},
		)
		accessLevelSupported = true
	}
	// private access level condition
	if accessLevel == models.AccessLevelPrivate {
		conditions = append(
			conditions,
			bson.D{{"accessLevel", bson.D{{"$eq", models.AccessLevelPrivate}}}},
			bson.D{{"user", bson.D{{"$eq", userInfo.User}}}},
		)
		accessLevelSupported = true
	}
	// personal access level condition
	if accessLevel == models.AccessModePersonal {
		conditions = append(
			conditions,
			bson.D{{"user", bson.D{{"$eq", userInfo.User}}}},
		)
		accessLevelSupported = true
	}

	if !accessLevelSupported {
		return nil, errors.New("not supported: accessLevel")
	}

	filter := bson.D{{"$and", conditions}}

	findOptions := options.FindOptions{}
	if paginator.PageSize > 0 {
		findOptions.SetLimit(paginator.GetLimit())
		if paginator.Page > 0 {
			findOptions.SetSkip(paginator.GetOffset())
		}
	}

	cursor, err := s.configs.Find(ctx, filter, &findOptions)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &configs); err != nil {
		return nil, err
	}

	if configs == nil || len(configs) == 0 {
		return nil, ErrNotFound
	}

	return configs, nil
}

func (s *service) Create(config *models.ResourceConfig) (*models.ResourceConfig, error) {
	config.CreatedAt = time.Now().UTC()
	res, err := s.configs.InsertOne(context.TODO(), config)
	if err != nil {
		return nil, err
	}

	config.ID = res.InsertedID.(primitive.ObjectID)

	return config, nil
}

func (s *service) Get(id string, userInfo *models.UserInfo) (*models.ResourceConfig, error) {
	ctx := context.TODO()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", objectId}},
				bson.D{{"deleted", bson.D{{"$ne", true}}}},
			},
		},
	}

	cursor := s.configs.FindOne(ctx, filter)
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	var config *models.ResourceConfig
	if err := cursor.Decode(&config); err != nil {
		return nil, err
	}

	// TODO: move business logic to service layer
	if !config.IsAvailableForUser(userInfo) {
		return nil, ErrNotFound
	}

	return config, nil
}

func (s *service) Clone(id string, userInfo *models.UserInfo) (*models.ResourceConfig, error) {
	// TODO: change input params: config, userInfo

	clone, err := s.Get(id, userInfo)
	if err != nil {
		return nil, err
	}

	clone.ID = primitive.NewObjectID()
	clone.UpdatedAt = nil
	clone.OwnerUserInfo = &models.OwnerUserInfo{
		UserId:   userInfo.UserId,
		UserName: userInfo.User,
		UserOrg:  userInfo.Org,
	}

	for {
		exists, _ := s.FindByName(clone.Name, userInfo)
		if exists == nil {
			break
		}
		clone.Name += " copy"
	}

	return s.Create(clone)
}

func (s *service) Update(config *models.ResourceConfig) (*models.ResourceConfig, error) {
	ctx := context.TODO()

	t := time.Now().UTC()
	config.UpdatedAt = &t

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", config.ID}},
				bson.D{{"deleted", bson.D{{"$ne", true}}}},
			},
		},
	}

	cursor := s.configs.FindOneAndReplace(ctx, filter, config)
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func (s *service) Delete(id string) error {
	ctx := context.TODO()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", objectId}},
				bson.D{{"deleted", bson.D{{"$ne", true}}}},
			},
		},
	}

	update := bson.D{
		{"$set",
			bson.D{
				{"deleted", true},
				{"updatedAt", time.Now().UTC()},
			},
		},
	}

	cursor := s.configs.FindOneAndUpdate(ctx, filter, update)
	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}

func (s *service) AddEvent(event *models.Event) error {
	event.CreatedAt = time.Now().UTC()
	_, err := s.events.InsertOne(context.TODO(), event)
	return err
}

func (s *service) CountEvents(filter *models.EventFilter) (int64, error) {
	filterOptions := bson.A{
		bson.D{{"configId", bson.D{{"$eq", filter.ConfigID}}}},
		bson.D{{"type", bson.D{{"$eq", filter.Type}}}},
	}

	if filter.From != nil || filter.To != nil {
		filterTimeOptions := bson.D{}

		if filter.From != nil {
			filterTimeOptions = append(filterTimeOptions, bson.E{"$gte", filter.From})
		}

		if filter.To != nil {
			filterTimeOptions = append(filterTimeOptions, bson.E{"$lte", filter.To})
		}

		filterOptions = append(filterOptions, bson.D{{"createdAt", filterTimeOptions}})
	}

	filterDb := bson.D{
		{"$and", filterOptions},
	}

	count, err := s.events.CountDocuments(context.TODO(), filterDb)
	if err != nil {
		return 0, err
	}

	return count, nil
}
