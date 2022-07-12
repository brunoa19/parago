package terraform

import (
	"fmt"
	"strconv"
	"strings"

	"shipa-gen/src/shipa"

	"github.com/iancoleman/strcase"
)

func GenerateFramework(cfg shipa.FrameworkConfig) *shipa.Result {
	header := genMain()

	content := genFramework(cfg)

	if len(content) == 0 {
		return nil
	}

	return &shipa.Result{
		Filename: "main.tf",
		Header:   header,
		Content:  content,
	}
}

func genDependsOn(dependsOn []string, genVar func(string) string) string {
	if dependsOn == nil {
		return ""
	}

	var deps []string
	for _, app := range dependsOn {
		deps = append(deps, genVar(app))
	}

	return fmt.Sprintf("depends_on = [%s]", strings.Join(deps, ", "))
}

func genFramework(cfg shipa.FrameworkConfig) string {
	root := newObject()
	root.dependsOn = genDependsOn(cfg.DependsOn, genFrameworkName)
	framework := newObject()
	root.addObject("framework", framework)

	framework.addField("name", stringValue(cfg.Name))
	framework.addField("provisioner", stringValue("kubernetes"))

	if cfg.Resources != nil {
		resources := newObject()
		framework.addObject("resources", resources)

		if cfg.Resources.General != nil {
			g := cfg.Resources.General
			general := newObject()
			resources.addObject("general", general)

			// setup
			if g.Setup != nil {
				setup := newObject()
				general.addObject("setup", setup)

				setup.addField("default", strconv.FormatBool(g.Setup.Default))
				setup.addField("public", strconv.FormatBool(g.Setup.Public))
				if g.Setup.KubernetesNamespace != "" {
					setup.addField("kubernetes_namespace", stringValue(g.Setup.KubernetesNamespace))
				}
			}

			// plan
			if g.Plan != nil && g.Plan.Name != "" {
				plan := newObject()
				general.addObject("plan", plan)

				plan.addField("name", stringValue(g.Plan.Name))
			}

			// security
			if g.Security != nil {
				security := newObject()
				general.addObject("security", security)

				security.addField("disable_scan", strconv.FormatBool(g.Security.DisableScan))
				if len(g.Security.IgnoreComponents) > 0 {
					security.addField("ignore_components", stringArray(g.Security.IgnoreComponents))
				}
				if len(g.Security.IgnoreCVES) > 0 {
					security.addField("ignore_cves", stringArray(g.Security.IgnoreCVES))
				}
			}

			// access
			if g.Access != nil && len(g.Access.Append) > 0 {
				access := newObject()
				general.addObject("access", access)

				access.addField("append", stringArray(g.Access.Append))
			}

			// router
			if g.Router != "" {
				general.addField("router", stringValue(g.Router))
			}

			// volumes
			if len(g.Volumes) > 0 {
				general.addField("volumes", stringArray(g.Volumes))
			}

			// container policy
			if g.ContainerPolicy != nil && len(g.ContainerPolicy.AllowedHosts) > 0 {
				containerPolicy := newObject()
				general.addObject("container_policy", containerPolicy)

				containerPolicy.addField("allowed_hosts", stringArray(g.ContainerPolicy.AllowedHosts))
			}

			// node selectors
			if g.NodeSelectors != nil {
				nodeSelectors := newObject()
				general.addObject("node_selectors", nodeSelectors)

				if g.NodeSelectors.Terms != nil && (g.NodeSelectors.Terms.Environment != "" || g.NodeSelectors.Terms.OS != "") {
					terms := newObject()
					nodeSelectors.addObject("terms", terms)

					terms.addField("environment", stringValue(g.NodeSelectors.Terms.Environment))
					terms.addField("os", stringValue(g.NodeSelectors.Terms.OS))
				}

				nodeSelectors.addField("strict", strconv.FormatBool(g.NodeSelectors.Strict))
			}

			// pod auto-scaler
			if g.PodAutoScaler != nil {
				podAutoScaler := newObject()
				general.addObject("pod_auto_scaler", podAutoScaler)

				podAutoScaler.addField("min_replicas", strconv.Itoa(g.PodAutoScaler.MinReplicas))
				podAutoScaler.addField("max_replicas", strconv.Itoa(g.PodAutoScaler.MaxReplicas))
				podAutoScaler.addField("target_cpu_utilization_percentage", strconv.Itoa(g.PodAutoScaler.TargetCPUUtilizationPercentage))
				podAutoScaler.addField("disable_app_override", strconv.FormatBool(g.PodAutoScaler.DisableAppOverride))
			}

			// domain policy
			if g.DomainPolicy != nil && len(g.DomainPolicy.AllowedCnames) > 0 {
				domainPolicy := newObject()
				general.addObject("domain_policy", domainPolicy)
				domainPolicy.addField("allowed_cnames", stringArray(g.DomainPolicy.AllowedCnames))
			}

			// app auto-discovery
			if g.AppAutoDiscovery != nil {
				appAutoDiscovery := newObject()
				general.addObject("app_auto_discovery", appAutoDiscovery)

				var labels []*Object
				for _, l := range g.AppAutoDiscovery.AppSelector {
					if l.Label == "" {
						continue
					}
					labels = append(labels, newObject().addField("label", stringValue(l.Label)))
				}

				if labels != nil {
					appAutoDiscovery.addListOfObjects("app_selector", labels)
				}

				appAutoDiscovery.addField("suffix", stringValue(g.AppAutoDiscovery.Suffix))
			}

			// network-policy
			if g.NetworkPolicy != nil {
				networkPolicy := newObject()
				general.addObject("network_policy", networkPolicy)

				// ingress
				if g.NetworkPolicy.Ingress != nil {
					networkPolicy.addObject("ingress", genFrameworkNetworkPolicyConfig(g.NetworkPolicy.Ingress))
				}

				// egress
				if g.NetworkPolicy.Egress != nil {
					networkPolicy.addObject("egress", genFrameworkNetworkPolicyConfig(g.NetworkPolicy.Egress))
				}

				networkPolicy.addField("disable_app_policies", strconv.FormatBool(g.NetworkPolicy.DisableAppPolicies))
			}
		}
	}

	return fmt.Sprintf(`
# Set volume
resource "shipa_framework" "%s" %s
`, strcase.ToSnake(cfg.Name), root.String())
}

