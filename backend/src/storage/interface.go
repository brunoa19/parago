package storage

import "shipa-gen/src/models"

type Storage interface {
	FindByName(name string, userInfo *models.UserInfo) ([]*models.ResourceConfig, error)
	List(userInfo *models.UserInfo, accessLevel models.AccessLevelType, paginator models.Paginator) ([]*models.ResourceConfig, error)
	Search(search string, userInfo *models.UserInfo) ([]*models.ResourceConfig, error)
	Create(config *models.ResourceConfig) (*models.ResourceConfig, error)
	Get(id string, userInfo *models.UserInfo) (*models.ResourceConfig, error)
	Clone(id string, userInfo *models.UserInfo) (*models.ResourceConfig, error)
	Update(config *models.ResourceConfig) (*models.ResourceConfig, error)
	Delete(id string) error

	AddEvent(event *models.Event) error
	CountEvents(filter *models.EventFilter) (int64, error)
	CountEventsMetrics(configId string, metrics []string) ([]models.Metric, error)
}
