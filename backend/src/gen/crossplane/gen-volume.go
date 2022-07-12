package crossplane

import (
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func GenerateVolume(cfg shipa.VolumeConfig) *shipa.Result {
	volume := genVolume(cfg)
	if volume == nil {
		return nil
	}

	data, _ := yaml.Marshal(volume)
	return &shipa.Result{
		Filename:  "crossplane.yaml",
		Content:   string(data),
		Separator: "\n---\n",
	}
}

func genVolume(cfg shipa.VolumeConfig) *Volume {
	// required fields
	if cfg.Name == "" || cfg.Capacity == "" || cfg.Plan == "" {
		return nil
	}

	volume := &Volume{
		Header: Header{
			ApiVersion: apiVersion,
			Kind:       "Volume",
			Metadata:   Metadata{Name: cfg.Name},
		},
	}
	volume.Spec.ForProvider = VolumeSpec{
		Name:        cfg.Name,
		Capacity:    cfg.Capacity,
		Plan:        cfg.Plan,
		AccessModes: cfg.AccessModes,
		Opts:        genVolumeOptions(cfg.Opts),
	}

	return volume
}
