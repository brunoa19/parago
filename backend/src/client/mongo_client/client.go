package mongo_client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"shipa-gen/src/configuration"
	"time"
)

type MongoClientImpl struct {
	mongoClient     *mongo.Client
	mongoContext    context.Context
	mongoCancelFunc context.CancelFunc
}

func NewClient(appConfig *configuration.Configuration) (MongoClient, error) {
	dbConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(appConfig.MongoUrl))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	if err := dbConn.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("failed to ping mongo db: ", err)
	}

	return &MongoClientImpl{
		mongoClient:     dbConn,
		mongoContext:    ctx,
		mongoCancelFunc: cancel,
	}, nil
}

func (c *MongoClientImpl) GetClient() *mongo.Client {
	return c.mongoClient
}

func (c *MongoClientImpl) Disconnect() error {
	c.mongoCancelFunc()

	return c.mongoClient.Disconnect(nil)
}
