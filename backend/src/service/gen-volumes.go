package service

import (
	"errors"
	"fmt"

	"shipa-gen/src/gen/ansible"
	"shipa-gen/src/gen/cloudformation"
	"shipa-gen/src/gen/crossplane"
	"shipa-gen/src/gen/github"
	"shipa-gen/src/gen/pulumi"
	"shipa-gen/src/gen/terraform"
	"shipa-gen/src/models"
	"shipa-gen/src/shipa"
)

func GenerateVolumes(cfg shipa.VolumesConfig) (*models.Payload, error) {
	out := &models.Payload{}
	var results []*shipa.Result
	for _, volume := range cfg.Volumes {
		volume.Provider = cfg.Provider
		file, err := generateVolume(volume)
		if err != nil {
			out.Errors = append(out.Errors, models.Error{
				Name:  volume.Name,
				Error: err.Error(),
			})
			continue
		}

		if file != nil {
			results = append(results, file)
		}
	}

	file, err := combineResultFiles(results)
	if err != nil {
		return nil, err
	}

	out.File = file
	return out, nil
}

func generateVolume(cfg shipa.VolumeConfig) (*shipa.Result, error) {
	var data *shipa.Result
	switch cfg.Provider {
	case models.ProviderCrossplane:
		data = crossplane.GenerateVolume(cfg)
	case models.ProviderCloudformation:
		data = cloudformation.GenerateVolume(cfg)
	case models.ProviderGithub, models.ProviderGitlab:
		data = github.GenerateVolume(cfg)
	case models.ProviderAnsible:
		data = ansible.GenerateVolume(cfg)
	case models.ProviderTerraform:
		data = terraform.GenerateVolume(cfg)
	case models.ProviderPulumi:
		data = pulumi.GenerateVolume(cfg)
	default:
		return nil, fmt.Errorf("not supported provider: %s", cfg.Provider)
	}

	if data == nil {
		return nil, errors.New("not data was generated")
	}

	data.Name = cfg.Name
	data.DependsOn = cfg.DependsOn
	return data, nil
}
