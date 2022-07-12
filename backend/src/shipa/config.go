package shipa

type AppsConfig struct {
	Provider string   `json:"provider"`
	Apps     []Config `json:"apps"`
}

type Config struct {
	Provider  string `json:"-"`
	AppName   string `json:"appName,omitempty"`
	Team      string `json:"team,omitempty"`
	Framework string `json:"framework,omitempty"`
	Plan      string `json:"plan,omitempty"`
	Tags      string `json:"tags,omitempty"`

	Image          string          `json:"image,omitempty"`
	RegistryUser   string          `json:"registryUser,omitempty"`
	RegistrySecret string          `json:"registrySecret,omitempty"`
	Port           int64           `json:"port"`
	CanarySettings *CanarySettings `json:"canarySettings,omitempty"`
	PodAutoScaler  *PodAutoScaler  `json:"podAutoScaler,omitempty"`

	Cname   string `json:"cname,omitempty"`
	Encrypt bool   `json:"encrypt"`

	// deprecated
	EnvName string `json:"envName,omitempty"`
	// deprecated
	EnvValue string `json:"envValue,omitempty"`

	Envs []Env `json:"envs,omitempty"`
	// deprecated
	Norestart bool `json:"norestart"`
	// deprecated
	Private bool `json:"private"`

	NetworkPolicy *NetworkPolicy `json:"network-policy,omitempty"`
	Volumes       []*Volume      `json:"volumes,omitempty"`
	DependsOn     []string       `json:"dependsOn,omitempty"`
}

type CanarySettings struct {
	StepInterval int64 `json:"stepInterval"`
	StepWeight   int64 `json:"stepWeight"`
	Steps        int64 `json:"steps"`
}

type PodAutoScaler struct {
	MaxReplicas                    int64 `json:"maxReplicas"`
	MinReplicas                    int64 `json:"minReplicas"`
	TargetCPUUtilizationPercentage int64 `json:"targetCPUUtilizationPercentage"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Result struct {
	Name      string
	Filename  string
	Header    string
	Content   string
	Separator string
	DependsOn []string
}

type Volume struct {
	Name string         `json:"name"`
	Path string         `json:"mountPath"`
	Opts *VolumeOptions `json:"mountOptions,omitempty"`
}

type VolumeOptions struct {
	Prop1 string `json:"additionalProp1,omitempty"`
	Prop2 string `json:"additionalProp2,omitempty"`
	Prop3 string `json:"additionalProp3,omitempty"`
}
