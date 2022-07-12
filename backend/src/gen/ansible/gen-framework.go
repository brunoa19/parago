package ansible

import (
	"encoding/json"

	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func GenerateFramework(cfg shipa.FrameworkConfig) *shipa.Result {
	play := newPlay()

	framework := genFramework(cfg)
	if framework != nil {
		play.Tasks = append(play.Tasks, framework)
	}

	if len(play.Tasks) == 0 {
		return nil
	}

	content, _ := yaml.Marshal([]interface{}{play})
	return &shipa.Result{
		Filename: "play.yml",
		Content:  string(content),
	}
}

func genFramework(cfg shipa.FrameworkConfig) *FrameworkTask {
	// required fields
	if cfg.Name == "" {
		return nil
	}

	t := newFrameworkTask()
	t.Framework = Framework{
		Shipa: credentials,
		Name:  cfg.Name,
	}

	data, _ := json.Marshal(cfg.Resources)
	json.Unmarshal(data, &t.Framework.Resources)

	return t
}
