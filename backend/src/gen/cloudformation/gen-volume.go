package cloudformation

import (
	"fmt"

	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func genVolumeName(name string) string {
	return genResourceName("Volume", name)
}

func GenerateVolume(cfg shipa.VolumeConfig) *shipa.Result {
	resources := make(map[string]interface{})

	volume := genVolume(cfg)
	if volume != nil {
		resources[genVolumeName(cfg.Name)] = volume
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

func genVolume(cfg shipa.VolumeConfig) interface{} {
	// required fields
	if cfg.Name == "" || cfg.Capacity == "" || cfg.Plan == "" {
		return nil
	}

	return &Volume{
		Type: "Shipa::Volume::Item",
		Properties: VolumeProperties{
			Shipa: Shipa{
				ShipaHost:  shipaHost,
				ShipaToken: shipaToken,
			},
			Name:        cfg.Name,
			Capacity:    cfg.Capacity,
			Plan:        cfg.Plan,
			AccessModes: cfg.AccessModes,
			Opts:        genVolumeOptions(cfg.Opts),
		},
		DependsOn: genDependencies(cfg.DependsOn, genVolumeName),
	}

}
