package auth_mongo_storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"shipa-gen/src/service/authorisation_service"
)

type AuthMongoStorage struct {
	usersCollection *mongo.Collection
}

type AuthUserInfo struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	SessionToken string             `bson:"sessionToken"`
	authorisation_service.AuthUserInfo
}

func NewAuthMongoStorage(mongoClient *mongo.Client, dbName string, collectionName string) authorisation_service.AuthorizationStorage {
	database := mongoClient.Database(dbName)

	return &AuthMongoStorage{
		usersCollection: database.Collection(collectionName),
	}
}

func (s *AuthMongoStorage) Store(sessionToken string, userInfo authorisation_service.AuthUserInfo) error {
	ctx := context.TODO()
	storedUserInfo, err := s.getUserInfo(sessionToken, ctx)
	if err != nil {
		return err
	}

	storeUserInfo := AuthUserInfo{
		Id:           primitive.ObjectID{},
		SessionToken: sessionToken,
		AuthUserInfo: userInfo,
	}

	if storedUserInfo != nil {
		storeUserInfo.Id = storedUserInfo.Id
		_, err := s.usersCollection.UpdateByID(ctx, storedUserInfo.Id, storeUserInfo)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = s.usersCollection.InsertOne(ctx, storeUserInfo)

	return err
}

func (s *AuthMongoStorage) Load(sessionToken string) (*authorisation_service.AuthUserInfo, error) {
	ctx := context.TODO()
	userInfo, err := s.getUserInfo(sessionToken, ctx)
	if err != nil {
		return nil, err
	}

	if userInfo == nil {
		return nil, nil
	}

	return &userInfo.AuthUserInfo, nil
}

func (s *AuthMongoStorage) Remove(sessionToken string) error {
	ctx := context.TODO()
	storedUserInfo, err := s.getUserInfo(sessionToken, ctx)
	if err != nil {
		return err
	}
	if storedUserInfo == nil {
		return nil
	}

	_, err = s.usersCollection.DeleteOne(ctx, bson.D{{"_id", storedUserInfo.Id}})

	return err
}

func (s *AuthMongoStorage) getUserInfo(sessionToken string, ctx context.Context) (*AuthUserInfo, error) {
	filter := bson.D{{"sessionToken", sessionToken}}

	cursor := s.usersCollection.FindOne(ctx, filter)
	if err := cursor.Err(); err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			return nil, err
		}
		return nil, nil
	}

	var userInfo *AuthUserInfo
	if err := cursor.Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
