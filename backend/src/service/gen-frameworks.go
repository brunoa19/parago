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

func GenerateFrameworks(cfg shipa.FrameworksConfig) (*models.Payload, error) {
	out := &models.Payload{}
	var results []*shipa.Result
	for _, framework := range cfg.Frameworks {
		framework.Provider = cfg.Provider
		file, err := generateFramework(framework)
		if err != nil {
			out.Errors = append(out.Errors, models.Error{
				Name:  framework.Name,
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

func generateFramework(cfg shipa.FrameworkConfig) (*shipa.Result, error) {
	var data *shipa.Result
	switch cfg.Provider {
	case models.ProviderCrossplane:
		data = crossplane.GenerateFramework(cfg)
	case models.ProviderCloudformation:
		data = cloudformation.GenerateFramework(cfg)
	case models.ProviderGithub, models.ProviderGitlab:
		data = github.GenerateFramework(cfg)
	case models.ProviderAnsible:
		data = ansible.GenerateFramework(cfg)
	case models.ProviderTerraform:
		data = terraform.GenerateFramework(cfg)
	case models.ProviderPulumi:
		data = pulumi.GenerateFramework(cfg)
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
