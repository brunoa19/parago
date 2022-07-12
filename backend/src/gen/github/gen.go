package github

import (
	"encoding/json"
	"fmt"

	"shipa-gen/src/gen/utils"
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func Generate(cfg shipa.Config) *shipa.Result {
	var action Action

	action.NetworkPolicy = genNetworkPolicy(cfg)
	action.AppCname = genAppCname(cfg)
	action.AppDeploy = genAppDeploy(cfg)
	if action.AppDeploy == nil {
		action.App = genApp(cfg)
		action.AppEnv = genAppEnv(cfg)
	}

	if action.App == nil && action.AppEnv == nil &&
		action.AppCname == nil && action.AppDeploy == nil {
		return nil
	}

	data, _ := yaml.Marshal(action)
	return &shipa.Result{
		Filename: "shipa-action.yml",
		Content:  string(data),
	}
}

func genAppDeploy(cfg shipa.Config) *AppDeploy {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	return &AppDeploy{
		App:   cfg.AppName,
		Image: cfg.Image,
		AppConfig: &AppDeployConfig{
			Team:      cfg.Team,
			Framework: cfg.Framework,
			Plan:      cfg.Plan,
			Tags:      utils.ParseValues(cfg.Tags),
			Env:       genEnvs(cfg),
		},
		Registry:       genAppDeployRegistry(cfg),
		Port:           genAppDeployPort(cfg),
		Volumes:        genAppDeployVolumes(cfg),
		CanarySettings: genAppDeployCanarySettings(cfg),
		PodAutoScaler:  genAppDeployPodAutoScaler(cfg),
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

func genAppDeployRegistry(cfg shipa.Config) *AppDeployRegistry {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return nil
	}

	return &AppDeployRegistry{
		User:   cfg.RegistryUser,
		Secret: cfg.RegistrySecret,
	}
}

func genAppDeployPort(cfg shipa.Config) *AppDeployPort {
	if cfg.Port == 0 {
		return nil
	}

	return &AppDeployPort{
		Number:   cfg.Port,
		Protocol: "TCP",
	}
}

func genAppCname(cfg shipa.Config) *AppCname {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	return &AppCname{
		App:       cfg.AppName,
		Cname:     cfg.Cname,
		Encrypted: cfg.Encrypt,
	}
}

func genApp(cfg shipa.Config) *App {
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	return &App{
		Name:      cfg.AppName,
		TeamOwner: cfg.Team,
		Pool:      cfg.Framework,
		Plan:      cfg.Plan,
		Tags:      utils.ParseValues(cfg.Tags),
	}
}

func genAppEnv(cfg shipa.Config) *AppEnv {
	if cfg.AppName == "" || len(cfg.Envs) == 0 {
		return nil
	}

	var envs []*Env
	for _, env := range cfg.Envs {
		envs = append(envs, &Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	return &AppEnv{
		App:       cfg.AppName,
		Envs:      envs,
		NoRestart: cfg.Norestart,
		Private:   cfg.Private,
	}
}

func genNetworkPolicy(cfg shipa.Config) *NetworkPolicy {
	if cfg.AppName == "" || cfg.NetworkPolicy == nil {
		return nil
	}

	policy := &NetworkPolicy{App: cfg.AppName}
	data, _ := json.Marshal(cfg.NetworkPolicy)
	json.Unmarshal(data, policy)

	return policy
}
