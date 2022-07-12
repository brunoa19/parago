package shipa

type FrameworksConfig struct {
	Provider   string            `json:"provider"`
	Frameworks []FrameworkConfig `json:"frameworks"`
}

type FrameworkConfig struct {
	// required
	Provider string `json:"-" yaml:"-"`
	Name     string `json:"name" yaml:"shipaFramework"`

	Resources *Resources `json:"resources,omitempty" yaml:"resources,omitempty"`

	DependsOn []string `json:"dependsOn,omitempty" yaml:"-"`
}

type Resources struct {
	General *General `json:"general,omitempty" yaml:"general,omitempty"`
}

type General struct {
	Setup            *FrameworkSetup             `json:"setup,omitempty" yaml:"setup,omitempty"`
	Plan             *FrameworkPlan              `json:"plan,omitempty" yaml:"plan,omitempty"`
	Security         *FrameworkSecurity          `json:"security,omitempty" yaml:"security,omitempty"`
	Access           *FrameworkAccess            `json:"access,omitempty" yaml:"access,omitempty"`
	Router           string                      `json:"router,omitempty" yaml:"router,omitempty"`
	Volumes          []string                    `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	ContainerPolicy  *FrameworkContainerPolicy   `json:"containerPolicy,omitempty" yaml:"containerPolicy,omitempty"`
	NodeSelectors    *FrameworkNodeSelectors     `json:"nodeSelectors,omitempty" yaml:"nodeSelectors,omitempty"`
	PodAutoScaler    *FrameworkPodAutoScaler     `json:"podAutoScaler,omitempty" yaml:"podAutoScaler,omitempty"`
	DomainPolicy     *FrameworkDomainPolicy      `json:"domainPolicy,omitempty" yaml:"domainPolicy,omitempty"`
	AppAutoDiscovery *FrameworkAppAutoDiscovery  `json:"appAutoDiscovery,omitempty" yaml:"appAutoDiscovery,omitempty"`
	NetworkPolicy    *FrameworkPoolNetworkPolicy `json:"networkPolicy,omitempty" yaml:"networkPolicy,omitempty"`
}

type FrameworkAppAutoDiscovery struct {
	AppSelector []*AppSelectorLabels `json:"appSelector,omitempty" yaml:"appSelector,omitempty"`
	Suffix      string               `json:"suffix" yaml:"suffix"`
}

type AppSelectorLabels struct {
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
}

type FrameworkDomainPolicy struct {
	AllowedCnames []string `json:"allowedCnames,omitempty" yaml:"allowedCnames,omitempty"`
}

type FrameworkPodAutoScaler struct {
	MinReplicas                    int  `json:"minReplicas" yaml:"minReplicas"`
	MaxReplicas                    int  `json:"maxReplicas" yaml:"maxReplicas"`
	TargetCPUUtilizationPercentage int  `json:"targetCPUUtilizationPercentage" yaml:"targetCPUUtilizationPercentage"`
	DisableAppOverride             bool `json:"disableAppOverride" yaml:"disableAppOverride"`
}

type FrameworkNodeSelectors struct {
	Terms  *NodeSelectorsTerms `json:"terms,omitempty" yaml:"terms,omitempty"`
	Strict bool                `json:"strict" yaml:"strict"`
}

type NodeSelectorsTerms struct {
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`
	OS          string `json:"os,omitempty" yaml:"os,omitempty"`
}

type FrameworkContainerPolicy struct {
	AllowedHosts []string `json:"allowedHosts,omitempty" yaml:"allowedHosts,omitempty"`
}

type FrameworkAccess struct {
	Append []string `json:"append,omitempty" yaml:"append,omitempty"`
}

type FrameworkSecurity struct {
	DisableScan      bool     `json:"disableScan" yaml:"disableScan"`
	IgnoreComponents []string `json:"ignoreComponents,omitempty" yaml:"ignoreComponents,omitempty"`
	IgnoreCVES       []string `json:"ignoreCves,omitempty" yaml:"ignoreCves,omitempty"`
}

type FrameworkPlan struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type FrameworkSetup struct {
	Default             bool   `json:"default" yaml:"default"`
	Public              bool   `json:"public" yaml:"public"`
	KubernetesNamespace string `json:"kubernetesNamespace,omitempty" yaml:"kubernetesNamespace,omitempty"`
}

type FrameworkPoolNetworkPolicy struct {
	Ingress            *FrameworkNetworkPolicyConfig `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	Egress             *FrameworkNetworkPolicyConfig `json:"egress,omitempty" yaml:"egress,omitempty"`
	DisableAppPolicies bool                          `json:"disableAppPolicies" yaml:"disableAppPolicies"`
}

type FrameworkNetworkPolicyConfig struct {
	PolicyMode        string                        `json:"policy_mode,omitempty" yaml:"policy_mode,omitempty"`
	CustomRules       []*FrameworkNetworkPolicyRule `json:"custom_rules,omitempty" yaml:"custom_rules,omitempty"`
	ShipaRules        []*FrameworkNetworkPolicyRule `json:"shipa_rules,omitempty" yaml:"shipa_rules,omitempty"`
	ShipaRulesEnabled []string                      `json:"shipa_rules_enabled,omitempty" yaml:"shipa_rules_enabled,omitempty"`
}

type FrameworkNetworkPolicyRule struct {
	ID           string                  `json:"id,omitempty" yaml:"id,omitempty"`
	Enabled      bool                    `json:"enabled" yaml:"enabled"`
	Description  string                  `json:"description,omitempty" yaml:"description,omitempty"`
	Ports        []*FrameworkNetworkPort `json:"ports,omitempty" yaml:"ports,omitempty"`
	Peers        []*FrameworkNetworkPeer `json:"peers,omitempty" yaml:"peers,omitempty"`
	AllowedApps  []string                `json:"allowed_apps,omitempty" yaml:"allowed_apps,omitempty"`
	AllowedPools []string                `json:"allowed_pools,omitempty" yaml:"allowed_frameworks,omitempty"`
}

type FrameworkNetworkPort struct {
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
}

type FrameworkNetworkPeer struct {
	PodSelector       *NetworkPeerSelector `json:"podSelector,omitempty" yaml:"podSelector,omitempty"`
	NamespaceSelector *NetworkPeerSelector `json:"namespaceSelector,omitempty" yaml:"namespaceSelector,omitempty"`
	IPBlock           []string             `json:"ipBlock,omitempty" yaml:"ipBlock,omitempty"`
}

type NetworkPeerSelector struct {
	MatchLabels      map[string]string     `json:"matchLabels,omitempty" yaml:"matchLabels,omitempty"`
	MatchExpressions []*SelectorExpression `json:"matchExpressions,omitempty" yaml:"matchExpressions,omitempty"`
}

type SelectorExpression struct {
	Key      string   `json:"key,omitempty" yaml:"key,omitempty"`
	Operator string   `json:"operator,omitempty" yaml:"operator,omitempty"`
	Values   []string `json:"values,omitempty" yaml:"values,omitempty"`
}
