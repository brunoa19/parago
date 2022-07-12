package ansible

import "shipa-gen/src/shipa"

func newPlay() *Play {
	return &Play{
		Hosts: "localhost",
		Vars: []map[string]string{
			{"shipa_host": "<host>"},
			{"shipa_token": "<token>"},
		},
	}
}

type Play struct {
	Hosts string `yaml:"hosts"`
	Vars  []map[string]string
	Tasks []interface{} `yaml:"tasks"`
}

type AppTask struct {
	Name string `yaml:"name"`
	App  App    `yaml:"shipa_application"`
}

func newAppTask() *AppTask {
	return &AppTask{
		Name: "Create shipa application",
	}
}

type Shipa struct {
	ShipaHost  string `yaml:"shipa_host"`
	ShipaToken string `yaml:"shipa_token"`
}

type App struct {
	Shipa     `yaml:",inline"`
	Name      string   `yaml:"name"`
	Teamowner string   `yaml:"teamowner,omitempty"`
	Framework string   `yaml:"framework,omitempty"`
	Plan      string   `yaml:"plan,omitempty"`
	Tags      []string `yaml:"tags,omitempty"`
}

type AppEnvTask struct {
	Name   string `yaml:"name"`
	AppEnv AppEnv `yaml:"shipa_app_env"`
}

func newAppEnvTask() *AppEnvTask {
	return &AppEnvTask{
		Name: "Create shipa app env",
	}
}

type AppEnv struct {
	Shipa     `yaml:",inline"`
	App       string `yaml:"app"`
	Envs      []Env  `yaml:"envs"`
	Norestart bool   `yaml:"norestart"`
	Private   bool   `yaml:"private"`
}

type Env struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type AppCnameTask struct {
	Name     string   `yaml:"name"`
	AppCname AppCname `yaml:"shipa_app_cname"`
}

func newAppCnameTask() *AppCnameTask {
	return &AppCnameTask{
		Name: "Create shipa app cname",
	}
}

type AppCname struct {
	Shipa   `yaml:",inline"`
	App     string `yaml:"app"`
	Cname   string `yaml:"cname"`
	Encrypt bool   `yaml:"encrypt"`
}

type AppDeployTask struct {
	Name      string    `yaml:"name"`
	AppDeploy AppDeploy `yaml:"shipa_app_deploy"`
}

func newAppDeployTask() *AppDeployTask {
	return &AppDeployTask{
		Name: "Deploy shipa application",
	}
}

type AppDeploy struct {
	Shipa          `yaml:",inline"`
	App            string                   `yaml:"app"`
	Image          string                   `yaml:"image"`
	AppConfig      *AppConfig               `yaml:"appConfig"`
	Registry       *Registry                `yaml:"registry,omitempty"`
	Port           *Port                    `yaml:"port,omitempty"`
	Volumes        []*AppDeployVolume       `yaml:"volumes,omitempty"`
	CanarySettings *AppDeployCanarySettings `yaml:"canarySettings,omitempty"`
	PodAutoScaler  *AppDeployPodAutoScaler  `yaml:"podAutoScaler,omitempty"`
}

type AppDeployCanarySettings struct {
	StepInterval int64 `yaml:"stepInterval"`
	StepWeight   int64 `yaml:"stepWeight"`
	Steps        int64 `yaml:"steps"`
}

type AppDeployPodAutoScaler struct {
	MaxReplicas                    int64 `yaml:"maxReplicas"`
	MinReplicas                    int64 `yaml:"minReplicas"`
	TargetCPUUtilizationPercentage int64 `yaml:"targetCPUUtilizationPercentage"`
}

type AppDeployVolume struct {
	Name    string         `yaml:"name"`
	Path    string         `yaml:"mountPath"`
	Options *VolumeOptions `yaml:"mountOptions,omitempty"`
}

type VolumeOptions struct {
	Prop1 string `yaml:"additionalProp1,omitempty"`
	Prop2 string `yaml:"additionalProp2,omitempty"`
	Prop3 string `yaml:"additionalProp3,omitempty"`
}

type Port struct {
	Number   int64  `yaml:"number"`
	Protocol string `yaml:"protocol"`
}

type Registry struct {
	User   string `yaml:"user"`
	Secret string `yaml:"secret"`
}

type AppConfig struct {
	Team      string   `yaml:"team"`
	Framework string   `yaml:"framework"`
	Plan      string   `yaml:"plan,omitempty"`
	Tags      []string `yaml:"tags,omitempty"`
	Env       []string `yaml:"env,omitempty"`
}

type VolumeTask struct {
	Name   string `yaml:"name"`
	Volume Volume `yaml:"shipa_volume"`
}

func newVolumeTask() *VolumeTask {
	return &VolumeTask{
		Name: "Create shipa volume",
	}
}

type Volume struct {
	Shipa `yaml:",inline"`
	// required
	Name     string `yaml:"name"`
	Capacity string `yaml:"capacity"`
	Plan     string `yaml:"plan"`
	// optional
	AccessModes string         `yaml:"accessModes,omitempty"`
	Opts        *VolumeOptions `yaml:"opts,omitempty"`
}

type FrameworkTask struct {
	Name      string    `yaml:"name"`
	Framework Framework `yaml:"shipa_framework"`
}

func newFrameworkTask() *FrameworkTask {
	return &FrameworkTask{
		Name: "Create shipa framework",
	}
}

type Framework struct {
	Shipa `yaml:",inline"`
	// required
	Name      string           `yaml:"name"`
	Resources *shipa.Resources `json:"resources,omitempty" yaml:"resources,omitempty"`
}
