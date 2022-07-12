package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shipa-gen/src/models"
)

func (h *handler) AuthLoginHandler(c *gin.Context) {

	loginRequest := models.AuthLoginRequest{}

	if err := c.BindJSON(&loginRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	userInfo, err := h.auth.AuthLogin(loginRequest.Login, loginRequest.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	if userInfo == nil {
		c.IndentedJSON(http.StatusForbidden, JsonError{"authorisation_service failed"})
		return
	}

	c.IndentedJSON(http.StatusOK, userInfo)
}

func (h *handler) AuthLogoutHandler(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfoAuthorized(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	err = h.auth.AuthLogout(userInfo.Token)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, models.UserInfo{})
}

func (h *handler) AuthUserHandler(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfoAuthorized(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, userInfo)
}
