package handlers

import (
	"net/http"

	"shipa-gen/src/service"
	"shipa-gen/src/shipa"

	"github.com/gin-gonic/gin"
)

type JsonError struct {
	Error string `json:"error"`
}

func GenerateAppsHandler(c *gin.Context) {
	var cfg shipa.AppsConfig

	if err := c.BindJSON(&cfg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	if len(cfg.Apps) == 0 {
		c.IndentedJSON(http.StatusBadRequest, JsonError{"empty input data"})
		return
	}

	if cfg.Provider == "" {
		c.IndentedJSON(http.StatusBadRequest, JsonError{"provider not set"})
		return
	}

	out, err := service.GenerateApps(cfg)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, out)
}

func GenerateVolumesHandler(c *gin.Context) {
	var cfg shipa.VolumesConfig

	if err := c.BindJSON(&cfg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	if len(cfg.Volumes) == 0 {
		c.IndentedJSON(http.StatusBadRequest, JsonError{"empty input data"})
		return
	}

	if cfg.Provider == "" {
		c.IndentedJSON(http.StatusBadRequest, JsonError{"provider not set"})
		return
	}

	out, err := service.GenerateVolumes(cfg)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, out)
}

func GenerateFrameworksHandler(c *gin.Context) {
	var cfg shipa.FrameworksConfig

	if err := c.BindJSON(&cfg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	if len(cfg.Frameworks) == 0 {
		c.IndentedJSON(http.StatusBadRequest, JsonError{"empty input data"})
		return
	}

	if cfg.Provider == "" {
		c.IndentedJSON(http.StatusBadRequest, JsonError{"provider not set"})
		return
	}

	out, err := service.GenerateFrameworks(cfg)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, out)
}
