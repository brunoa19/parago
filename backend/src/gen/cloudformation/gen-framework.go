package cloudformation

import (
	"encoding/json"
	"fmt"

	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func genFrameworkName(name string) string {
	return genResourceName("Framework", name)
}

func GenerateFramework(cfg shipa.FrameworkConfig) *shipa.Result {
	resources := make(map[string]interface{})

	volume := genFramework(cfg)
	if volume != nil {
		resources[genFrameworkName(cfg.Name)] = volume
	}

	if len(resources) == 0 {
		return nil
	}

	data, _ := yaml.Marshal(map[string]interface{}{
		header: resources,
	})

	headLine := fmt.Sprintf("\n%s:\n", header)
	content := string(data)[len(headLine):]

	return &shipa.Result{
		Filename: "cloudformation.yaml",
		Content:  content,
		Header:   headLine,
	}
}

func genFramework(cfg shipa.FrameworkConfig) interface{} {
	// required fields
	if cfg.Name == "" {
		return nil
	}

	framework := &Framework{
		Type: "Shipa::Framework::Item",
		Properties: FrameworkProperties{
			Shipa: Shipa{
				ShipaHost:  shipaHost,
				ShipaToken: shipaToken,
			},
		},
		DependsOn: genDependencies(cfg.DependsOn, genFrameworkName),
	}

	data, _ := json.Marshal(cfg)
	json.Unmarshal(data, &framework.Properties)

	return framework
}
