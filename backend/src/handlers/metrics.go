package handlers

import (
	"net/http"
	"shipa-gen/src/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type MetricsResponse struct {
	Metrics map[string]models.Metric `json:"metrics"`
}

func (h *handler) GetMetricsResourceConfig(c *gin.Context) {
	id := c.Param("id")

	metricTypes := c.Query("type")
	var requestMetrics []string
	if len(metricTypes) > 0 {
		requestMetrics = strings.Split(metricTypes, ",")
	} else {
		requestMetrics = []string{models.EventTypeClone}
	}
	userInfo, err := h.GetCurrentUserInfo(c)

	metrics, err := h.srv.CountEventsMetrics(id, requestMetrics, userInfo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	response := MetricsResponse{
		Metrics: map[string]models.Metric{},
	}
	for _, metric := range metrics {
		response.Metrics[metric.Name] = metric
	}
	for _, metricName := range requestMetrics {
		if _, ok := response.Metrics[metricName]; !ok {
			response.Metrics[metricName] = models.Metric{Name: metricName}
		}
	}

	c.IndentedJSON(http.StatusOK, response)
}
