package cloudformation

type Framework struct {
	Type       string              `yaml:"Type"`
	Properties FrameworkProperties `yaml:"Properties"`
	DependsOn  interface{}         `yaml:"DependsOn,omitempty"`
}

type FrameworkProperties struct {
	Shipa     `yaml:",inline"`
	Name      string     `json:"name" yaml:"Name"`
	Resources *Resources `json:"resources,omitempty" yaml:"Resources,omitempty"`
}

type Resources struct {
	General *General `json:"general,omitempty" yaml:"General,omitempty"`
}

type General struct {
	Setup            *FrameworkSetup             `json:"setup,omitempty" yaml:"Setup,omitempty"`
	Plan             *FrameworkPlan              `json:"plan,omitempty" yaml:"Plan,omitempty"`
	Security         *FrameworkSecurity          `json:"security,omitempty" yaml:"Security,omitempty"`
	Access           *FrameworkAccess            `json:"access,omitempty" yaml:"Access,omitempty"`
	Router           string                      `json:"router,omitempty" yaml:"Router,omitempty"`
	Volumes          []string                    `json:"volumes,omitempty" yaml:"Volumes,omitempty"`
	ContainerPolicy  *FrameworkContainerPolicy   `json:"containerPolicy,omitempty" yaml:"ContainerPolicy,omitempty"`
	NodeSelectors    *FrameworkNodeSelectors     `json:"nodeSelectors,omitempty" yaml:"NodeSelectors,omitempty"`
	PodAutoScaler    *FrameworkPodAutoScaler     `json:"podAutoScaler,omitempty" yaml:"PodAutoScaler,omitempty"`
	DomainPolicy     *FrameworkDomainPolicy      `json:"domainPolicy,omitempty" yaml:"DomainPolicy,omitempty"`
	AppAutoDiscovery *FrameworkAppAutoDiscovery  `json:"appAutoDiscovery,omitempty" yaml:"AppAutoDiscovery,omitempty"`
	NetworkPolicy    *FrameworkPoolNetworkPolicy `json:"networkPolicy,omitempty" yaml:"NetworkPolicy,omitempty"`
}

type FrameworkAppAutoDiscovery struct {
	AppSelector []*AppSelectorLabels `json:"appSelector,omitempty" yaml:"AppSelector,omitempty"`
	Suffix      string               `json:"suffix" yaml:"Suffix"`
}

type AppSelectorLabels struct {
	Label string `json:"label,omitempty" yaml:"Label,omitempty"`
}

type FrameworkDomainPolicy struct {
	AllowedCnames []string `json:"allowedCnames,omitempty" yaml:"AllowedCnames,omitempty"`
}

type FrameworkPodAutoScaler struct {
	MinReplicas                    int  `json:"minReplicas" yaml:"MinReplicas"`
	MaxReplicas                    int  `json:"maxReplicas" yaml:"MaxReplicas"`
	TargetCPUUtilizationPercentage int  `json:"targetCPUUtilizationPercentage" yaml:"TargetCPUUtilizationPercentage"`
	DisableAppOverride             bool `json:"disableAppOverride" yaml:"DisableAppOverride"`
}

type FrameworkNodeSelectors struct {
	Terms  *NodeSelectorsTerms `json:"terms,omitempty" yaml:"Terms,omitempty"`
	Strict bool                `json:"strict" yaml:"Strict"`
}

type NodeSelectorsTerms struct {
	Environment string `json:"environment,omitempty" yaml:"Environment,omitempty"`
	OS          string `json:"os,omitempty" yaml:"OS,omitempty"`
}

type FrameworkContainerPolicy struct {
	AllowedHosts []string `json:"allowedHosts,omitempty" yaml:"AllowedHosts,omitempty"`
}

type FrameworkAccess struct {
	Append []string `json:"append,omitempty" yaml:"Append,omitempty"`
}

type FrameworkSecurity struct {
	DisableScan      bool     `json:"disableScan" yaml:"DisableScan"`
	IgnoreComponents []string `json:"ignoreComponents,omitempty" yaml:"IgnoreComponents,omitempty"`
	IgnoreCVES       []string `json:"ignoreCves,omitempty" yaml:"IgnoreCves,omitempty"`
}

type FrameworkPlan struct {
	Name string `json:"name,omitempty" yaml:"Name,omitempty"`
}

type FrameworkSetup struct {
	Default             bool   `json:"default" yaml:"Default"`
	Public              bool   `json:"public" yaml:"Public"`
	KubernetesNamespace string `json:"kubernetesNamespace,omitempty" yaml:"KubernetesNamespace,omitempty"`
}

type FrameworkPoolNetworkPolicy struct {
	Ingress            *FrameworkNetworkPolicyConfig `json:"ingress,omitempty" yaml:"Ingress,omitempty"`
	Egress             *FrameworkNetworkPolicyConfig `json:"egress,omitempty" yaml:"Egress,omitempty"`
	DisableAppPolicies bool                          `json:"disableAppPolicies" yaml:"DisableAppPolicies"`
}

type FrameworkNetworkPolicyConfig struct {
	PolicyMode        string                        `json:"policy_mode,omitempty" yaml:"PolicyMode,omitempty"`
	CustomRules       []*FrameworkNetworkPolicyRule `json:"custom_rules,omitempty" yaml:"CustomRules,omitempty"`
	ShipaRules        []*FrameworkNetworkPolicyRule `json:"shipa_rules,omitempty" yaml:"ShipaRules,omitempty"`
	ShipaRulesEnabled []string                      `json:"shipa_rules_enabled,omitempty" yaml:"ShipaRulesEnabled,omitempty"`
}

type FrameworkNetworkPolicyRule struct {
	ID           string                  `json:"id,omitempty" yaml:"ID,omitempty"`
	Enabled      bool                    `json:"enabled" yaml:"Enabled"`
	Description  string                  `json:"description,omitempty" yaml:"Description,omitempty"`
	Ports        []*FrameworkNetworkPort `json:"ports,omitempty" yaml:"Ports,omitempty"`
	Peers        []*FrameworkNetworkPeer `json:"peers,omitempty" yaml:"Peers,omitempty"`
	AllowedApps  []string                `json:"allowed_apps,omitempty" yaml:"AllowedApps,omitempty"`
	AllowedPools []string                `json:"allowed_pools,omitempty" yaml:"AllowedPools,omitempty"`
}

type FrameworkNetworkPort struct {
	Protocol string `json:"protocol,omitempty" yaml:"Protocol,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"Port,omitempty"`
}

type FrameworkNetworkPeer struct {
	PodSelector       *NetworkPeerSelector `json:"podSelector,omitempty" yaml:"PodSelector,omitempty"`
	NamespaceSelector *NetworkPeerSelector `json:"namespaceSelector,omitempty" yaml:"NamespaceSelector,omitempty"`
	IPBlock           []string             `json:"ipBlock,omitempty" yaml:"IPBlock,omitempty"`
}

type NetworkPeerSelector struct {
	MatchLabels      map[string]string     `json:"matchLabels,omitempty" yaml:"MatchLabels,omitempty"`
	MatchExpressions []*SelectorExpression `json:"matchExpressions,omitempty" yaml:"MatchExpressions,omitempty"`
}

type SelectorExpression struct {
	Key      string   `json:"key,omitempty" yaml:"Key,omitempty"`
	Operator string   `json:"operator,omitempty" yaml:"Operator,omitempty"`
	Values   []string `json:"values,omitempty" yaml:"Values,omitempty"`
}
