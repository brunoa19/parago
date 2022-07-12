package crossplane

import (
	"encoding/json"
	"fmt"
	"strings"

	"shipa-gen/src/gen/utils"
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

const apiVersion = "shipa.crossplane.io/v1alpha1"

func Generate(cfg shipa.Config) *shipa.Result {
	var resource []interface{}

	appDeploy := genAppDeploy(cfg)
	if appDeploy != nil {
		resource = append(resource, appDeploy)
	} else {
		app := genApp(cfg)
		if app != nil {
			resource = append(resource, app)
		}

		appEnv := genAppEnv(cfg)
		if appEnv != nil {
			resource = append(resource, appEnv)
		}
	}

	appCname := genAppCname(cfg)
	if appCname != nil {
		resource = append(resource, appCname)
	}

	policy := genNetworkPolicy(cfg)
	if policy != nil {
		resource = append(resource, policy)
	}

	if len(resource) == 0 {
		return nil
	}

	var contents []string
	for _, r := range resource {
		data, _ := yaml.Marshal(r)
		contents = append(contents, string(data))
	}

	return &shipa.Result{
		Filename:  "crossplane.yaml",
		Content:   strings.Join(contents, "---\n"),
		Separator: "\n---\n",
	}
}

func genApp(cfg shipa.Config) *App {
	// required fields
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	app := &App{
		ApiVersion: apiVersion,
		Kind:       "App",
	}
	app.Metadata.Name = cfg.AppName
	app.Spec.ForProvider = AppForProvider{
		Name:      cfg.AppName,
		TeamOwner: cfg.Team,
		Framework: cfg.Framework,
		Plan:      cfg.Plan,
		Tags:      utils.ParseValues(cfg.Tags),
	}

	return app
}

func genAppDeploy(cfg shipa.Config) *AppDeploy {
	// required fields
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	appDeploy := &AppDeploy{
		ApiVersion: apiVersion,
		Kind:       "AppDeploy",
	}
	appDeploy.Metadata.Name = cfg.AppName
	appDeploy.Spec.ForProvider = AppDeployParameters{
		App:   cfg.AppName,
		Image: cfg.Image,
		AppConfig: AppDeployConfig{
			Team:      cfg.Team,
			Framework: cfg.Framework,
			Plan:      cfg.Plan,
			Tags:      utils.ParseValues(cfg.Tags),
			Env:       genAppDeployEnvs(cfg),
		},
		Registry:       genAppDeployRegistry(cfg),
		Port:           genAppDeployPort(cfg),
		Volumes:        genAppDeployVolumes(cfg),
		CanarySettings: genAppDeployCanarySettings(cfg),
		PodAutoScaler:  genAppDeployPodAutoScaler(cfg),
	}

	return appDeploy
}

func genAppDeployEnvs(cfg shipa.Config) []string {
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

func genAppDeployVolumes(cfg shipa.Config) (out []*AppDeployVolume) {
	for _, volume := range cfg.Volumes {
		out = append(out, genAppDeployVolume(volume))
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

func genAppDeployPort(cfg shipa.Config) *AppDeployPort {
	if cfg.Port == 0 {
		return nil
	}

	return &AppDeployPort{
		Number:   cfg.Port,
		Protocol: "TCP",
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

func genAppCname(cfg shipa.Config) *AppCname {
	// required fields
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	appCname := &AppCname{
		ApiVersion: apiVersion,
		Kind:       "AppCname",
	}
	appCname.Metadata.Name = cfg.AppName
	appCname.Spec.ForProvider = AppCnameForProvider{
		App:     cfg.AppName,
		Cname:   cfg.Cname,
		Encrypt: cfg.Encrypt,
	}

	return appCname
}

func genAppEnv(cfg shipa.Config) *AppEnv {
	// required fields
	if cfg.AppName == "" || len(cfg.Envs) == 0 {
		return nil
	}

	appEnv := &AppEnv{
		ApiVersion: apiVersion,
		Kind:       "AppEnv",
	}
	appEnv.Metadata.Name = cfg.AppName
	p := &appEnv.Spec.ForProvider
	p.App = cfg.AppName
	p.AppEnv.Norestart = cfg.Norestart
	p.AppEnv.Private = cfg.Private
	for _, env := range cfg.Envs {
		p.AppEnv.Envs = append(p.AppEnv.Envs, Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	return appEnv
}

func genNetworkPolicy(cfg shipa.Config) *NetworkPolicy {
	// required fields
	if cfg.AppName == "" || cfg.NetworkPolicy == nil {
		return nil
	}

	policy := &NetworkPolicy{
		ApiVersion: apiVersion,
		Kind:       "NetworkPolicy",
	}
	policy.Metadata.Name = cfg.AppName
	p := &policy.Spec.ForProvider
	p.App = cfg.AppName
	p.NetworkPolicy = &NetworkPolicyDetails{}

	data, _ := json.Marshal(cfg.NetworkPolicy)
	json.Unmarshal(data, p.NetworkPolicy)

	return policy
}
