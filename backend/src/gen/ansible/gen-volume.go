package ansible

import (
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func GenerateVolume(cfg shipa.VolumeConfig) *shipa.Result {
	play := newPlay()

	volume := genVolume(cfg)
	if volume != nil {
		play.Tasks = append(play.Tasks, volume)
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

func genVolume(cfg shipa.VolumeConfig) *VolumeTask {
	// required fields
	if cfg.Name == "" || cfg.Capacity == "" || cfg.Plan == "" {
		return nil
	}

	t := newVolumeTask()
	t.Volume = Volume{
		Shipa:       credentials,
		Name:        cfg.Name,
		Capacity:    cfg.Capacity,
		Plan:        cfg.Plan,
		AccessModes: cfg.AccessModes,
		Opts:        genVolumeOptions(cfg.Opts),
	}
	return t
}
