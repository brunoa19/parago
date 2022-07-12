package github

import (
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func GenerateFramework(cfg shipa.FrameworkConfig) *shipa.Result {
	var action Action

	action.Framework = genFramework(cfg)

	if action.Framework == nil {
		return nil
	}

	data, _ := yaml.Marshal(action)
	return &shipa.Result{
		Filename: "shipa-action.yml",
		Content:  string(data),
	}
}

func genFramework(cfg shipa.FrameworkConfig) *Framework {
	// required fields
	if cfg.Name == "" {
		return nil
	}

	return &Framework{
		Name:      cfg.Name,
		Resources: cfg.Resources,
		DependsOn: cfg.DependsOn,
	}
}
