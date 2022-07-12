package shipa

type NetworkPolicy struct {
	Ingress    *NetworkPolicyConfig `json:"ingress,omitempty"`
	Egress     *NetworkPolicyConfig `json:"egress,omitempty"`
	RestartApp bool                 `json:"restart_app"`
}

type NetworkPolicyConfig struct {
	PolicyMode  string               `json:"policy_mode,omitempty"`
	CustomRules []*NetworkPolicyRule `json:"custom_rules,omitempty"`
}

type NetworkPolicyRule struct {
	ID                string         `json:"id,omitempty"`
	Enabled           bool           `json:"enabled"`
	Description       string         `json:"description,omitempty"`
	Ports             []*NetworkPort `json:"ports,omitempty"`
	AllowedApps       []string       `json:"allowed_apps,omitempty"`
	AllowedFrameworks []string       `json:"allowed_frameworks,omitempty"`
}

type NetworkPort struct {
	Protocol string `json:"protocol,omitempty"`
	Port     int    `json:"port,omitempty"`
}
