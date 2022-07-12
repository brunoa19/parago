package builder

import (
	"log"
	"shipa-gen/src/client/mongo_client"
	"shipa-gen/src/configuration"
	"shipa-gen/src/service/authorisation_service"
	"shipa-gen/src/storage/auth_mem_storage"
	"shipa-gen/src/storage/auth_mongo_storage"
)

func NewAuthStorage(client mongo_client.MongoClient, conf *configuration.Configuration) authorisation_service.AuthorizationStorage {
	if conf.AuthStorage == authorisation_service.AuthStorageMongo {
		log.Println("auth storage type is mongo")
		return auth_mongo_storage.NewAuthMongoStorage(client.GetClient(), conf.AuthMongoDb, conf.AuthMongoUserCollection)
	}
	log.Println("auth storage type is memory")

	return auth_mem_storage.NewAuthMemStorage()
}
