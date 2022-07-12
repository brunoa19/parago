package models

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessLevelType string

const (
	AccessLevelPrivate      = "private"
	AccessLevelOrganization = "organization"
	AccessLevelPublic       = "public"
	AccessModePersonal      = "personal"
)

type ResourceType string

const (
	ResourceTypeApp       = "app"
	ResourceTypeVolume    = "volume"
	ResourceTypeFramework = "framework"
)

type Provider string

const (
	ProviderCrossplane     = "crossplane"
	ProviderCloudformation = "cloudformation"
	ProviderGithub         = "github"
	ProviderGitlab         = "gitlab"
	ProviderAnsible        = "ansible"
	ProviderTerraform      = "terraform"
	ProviderPulumi         = "pulumi"
	ProviderHelmChart      = "helm_chart"
)

func (p Provider) Validate() error {
	switch p {
	case "":
		return errors.New("provider not set")
	case ProviderCrossplane, ProviderCloudformation, ProviderGithub, ProviderGitlab, ProviderAnsible, ProviderTerraform, ProviderPulumi:
		return nil
	default:
		return fmt.Errorf("invalid provider, use one of: %s", []Provider{ProviderCrossplane, ProviderCloudformation, ProviderGithub, ProviderGitlab, ProviderAnsible, ProviderTerraform, ProviderPulumi})
	}
}

type OwnerUserInfo struct {
	UserId   string `json:"userId" bson:"userId"`
	UserName string `json:"userName" bson:"userName"`
	UserOrg  string `json:"userOrg" bson:"userOrg"`
}

type Metadata struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	AccessLevel AccessLevelType    `json:"accessLevel" bson:"accessLevel"` // Private | Organization | Public
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Deleted     bool               `bson:"deleted" json:"-"`

	Provider      string         `json:"provider" bson:"provider"`
	Type          ResourceType   `json:"resourceType" bson:"resourceType"` // app | volume | framework
	OwnerUserInfo *OwnerUserInfo `json:"userInfo,omitempty" bson:"userInfo,omitempty"`
}

type ResourceConfig struct {
	Metadata   `json:",inline" bson:",inline"`
	Data       map[string]interface{} `json:"data" bson:"data"`
	Definition map[string]interface{} `json:"definition,omitempty" bson:"definition,omitempty"`
}

func (r *ResourceConfig) IsAvailableForUser(user *UserInfo) bool {
	switch {
	case r.AccessLevel == AccessLevelPublic:
		return true
	case r.AccessLevel == AccessLevelOrganization && r.OwnerUserInfo.UserOrg == user.Org:
		return true
	case r.OwnerUserInfo.UserId == user.UserId || r.OwnerUserInfo.UserName == user.User:
		return true
	default:
		return false
	}
}

type ResourceWithPayload struct {
	ResourceConfig `json:",inline"`
	Payload        *Payload `json:"payload"`
}

type ResourceConfigRequest struct {
	AccessLevel AccessLevelType        `json:"accessLevel"` // Private | Organization | Public
	Description string                 `json:"description,omitempty"`
	Provider    Provider               `json:"provider"`
	Type        ResourceType           `json:"resourceType"` // app | volume | framework
	Data        RequestData            `json:"data"`
	Definition  map[string]interface{} `json:"definition,omitempty"`
}

type UpdateResourceConfigRequest struct {
	ResourceConfigRequest
}

type CreateResourceConfigRequest struct {
	Name string `json:"name"`
	ResourceConfigRequest
}

func (r *ResourceConfigRequest) Init() {
	r.Provider = Provider(r.Data.Provider())
	r.Type = ResourceType(r.Data.ResourceType())
}

func (r ResourceConfigRequest) Validate() error {
	if err := r.AccessLevel.Validate(); err != nil {
		return err
	}

	if err := r.Type.Validate(); err != nil {
		return err
	}

	if r.Data == nil || len(r.Data) == 0 {
		return errors.New("missing config data")
	}

	if err := r.Provider.Validate(); err != nil {
		return err
	}

	return nil
}

type RequestData map[string]interface{}

func (r RequestData) Provider() string {
	if r == nil {
		return ""
	}

	val, ok := r["provider"].(string)
	if !ok {
		return ""
	}

	return val
}

func (r RequestData) ResourceType() string {
	if r == nil {
		return ""
	}

	for key := range r {
		switch key {
		case ResourceTypeApp + "s":
			return ResourceTypeApp
		case ResourceTypeVolume + "s":
			return ResourceTypeVolume
		case ResourceTypeFramework + "s":
			return ResourceTypeFramework
		}
	}

	return ""
}

type Payload struct {
	File   *FileData `json:"file,omitempty"`
	Errors []Error   `json:"errors,omitempty"`
}

type Error struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}

type FileData struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (l AccessLevelType) Validate() error {
	switch l {
	case AccessLevelPrivate, AccessLevelPublic, AccessLevelOrganization:
		return nil
	default:
		return fmt.Errorf("invalid accessLevel, use one of: %v", []AccessLevelType{AccessLevelPrivate, AccessLevelPublic, AccessLevelOrganization})
	}
}

func (r ResourceType) Validate() error {
	switch r {
	case ResourceTypeApp, ResourceTypeVolume, ResourceTypeFramework:
		return nil
	default:
		return fmt.Errorf("invalid resourceType, use one of: %v", []ResourceType{ResourceTypeApp, ResourceTypeVolume, ResourceTypeFramework})
	}
}

type Paginator struct {
	Page     int
	PageSize int
}

func (p *Paginator) GetLimit() int64 {
	return int64(p.PageSize)
}

func (p *Paginator) GetOffset() int64 {
	if p.Page > 1 {
		return int64((p.Page - 1) * p.PageSize)
	}

	return 0
}

type EventType string

const (
	EventTypeLoad  = "load"
	EventTypeClone = "clone"
)

type Event struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ConfigID  string             `json:"configId" bson:"configId"`
	Type      EventType          `json:"type" bson:"type"`
	Org       string             `json:"org" bson:"org"`
	User      string             `json:"user" bson:"user"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type EventFilter struct {
	ConfigID string
	Type     EventType
	From     *time.Time
	To       *time.Time
}

type Metric struct {
	Name     string `json:"name" bson:"name"`
	Total    int64  `json:"total" bson:"total"`
	PerDay   int64  `json:"perDay" bson:"perDay"`
	PerWeek  int64  `json:"perWeek" bson:"perWeek"`
	PerMonth int64  `json:"perMonth" bson:"perMonth"`
	PerYear  int64  `json:"perYear" bson:"perYear"`
}
