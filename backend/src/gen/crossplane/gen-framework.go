package crossplane

import (
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func GenerateFramework(cfg shipa.FrameworkConfig) *shipa.Result {
	framework := genFramework(cfg)
	if framework == nil {
		return nil
	}

	data, _ := yaml.Marshal(framework)
	return &shipa.Result{
		Filename:  "crossplane.yaml",
		Content:   string(data),
		Separator: "\n---\n",
	}
}

func genFramework(cfg shipa.FrameworkConfig) *Framework {
	// required fields
	if cfg.Name == "" {
		return nil
	}

	framework := &Framework{
		Header: Header{
			ApiVersion: apiVersion,
			Kind:       "Framework",
			Metadata:   Metadata{Name: cfg.Name},
		},
	}
	framework.Spec.ForProvider = cfg

	return framework
}