func genFrameworkName(name string) string {
	return "shipa_framework." + strcase.ToSnake(name)
}

func genFrameworkNetworkPolicyConfig(cfg *shipa.FrameworkNetworkPolicyConfig) *Object {
	obj := newObject()

	if cfg.PolicyMode != "" {
		obj.addField("policy_mode", stringValue(cfg.PolicyMode))
	}

	var customRules []*Object
	for _, r := range cfg.CustomRules {
		customRules = append(customRules, genFrameworkNetworkPolicyRule(r))
	}

	if customRules != nil {
		obj.addListOfObjects("custom_rules", customRules)
	}

	var shipaRules []*Object
	for _, r := range cfg.ShipaRules {
		shipaRules = append(shipaRules, genFrameworkNetworkPolicyRule(r))
	}

	if shipaRules != nil {
		obj.addListOfObjects("shipa_rules", shipaRules)
	}

	if cfg.ShipaRulesEnabled != nil {
		obj.addField("shipa_rules_enabled", stringArray(cfg.ShipaRulesEnabled))
	}

	return obj
}

func genFrameworkNetworkPolicyRule(cfg *shipa.FrameworkNetworkPolicyRule) *Object {
	obj := newObject()

	if cfg.ID != "" {
		obj.addField("id", stringValue(cfg.ID))
	}

	obj.addField("enabled", strconv.FormatBool(cfg.Enabled))

	if cfg.Description != "" {
		obj.addField("description", stringValue(cfg.Description))
	}

	// ports
	var ports []*Object
	for _, p := range cfg.Ports {
		ports = append(ports, genFrameworkNetworkPort(p))
	}

	if ports != nil {
		obj.addListOfObjects("ports", ports)
	}

	// peers
	var peers []*Object
	for _, p := range cfg.Peers {
		ports = append(ports, genFrameworkNetworkPeer(p))
	}

	if peers != nil {
		obj.addListOfObjects("peers", peers)
	}

	if len(cfg.AllowedApps) > 0 {
		obj.addField("allowed_apps", stringArray(cfg.AllowedApps))
	}

	if len(cfg.AllowedPools) > 0 {
		obj.addField("allowed_frameworks", stringArray(cfg.AllowedPools))
	}

	return obj
}

func genFrameworkNetworkPort(cfg *shipa.FrameworkNetworkPort) *Object {
	obj := newObject()
	if cfg.Protocol != "" {
		obj.addField("protocol", stringValue(cfg.Protocol))
	}

	if cfg.Port > 0 {
		obj.addField("port", strconv.Itoa(cfg.Port))
	}

	return obj
}

func genFrameworkNetworkPeer(cfg *shipa.FrameworkNetworkPeer) *Object {
	obj := newObject()

	if cfg.PodSelector != nil {
		obj.addObject("pod_selector", genNetworkPeerSelector(cfg.PodSelector))
	}

	if cfg.NamespaceSelector != nil {
		obj.addObject("namespace_selector", genNetworkPeerSelector(cfg.NamespaceSelector))
	}

	if len(cfg.IPBlock) > 0 {
		obj.addField("ip_block", stringArray(cfg.IPBlock))
	}

	return obj
}

func genNetworkPeerSelector(cfg *shipa.NetworkPeerSelector) *Object {
	obj := newObject()

	if cfg.MatchLabels != nil {
		obj.addObject("match_labels", genMatchLabels(cfg.MatchLabels))
	}

	var expressions []*Object
	for _, selector := range cfg.MatchExpressions {
		expressions = append(expressions, genSelectorExpression(selector))
	}

	if expressions != nil {
		obj.addListOfObjects("match_expressions", expressions)
	}

	return obj
}

func genSelectorExpression(cfg *shipa.SelectorExpression) *Object {
	obj := newObject()

	if cfg.Key != "" {
		obj.addField("key", stringValue(cfg.Key))
	}

	if cfg.Operator != "" {
		obj.addField("operator", stringValue(cfg.Operator))
	}

	if len(cfg.Values) > 0 {
		obj.addField("values", stringArray(cfg.Values))
	}

	return obj
}

func genMatchLabels(cfg map[string]string) *Object {
	obj := newObject()

	for key, val := range cfg {
		obj.addField(key, stringValue(val))
	}

	return obj
}
