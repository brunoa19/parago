package cloudformation

type NetworkPolicy struct {
	Type       string                  `yaml:"Type"`
	Properties NetworkPolicyProperties `yaml:"Properties"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

type NetworkPolicyProperties struct {
	App        string               `yaml:"App"`
	ShipaHost  string               `yaml:"ShipaHost"`
	ShipaToken string               `yaml:"ShipaToken"`
	Ingress    *NetworkPolicyConfig `yaml:"Ingress,omitempty" json:"ingress,omitempty"`
	Egress     *NetworkPolicyConfig `yaml:"Egress,omitempty" json:"egress,omitempty"`
	RestartApp bool                 `yaml:"RestartApp" json:"restart_app"`
}

type NetworkPolicyConfig struct {
	PolicyMode  string               `yaml:"PolicyMode,omitempty" json:"policy_mode,omitempty"`
	CustomRules []*NetworkPolicyRule `yaml:"CustomRules,omitempty" json:"custom_rules,omitempty"`
}

type NetworkPolicyRule struct {
	ID                string         `yaml:"ID,omitempty" json:"id,omitempty"`
	Enabled           bool           `yaml:"Enabled" json:"enabled"`
	Description       string         `yaml:"Description,omitempty" json:"description,omitempty"`
	Ports             []*NetworkPort `yaml:"Ports,omitempty" json:"ports,omitempty"`
	AllowedApps       []string       `yaml:"AllowedApps,omitempty" json:"allowed_apps,omitempty"`
	AllowedFrameworks []string       `yaml:"AllowedFrameworks,omitempty" json:"allowed_frameworks,omitempty"`
}

type NetworkPort struct {
	Protocol string `yaml:"Protocol,omitempty" json:"protocol,omitempty"`
	Port     int    `yaml:"Port,omitempty" json:"port,omitempty"`
}
