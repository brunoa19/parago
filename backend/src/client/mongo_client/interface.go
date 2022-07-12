package mongo_client

import "go.mongodb.org/mongo-driver/mongo"

type MongoClient interface {
	GetClient() *mongo.Client
	Disconnect() error
}
