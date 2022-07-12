package cloudformation

import (
	"encoding/json"
	"fmt"

	"shipa-gen/src/gen/utils"
	"shipa-gen/src/shipa"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
)

const (
	header     = "Resources"
	shipaHost  = "{{resolve:secretsmanager:ShipaHost}}"
	shipaToken = "{{resolve:secretsmanager:ShipaToken}}"
)

func genResourceName(prefix, name string) string {
	return strcase.ToCamel(prefix) + strcase.ToCamel(name)
}

func genAppName(name string) string {
	return genResourceName("App", name)
}

func genAppDeployName(name string) string {
	return genResourceName("AppDeploy", name)
}

func genAppEnvName(name string) string {
	return genResourceName("AppEnv", name)
}

func genAppCnameName(name string) string {
	return genResourceName("AppCname", name)
}

func genNetworkPolicyName(name string) string {
	return genResourceName("NetworkPolicy", name)
}

func Generate(cfg shipa.Config) *shipa.Result {
	resources := make(map[string]interface{})

	appDeploy := genAppDeploy(cfg)
	if appDeploy != nil {
		resources[genAppDeployName(cfg.AppName)] = appDeploy
	} else {
		app := genApp(cfg)
		if app != nil {
			resources[genAppName(cfg.AppName)] = app
		}

		appEnv := genAppEnv(cfg)
		if appEnv != nil {
			resources[genAppEnvName(cfg.AppName)] = app
		}
	}

	appCname := genAppCname(cfg)
	if appCname != nil {
		resources[genAppCnameName(cfg.AppName)] = appCname
	}

	policy := genNetworkPolicy(cfg)
	if policy != nil {
		resources[genNetworkPolicyName(cfg.AppName)] = policy
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

func genApp(cfg shipa.Config) interface{} {
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	return &App{
		Type: "Shipa::Application::Item",
		Properties: AppProperties{
			Name:       cfg.AppName,
			ShipaHost:  shipaHost,
			ShipaToken: shipaToken,
			Teamowner:  cfg.Team,
			Framework:  cfg.Framework,
			Plan:       cfg.Plan,
			Tags:       utils.ParseValues(cfg.Tags),
		},
		DependsOn: genDependencies(cfg.DependsOn, genAppName),
	}
}

func genAppDeploy(cfg shipa.Config) interface{} {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	return &AppDeploy{
		Type: "Shipa::AppDeploy::Item",
		Properties: AppDeployProperties{
			App:            cfg.AppName,
			ShipaHost:      shipaHost,
			ShipaToken:     shipaToken,
			Image:          cfg.Image,
			Registry:       genAppDeployRegistry(cfg),
			Port:           genAppDeployPort(cfg),
			Volumes:        genAppDeployVolumes(cfg),
			CanarySettings: genAppDeployCanarySettings(cfg),
			PodAutoScaler:  genAppDeployPodAutoScaler(cfg),
			AppConfig: AppConfig{
				Team:      cfg.Team,
				Framework: cfg.Framework,
				Plan:      cfg.Plan,
				Tags:      utils.ParseValues(cfg.Tags),
				Env:       genEnvs(cfg),
			},
		},
		DependsOn: genDependencies(cfg.DependsOn, genAppDeployName),
	}
}

func genDependencies(dependsOn []string, genName func(string) string) interface{} {
	var deps []string
	for _, d := range dependsOn {
		deps = append(deps, genName(d))
	}

	switch len(deps) {
	case 0:
		return nil
	case 1:
		return deps[0]
	default:
		return deps
	}
}

func genEnvs(cfg shipa.Config) []string {
	if len(cfg.Envs) == 0 {
		return nil
	}

	var envs []string
	for _, env := range cfg.Envs {
		envs = append(envs, fmt.Sprintf("%s=%s", env.Name, env.Value))
	}
	return envs
}

func genAppDeployPodAutoScaler(cfg shipa.Config) *AppDeployPodAutoScaler {
	if cfg.PodAutoScaler == nil {
		return nil
	}

	return &AppDeployPodAutoScaler{
		MaxReplicas:                    cfg.PodAutoScaler.MaxReplicas,
		MinReplicas:                    cfg.PodAutoScaler.MinReplicas,
		TargetCPUUtilizationPercentage: cfg.PodAutoScaler.TargetCPUUtilizationPercentage,
	}
}

func genAppDeployCanarySettings(cfg shipa.Config) *AppDeployCanarySettings {
	if cfg.CanarySettings == nil {
		return nil
	}

	return &AppDeployCanarySettings{
		Steps:        cfg.CanarySettings.Steps,
		StepWeight:   cfg.CanarySettings.StepWeight,
		StepInterval: cfg.CanarySettings.StepInterval,
	}
}

func genAppDeployVolumes(cfg shipa.Config) (volumes []*AppDeployVolume) {
	for _, volume := range cfg.Volumes {
		volumes = append(volumes, genAppDeployVolume(volume))
	}
	return
}

func genAppDeployVolume(volume *shipa.Volume) *AppDeployVolume {
	if volume == nil {
		return nil
	}

	return &AppDeployVolume{
		Name:    volume.Name,
		Path:    volume.Path,
		Options: genVolumeOptions(volume.Opts),
	}
}

func genVolumeOptions(opts *shipa.VolumeOptions) *VolumeOptions {
	if opts == nil {
		return nil
	}

	return &VolumeOptions{
		Prop1: opts.Prop1,
		Prop2: opts.Prop2,
		Prop3: opts.Prop3,
	}
}

func genAppDeployRegistry(cfg shipa.Config) *Registry {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return nil
	}

	return &Registry{
		User:   cfg.RegistryUser,
		Secret: cfg.RegistrySecret,
	}
}

func genAppDeployPort(cfg shipa.Config) *Port {
	if cfg.Port == 0 {
		return nil
	}

	return &Port{
		Number:   cfg.Port,
		Protocol: "TCP",
	}
}

func genAppEnv(cfg shipa.Config) interface{} {
	if cfg.AppName == "" || len(cfg.Envs) == 0 {
		return nil
	}

	var envs []Env
	for _, env := range cfg.Envs {
		envs = append(envs, Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	return &AppEnv{
		Type: "Shipa::AppEnv::Item",
		Properties: AppEnvProperties{
			App:        cfg.AppName,
			ShipaHost:  shipaHost,
			ShipaToken: shipaToken,
			Norestart:  cfg.Norestart,
			Private:    cfg.Private,
			Envs:       envs,
		},
		DependsOn: genDependencies(cfg.DependsOn, genAppEnvName),
	}
}

func genAppCname(cfg shipa.Config) interface{} {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	return &AppCname{
		Type: "Shipa::AppCname::Item",
		Properties: AppCnameProperties{
			App:        cfg.AppName,
			ShipaHost:  shipaHost,
			ShipaToken: shipaToken,
			Cname:      cfg.Cname,
			Encrypt:    cfg.Encrypt,
		},
		DependsOn: genDependencies(cfg.DependsOn, genAppCnameName),
	}
}

func genNetworkPolicy(cfg shipa.Config) interface{} {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	policy := &NetworkPolicy{
		Type: "Shipa::NetworkPolicy::Item",
		Properties: NetworkPolicyProperties{
			App:        cfg.AppName,
			ShipaHost:  shipaHost,
			ShipaToken: shipaToken,
		},
		DependsOn: genDependencies(cfg.DependsOn, genNetworkPolicyName),
	}

	data, _ := json.Marshal(cfg.NetworkPolicy)
	json.Unmarshal(data, &policy.Properties)

	return policy
}
