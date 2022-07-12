package github

import (
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func GenerateVolume(cfg shipa.VolumeConfig) *shipa.Result {
	var action Action

	action.Volume = genVolume(cfg)

	if action.Volume == nil {
		return nil
	}

	data, _ := yaml.Marshal(action)
	return &shipa.Result{
		Filename: "shipa-action.yml",
		Content:  string(data),
	}
}

func genVolume(cfg shipa.VolumeConfig) *Volume {
	// required fields
	if cfg.Name == "" || cfg.Capacity == "" || cfg.Plan == "" {
		return nil
	}

	volume := &Volume{
		Name:        cfg.Name,
		Capacity:    cfg.Capacity,
		Plan:        cfg.Plan,
		AccessModes: cfg.AccessModes,
		Opts:        genVolumeOptions(cfg.Opts),
	}

	return volume
}
