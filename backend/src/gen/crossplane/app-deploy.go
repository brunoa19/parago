package crossplane

type AppDeploy struct {
	ApiVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   Metadata      `yaml:"metadata"`
	Spec       AppDeploySpec `yaml:"spec"`
}

type AppDeploySpec struct {
	ForProvider AppDeployParameters `yaml:"forProvider"`
}

type AppDeployParameters struct {
	App            string                   `json:"-" yaml:"app"`
	Image          string                   `json:"image" yaml:"image"`
	AppConfig      AppDeployConfig          `json:"appConfig" yaml:"appConfig"`
	CanarySettings *AppDeployCanarySettings `json:"canarySettings,omitempty" yaml:"canarySettings,omitempty"`
	PodAutoScaler  *AppDeployPodAutoScaler  `json:"podAutoScaler,omitempty" yaml:"podAutoScaler,omitempty"`
	Port           *AppDeployPort           `json:"port,omitempty" yaml:"port,omitempty"`
	Registry       *AppDeployRegistry       `json:"registry,omitempty" yaml:"registry,omitempty"`
	Volumes        []*AppDeployVolume       `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}

// AppDeployConfig - represents app deploy config
type AppDeployConfig struct {
	Team        string   `json:"team" yaml:"team"`
	Framework   string   `json:"framework" yaml:"framework"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Env         []string `json:"env,omitempty" yaml:"env,omitempty"`
	Plan        string   `json:"plan,omitempty" yaml:"plan,omitempty"`
	Router      string   `json:"router,omitempty" yaml:"router,omitempty"`
	Tags        []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// AppDeployCanarySettings - represents app deploy canary settings
type AppDeployCanarySettings struct {
	StepInterval int64 `json:"stepInterval" yaml:"stepInterval"`
	StepWeight   int64 `json:"stepWeight" yaml:"stepWeight"`
	Steps        int64 `json:"steps" yaml:"steps"`
}

// AppDeployPodAutoScaler - represents app deploy auto scaler
type AppDeployPodAutoScaler struct {
	MaxReplicas                    int64 `json:"maxReplicas" yaml:"maxReplicas"`
	MinReplicas                    int64 `json:"minReplicas" yaml:"minReplicas"`
	TargetCPUUtilizationPercentage int64 `json:"targetCPUUtilizationPercentage" yaml:"targetCPUUtilizationPercentage"`
}

// AppDeployPort - represents app deploy port
type AppDeployPort struct {
	Number   int64  `json:"number" yaml:"number"`
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}

// AppDeployRegistry - represents app deploy registry
type AppDeployRegistry struct {
	User   string `json:"user" yaml:"user"`
	Secret string `json:"secret" yaml:"secret"`
}

// AppDeployVolume - represents app deploy volume
type AppDeployVolume struct {
	Name    string         `json:"name" yaml:"name"`
	Path    string         `json:"mountPath" yaml:"mountPath"`
	Options *VolumeOptions `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`
}

// VolumeOptions - represents additional volume options
type VolumeOptions struct {
	Prop1 string `json:"additionalProp1,omitempty" yaml:"additionalProp1,omitempty"`
	Prop2 string `json:"additionalProp2,omitempty" yaml:"additionalProp2,omitempty"`
	Prop3 string `json:"additionalProp3,omitempty" yaml:"additionalProp3,omitempty"`
}
