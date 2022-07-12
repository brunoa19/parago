package ansible

import (
	"fmt"

	"shipa-gen/src/gen/utils"
	"shipa-gen/src/shipa"

	"gopkg.in/yaml.v2"
)

func Generate(cfg shipa.Config) *shipa.Result {
	play := newPlay()

	appDeploy := genAppDeploy(cfg)
	if appDeploy != nil {
		play.Tasks = append(play.Tasks, appDeploy)
	} else {
		app := genApp(cfg)
		if app != nil {
			play.Tasks = append(play.Tasks, app)
		}

		appEnv := genAppEnv(cfg)
		if appEnv != nil {
			play.Tasks = append(play.Tasks, appEnv)
		}
	}

	appCname := genAppCname(cfg)
	if appCname != nil {
		play.Tasks = append(play.Tasks, appCname)
	}

	policy := genNetworkPolicy(cfg)
	if policy != nil {
		play.Tasks = append(play.Tasks, policy)
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

func genAppDeploy(cfg shipa.Config) *AppDeployTask {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	t := newAppDeployTask()
	t.AppDeploy = AppDeploy{
		Shipa: credentials,
		App:   cfg.AppName,
		Image: cfg.Image,
		AppConfig: &AppConfig{
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
	return t
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

func genAppDeployPort(cfg shipa.Config) *Port {
	if cfg.Port == 0 {
		return nil
	}

	return &Port{
		Number:   cfg.Port,
		Protocol: "TCP",
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

func genAppCname(cfg shipa.Config) *AppCnameTask {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	t := newAppCnameTask()
	t.AppCname = AppCname{
		Shipa:   credentials,
		App:     cfg.AppName,
		Cname:   cfg.Cname,
		Encrypt: cfg.Encrypt,
	}
	return t
}

func genAppEnv(cfg shipa.Config) *AppEnvTask {
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

	t := newAppEnvTask()
	t.AppEnv = AppEnv{
		Shipa:     credentials,
		App:       cfg.AppName,
		Envs:      envs,
		Norestart: cfg.Norestart,
		Private:   cfg.Private,
	}
	return t
}

var credentials = Shipa{
	ShipaHost:  "{{ shipa_host }}",
	ShipaToken: "{{ shipa_token }}",
}

func genApp(cfg shipa.Config) *AppTask {
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	t := newAppTask()
	t.App = App{
		Shipa:     credentials,
		Name:      cfg.AppName,
		Teamowner: cfg.Team,
		Framework: cfg.Framework,
		Plan:      cfg.Plan,
		Tags:      utils.ParseValues(cfg.Tags),
	}
	return t
}

func genNetworkPolicy(cfg shipa.Config) *NetworkPolicyTask {
	if cfg.AppName == "" || cfg.NetworkPolicy == nil {
		return nil
	}

	t := newNetworkPolicyTask()
	t.NetworkPolicy = NetworkPolicy{
		Shipa: credentials,
		App:   cfg.AppName,
	}

	utils.CopyJsonData(cfg.NetworkPolicy, &t.NetworkPolicy)

	return t
}
