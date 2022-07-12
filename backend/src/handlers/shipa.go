package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) ShipaGetRequest(c *gin.Context) {
	_, extToken, err := h.GetCurrentUserInfoAuthorizedWithExtToken(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	path := c.Param("path")

	status, contentType, body, err := h.shipa.GetResponse(extToken, path)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	c.Data(status, contentType, body)
}
