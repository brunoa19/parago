package handlers

import (
	"net/http"
	"shipa-gen/src/client/shipa_client"
	"shipa-gen/src/storage"
	"strconv"
	"time"

	"shipa-gen/src/models"
	"shipa-gen/src/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	srv   service.Service
	auth  service.AuthorisationService
	shipa shipa_client.ShipaClient
}

func New(service service.Service, authService service.AuthorisationService, shipaClient shipa_client.ShipaClient) *handler {
	return &handler{
		srv:   service,
		auth:  authService,
		shipa: shipaClient,
	}
}

func (h *handler) GetCurrentUserInfoWithExtToken(c *gin.Context) (*models.UserInfo, string, error) {
	token := c.GetHeader("x-auth-token")
	userInfo, extToken, err := h.auth.AuthGetUser(token)
	if err != nil {
		return nil, extToken, err
	}
	if userInfo != nil {
		return userInfo, extToken, nil
	}

	return &models.UserInfo{
		Token: token,
	}, extToken, nil
}

func (h *handler) GetCurrentUserInfoAuthorizedWithExtToken(c *gin.Context) (*models.UserInfo, string, error) {
	userInfo, token, err := h.GetCurrentUserInfoWithExtToken(c)
	if err != nil {
		return nil, token, err
	}

	if userInfo.IsGuest() {
		return nil, token, ErrorUnauthorizedRequest
	}

	return userInfo, token, nil
}

func (h *handler) GetCurrentUserInfo(c *gin.Context) (*models.UserInfo, error) {
	userInfo, _, err := h.GetCurrentUserInfoWithExtToken(c)

	return userInfo, err
}

func (h *handler) GetCurrentUserInfoAuthorized(c *gin.Context) (*models.UserInfo, error) {
	userInfo, _, err := h.GetCurrentUserInfoAuthorizedWithExtToken(c)

	return userInfo, err
}

func (h *handler) GetHttpResponseCodeByError(err error) int {
	if err == ErrorUnauthorizedRequest {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func (h *handler) GetPaginator(c *gin.Context) models.Paginator {
	paginator := models.Paginator{}
	paginator.Page, _ = strconv.Atoi(c.Query("page"))
	paginator.PageSize, _ = strconv.Atoi(c.Query("pageSize"))

	return paginator
}

func (h *handler) RootUrlHandler(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}

func (h *handler) CreateResourceConfig(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfoAuthorized(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	var cfg models.CreateResourceConfigRequest

	if err := c.BindJSON(&cfg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	cfg.Init()

	if err := cfg.Validate(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	config, err := h.srv.Create(&cfg, userInfo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

type SearchResponse struct {
	Results []*models.Metadata `json:"results,omitempty"`
}

func (h *handler) SearchResourceConfig(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfo(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	search := c.Query("q")

	configs, err := h.srv.Search(search, userInfo)
	if err != nil {
		httpCode := http.StatusInternalServerError
		if err.Error() == storage.ErrNotFound.Error() {
			httpCode = http.StatusNotFound
		}
		c.IndentedJSON(httpCode, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, &SearchResponse{
		Results: configs,
	})
}

func (h *handler) GetResourceConfig(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfo(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	id := c.Param("id")

	config, err := h.srv.Get(id, userInfo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

func (h *handler) ListResourceConfigs(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfo(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	paginator := h.GetPaginator(c)
	configs, err := h.srv.List(userInfo, "personal", paginator)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		if err.Error() == storage.ErrNotFound.Error() {
			httpStatus = http.StatusNotFound
		}
		c.IndentedJSON(httpStatus, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, &SearchResponse{
		Results: configs,
	})
}

func (h *handler) ListResourceByTypeConfigs(c *gin.Context) {
	accessLevel := c.Param("accessLevel")

	user, err := h.GetCurrentUserInfo(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	paginator := h.GetPaginator(c)
	configs, err := h.srv.List(user, models.AccessLevelType(accessLevel), paginator)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		if err.Error() == storage.ErrNotFound.Error() {
			httpStatus = http.StatusNotFound
		}
		c.IndentedJSON(httpStatus, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, &SearchResponse{
		Results: configs,
	})
}

func (h *handler) CloneResourceConfig(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfoAuthorized(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	id := c.Param("id")

	config, err := h.srv.Clone(id, userInfo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

func (h *handler) UpdateResourceConfig(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfoAuthorized(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	id := c.Param("id")

	var cfg models.UpdateResourceConfigRequest

	if err := c.BindJSON(&cfg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	cfg.Init()

	if err := cfg.Validate(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	config, err := h.srv.Update(id, &cfg, userInfo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

func (h *handler) DeleteResourceConfig(c *gin.Context) {
	userInfo, err := h.GetCurrentUserInfoAuthorized(c)
	if err != nil {
		c.IndentedJSON(h.GetHttpResponseCodeByError(err), JsonError{err.Error()})
		return
	}

	id := c.Param("id")

	if err = h.srv.Delete(id, userInfo); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *handler) GetStatistics(c *gin.Context) {
	from, err := parseTime(c.Query("from"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	to, err := parseTime(c.Query("to"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, JsonError{err.Error()})
		return
	}

	filter := &models.EventFilter{
		ConfigID: c.Param("id"),
		Type:     models.EventType(c.Query("type")),
		From:     from,
		To:       to,
	}

	count, err := h.srv.CountEvents(filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, JsonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, &EventsCount{
		Count: count,
	})
}

type EventsCount struct {
	Count int64 `json:"count"`
}

func parseTime(val string) (*time.Time, error) {
	if val == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339Nano, val)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
