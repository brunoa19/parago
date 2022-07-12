package main

import (
	"log"
	"shipa-gen/src/builder"
	"shipa-gen/src/client/mongo_client"
	"shipa-gen/src/client/shipa_client"
	"shipa-gen/src/configuration"
	"shipa-gen/src/service/authorisation_service"
	"time"

	"shipa-gen/src/handlers"
	"shipa-gen/src/service"
	"shipa-gen/src/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := configuration.NewConfiguration()

	//chartOut := os.TempDir() + "/charts"
	//helmGenerator := helm.NewChartGenerator("helm/template", chartOut)
	//cfg := shipa.Config{
	//	AppName: "test_app",
	//}
	//path, err := helmGenerator.PrepareChart(&cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(path)
	//log.Fatal("Stop")

	mongoClient, err := mongo_client.NewClient(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = mongoClient.Disconnect()
	}()

	log.Println("Successfully connected to mongo db")

	gin.SetMode(gin.ReleaseMode) // gin.ReleaseMode

	store, err := storage.New(mongoClient.GetClient())
	if err != nil {
		log.Fatal(err)
	}

	shipaClient := shipa_client.NewClient(appConfig.ShipaServerBaseUrl)
	authStorage := builder.NewAuthStorage(mongoClient, appConfig)
	authorisationService := authorisation_service.NewService(shipaClient, authStorage, 5*time.Minute)

	srv := service.New(store)
	handler := handlers.New(srv, authorisationService, shipaClient)

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/", handler.RootUrlHandler)

	router.Group("/shipa-gen").
		POST("apps", handlers.GenerateAppsHandler).
		POST("volumes", handlers.GenerateVolumesHandler).
		POST("frameworks", handlers.GenerateFrameworksHandler)

	router.Group("/shipa-server").
		GET("/*path", handler.ShipaGetRequest)

	router.Group("/config").
		POST("", handler.CreateResourceConfig).
		GET("", handler.ListResourceConfigs).
		GET("list/:accessLevel", handler.ListResourceByTypeConfigs).
		GET(":id", handler.GetResourceConfig).
		POST(":id", handler.UpdateResourceConfig).
		DELETE(":id", handler.DeleteResourceConfig).
		GET(":id/metrics", handler.GetMetricsResourceConfig).
		GET("search", handler.SearchResourceConfig).
		GET("clone/:id", handler.CloneResourceConfig)

	router.Group("/auth").
		POST("login", handler.AuthLoginHandler).
		POST("logout", handler.AuthLogoutHandler).
		GET("user", handler.AuthUserHandler)

	router.GET("/statistics/:id", handler.GetStatistics)

	log.Println("Server starts listen on port", appConfig.BackendPort)
	err = router.Run(":" + appConfig.BackendPort)
	if err != nil {
		log.Fatal(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Disposition")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
