package crossplane

type NetworkPolicy struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   Metadata          `yaml:"metadata"`
	Spec       NetworkPolicySpec `yaml:"spec"`
}

type NetworkPolicySpec struct {
	ForProvider NetworkPolicyForProvider `yaml:"forProvider"`
}

type NetworkPolicyForProvider struct {
	App           string                `yaml:"app"`
	NetworkPolicy *NetworkPolicyDetails `yaml:"networkPolicy,omitempty"`
}

type NetworkPolicyDetails struct {
	Ingress    *NetworkPolicyConfig `yaml:"ingress,omitempty" json:"ingress,omitempty"`
	Egress     *NetworkPolicyConfig `yaml:"egress,omitempty" json:"egress,omitempty"`
	RestartApp bool                 `yaml:"restart_app" json:"restart_app"`
}

type NetworkPolicyConfig struct {
	PolicyMode  string               `yaml:"policy_mode,omitempty" json:"policy_mode,omitempty"`
	CustomRules []*NetworkPolicyRule `yaml:"custom_rules,omitempty" json:"custom_rules,omitempty"`
}

type NetworkPolicyRule struct {
	ID                string         `yaml:"id,omitempty" json:"id,omitempty"`
	Enabled           bool           `yaml:"enabled" json:"enabled"`
	Description       string         `yaml:"description,omitempty" json:"description,omitempty"`
	Ports             []*NetworkPort `yaml:"ports,omitempty" json:"ports,omitempty"`
	AllowedApps       []string       `yaml:"allowed_apps,omitempty" json:"allowed_apps,omitempty"`
	AllowedFrameworks []string       `yaml:"allowed_frameworks,omitempty" json:"allowed_frameworks,omitempty"`
}

type NetworkPort struct {
	Protocol string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
	Port     int    `yaml:"port,omitempty" json:"port,omitempty"`
}
