package configuration

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Configuration struct {
	BackendPort string `env:"SHIPA_GEN_BACKEND_PORT" envDefault:"8080"`
	MongoUrl    string `env:"SHIPA_GEN_MONGO_URI" envDefault:"mongodb://localhost:27017"`
	ApplicationMode string `json:"APPLICATION_MODE" envDefault:"release"`
	ShipaServerBaseUrl      string `env:"SHIPA_SERVER_BASE_URL" envDefault:"https://target.shipa.cloud:443"`
	AuthStorage             string `env:"SHIPA_GEN_AUTH_STORAGE" envDefault:"mongo"`
	AuthMongoDb             string `env:"SHIPA_GEN_AUTH_MONGO_DB" envDefault:"shipa-gen"`
	AuthMongoUserCollection string `env:"SHIPA_GEN_AUTH_MONGO_USER_COLLECTION" envDefault:"users"`
}

func NewConfiguration() *Configuration {

	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
