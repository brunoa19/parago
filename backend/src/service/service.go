package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"shipa-gen/src/models"
	"shipa-gen/src/shipa"
	"shipa-gen/src/storage"

	"github.com/iancoleman/strcase"
)

var (
	ErrAccessForbidden = errors.New("access forbidden")
)

type Service interface {
	List(userInfo *models.UserInfo, accessLevel models.AccessLevelType, paginator models.Paginator) ([]*models.Metadata, error)
	Search(search string, userInfo *models.UserInfo) ([]*models.Metadata, error)
	Create(config *models.CreateResourceConfigRequest, userInfo *models.UserInfo) (*models.Metadata, error)
	Get(id string, userInfo *models.UserInfo) (*models.ResourceWithPayload, error)
	Clone(id string, userInfo *models.UserInfo) (*models.Metadata, error)
	Update(id string, config *models.UpdateResourceConfigRequest, userInfo *models.UserInfo) (*models.Metadata, error)
	Delete(id string, userInfo *models.UserInfo) error

	CountEvents(filter *models.EventFilter) (int64, error)
	CountEventsMetrics(id string, metrics []string, userInfo *models.UserInfo) ([]models.Metric, error)
}

type AuthorisationService interface {
	AuthLogin(login string, password string) (*models.UserInfo, error)
	AuthLogout(sessionToken string) error
	AuthGetUser(sessionToken string) (*models.UserInfo, string, error)
}

type service struct {
	storage storage.Storage
}

func New(storage storage.Storage) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) List(userInfo *models.UserInfo, accessLevel models.AccessLevelType, paginator models.Paginator) ([]*models.Metadata, error) {
	configs, err := s.storage.List(userInfo, accessLevel, paginator)
	if err != nil {
		return nil, err
	}

	var results []*models.Metadata
	for _, c := range configs {
		results = append(results, &c.Metadata)
	}

	return results, nil
}

func (s *service) Search(search string, userInfo *models.UserInfo) ([]*models.Metadata, error) {
	configs, err := s.storage.Search(search, userInfo)
	if err != nil {
		return nil, err
	}

	var results []*models.Metadata
	for _, c := range configs {
		results = append(results, &c.Metadata)
	}

	return results, nil
}

func (s *service) Create(req *models.CreateResourceConfigRequest, userInfo *models.UserInfo) (*models.Metadata, error) {
	existing, _ := s.storage.FindByName(req.Name, userInfo)
	if existing != nil {
		return nil, fmt.Errorf("config with name: %s - already exists", req.Name)
	}

	config := &models.ResourceConfig{
		Metadata: models.Metadata{
			OwnerUserInfo: &models.OwnerUserInfo{
				UserId:   userInfo.UserId,
				UserName: userInfo.User,
				UserOrg:  userInfo.Org,
			},
		},
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	result, err := s.storage.Create(config)
	if err != nil {
		return nil, err
	}

	return &result.Metadata, nil
}

func (s *service) Get(id string, userInfo *models.UserInfo) (*models.ResourceWithPayload, error) {
	config, err := s.storage.Get(id, userInfo)
	if err != nil {
		return nil, err
	}

	payload, err := generatePayload(config)
	if err != nil {
		return nil, err
	}

	err = s.storage.AddEvent(&models.Event{
		ConfigID: id,
		User:     userInfo.User,
		Org:      userInfo.Org,
		Type:     models.EventTypeLoad,
	})
	if err != nil {
		return nil, err
	}

	if payload != nil && payload.File != nil {
		payload.File.Name = generateFilename(config.OwnerUserInfo.UserOrg, config.OwnerUserInfo.UserName, config.Provider, config.Name, payload.File.Name)
	}

	return &models.ResourceWithPayload{
		ResourceConfig: *config,
		Payload:        payload,
	}, nil
}

func generateFilename(org, user, provider, configName, filename string) string {
	return fmt.Sprintf("%s-%s",
		strcase.ToKebab(strings.Join([]string{org, user, provider, configName}, " ")),
		filename,
	)
}

func generatePayload(config *models.ResourceConfig) (*models.Payload, error) {
	switch config.Type {
	case models.ResourceTypeApp:
		var cfg shipa.AppsConfig
		decodeData(config.Data, &cfg)
		return GenerateApps(cfg)
	case models.ResourceTypeFramework:
		var cfg shipa.FrameworksConfig
		decodeData(config.Data, &cfg)
		return GenerateFrameworks(cfg)
	case models.ResourceTypeVolume:
		var cfg shipa.VolumesConfig
		decodeData(config.Data, &cfg)
		return GenerateVolumes(cfg)
	default:
		return nil, fmt.Errorf("not supported config type: %s", config.Type)
	}
}

func decodeData(data map[string]interface{}, dst interface{}) {
	raw, _ := json.Marshal(data)
	_ = json.Unmarshal(raw, dst)
}

func (s *service) Clone(id string, userInfo *models.UserInfo) (*models.Metadata, error) {
	config, err := s.storage.Clone(id, userInfo)
	if err != nil {
		return nil, err
	}

	err = s.storage.AddEvent(&models.Event{
		ConfigID: id,
		User:     userInfo.User,
		Org:      userInfo.Org,
		Type:     models.EventTypeClone,
	})
	if err != nil {
		return nil, err
	}

	return &config.Metadata, nil
}

func (s *service) Update(id string, req *models.UpdateResourceConfigRequest, userInfo *models.UserInfo) (*models.Metadata, error) {
	current, err := s.storage.Get(id, userInfo)
	if err != nil {
		return nil, err
	}

	if current.OwnerUserInfo.UserName != userInfo.User || current.OwnerUserInfo.UserOrg != userInfo.Org {
		return nil, storage.ErrNotFound
	}

	config := &models.ResourceConfig{}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	config.Name = current.Name
	config.ID = current.ID
	config.CreatedAt = current.CreatedAt
	config.OwnerUserInfo = &models.OwnerUserInfo{
		UserId:   userInfo.UserId,
		UserName: userInfo.User,
		UserOrg:  userInfo.Org,
	}

	result, err := s.storage.Update(config)
	if err != nil {
		return nil, err
	}

	return &result.Metadata, nil
}

func (s *service) Delete(id string, userInfo *models.UserInfo) error {
	config, err := s.storage.Get(id, userInfo)
	if err != nil {
		return err
	}

	if config.OwnerUserInfo.UserName != userInfo.User || config.OwnerUserInfo.UserOrg != userInfo.Org {
		return storage.ErrNotFound
	}

	return s.storage.Delete(id)
}

func (s *service) CountEventsMetrics(id string, metrics []string, userInfo *models.UserInfo) ([]models.Metric, error) {
	config, err := s.storage.Get(id, userInfo)
	if err != nil {
		return nil, err
	}

	if !config.IsAvailableForUser(userInfo) {
		return nil, ErrAccessForbidden
	}

	return s.storage.CountEventsMetrics(id, metrics)
}

func (s *service) CountEvents(filter *models.EventFilter) (int64, error) {
	return s.storage.CountEvents(filter)
}
