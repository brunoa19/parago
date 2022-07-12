package pulumi

import (
	"fmt"
	"strconv"
	"strings"

	"shipa-gen/src/shipa"
)

func GenerateFramework(cfg shipa.FrameworkConfig) *shipa.Result {
	header := genMain()

	var content string
	if hasFramework(cfg) {
		content = genFramework(cfg)
	}

	if len(content) == 0 {
		return nil
	}

	return &shipa.Result{
		Filename: "index.ts",
		Header:   header,
		Content:  content,
	}
}

func hasFramework(cfg shipa.FrameworkConfig) bool {
	return cfg.Name != ""
}

func stringValue(val string) string {
	return fmt.Sprintf(`"%s"`, val)
}

func stringArray(values []string) string {
	return fmt.Sprintf(`["%s"]`, strings.Join(values, `", "`))
}

func genFramework(cfg shipa.FrameworkConfig) string {
	root := newObject()
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
					setup.addField("kubernetesNamespace", stringValue(g.Setup.KubernetesNamespace))
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

				security.addField("disableScan", strconv.FormatBool(g.Security.DisableScan))
				if len(g.Security.IgnoreComponents) > 0 {
					security.addField("ignoreComponents", stringArray(g.Security.IgnoreComponents))
				}
				if len(g.Security.IgnoreCVES) > 0 {
					security.addField("ignoreCves", stringArray(g.Security.IgnoreCVES))
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
				general.addObject("containerPolicy", containerPolicy)

				containerPolicy.addField("allowedHosts", stringArray(g.ContainerPolicy.AllowedHosts))
			}

			// node selectors
			if g.NodeSelectors != nil {
				nodeSelectors := newObject()
				general.addObject("nodeSelectors", nodeSelectors)

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
				general.addObject("podAutoScaler", podAutoScaler)

				podAutoScaler.addField("minReplicas", strconv.Itoa(g.PodAutoScaler.MinReplicas))
				podAutoScaler.addField("maxReplicas", strconv.Itoa(g.PodAutoScaler.MaxReplicas))
				podAutoScaler.addField("targetCPUUtilizationPercentage", strconv.Itoa(g.PodAutoScaler.TargetCPUUtilizationPercentage))
				podAutoScaler.addField("disableAppOverride", strconv.FormatBool(g.PodAutoScaler.DisableAppOverride))
			}

			// domain policy
			if g.DomainPolicy != nil && len(g.DomainPolicy.AllowedCnames) > 0 {
				domainPolicy := newObject()
				general.addObject("domainPolicy", domainPolicy)
				domainPolicy.addField("allowedCnames", stringArray(g.DomainPolicy.AllowedCnames))
			}

			// app auto-discovery
			if g.AppAutoDiscovery != nil {
				appAutoDiscovery := newObject()
				general.addObject("appAutoDiscovery", appAutoDiscovery)

				var labels []*Object
				for _, l := range g.AppAutoDiscovery.AppSelector {
					if l.Label == "" {
						continue
					}
					labels = append(labels, newObject().addField("label", stringValue(l.Label)))
				}

				if labels != nil {
					appAutoDiscovery.addListOfObjects("appSelector", labels)
				}

				appAutoDiscovery.addField("suffix", stringValue(g.AppAutoDiscovery.Suffix))
			}

			// network-policy
			if g.NetworkPolicy != nil {
				networkPolicy := newObject()
				general.addObject("networkPolicy", networkPolicy)

				// ingress
				if g.NetworkPolicy.Ingress != nil {
					networkPolicy.addObject("ingress", genFrameworkNetworkPolicyConfig(g.NetworkPolicy.Ingress))
				}

				// egress
				if g.NetworkPolicy.Egress != nil {
					networkPolicy.addObject("egress", genFrameworkNetworkPolicyConfig(g.NetworkPolicy.Egress))
				}

				networkPolicy.addField("disableAppPolicies", strconv.FormatBool(g.NetworkPolicy.DisableAppPolicies))
			}
		}
	}

	return fmt.Sprintf(`
const %s = new shipa.Framework("%s", %s%s);
`, genFrameworkVarName(cfg.Name), genFrameworkResourceName(cfg.Name), root.String(), genDepends(cfg.DependsOn, genFrameworkVarName))
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
		obj.addObject("podSelector", genNetworkPeerSelector(cfg.PodSelector))
	}

	if cfg.NamespaceSelector != nil {
		obj.addObject("namespaceSelector", genNetworkPeerSelector(cfg.NamespaceSelector))
	}

	if len(cfg.IPBlock) > 0 {
		obj.addField("ipBlock", stringArray(cfg.IPBlock))
	}

	return obj
}

func genNetworkPeerSelector(cfg *shipa.NetworkPeerSelector) *Object {
	obj := newObject()

	if cfg.MatchLabels != nil {
		obj.addObject("matchLabels", genMatchLabels(cfg.MatchLabels))
	}

	var expressions []*Object
	for _, selector := range cfg.MatchExpressions {
		expressions = append(expressions, genSelectorExpression(selector))
	}

	if expressions != nil {
		obj.addListOfObjects("matchExpressions", expressions)
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

func genFrameworkVarName(name string) string {
	return genVarName("framework", name)
}

func genFrameworkResourceName(name string) string {
	return genResourceName("framework-", name)
}

func genIndent(indentationLevel int) string {
	if indentationLevel <= 0 {
		return ""
	}

	const spacesInTab = 4
	return strings.Repeat(" ", indentationLevel*spacesInTab)
}

type Object struct {
	fields []*Field
}

type Field struct {
	name    string
	content *string
	object  *Object
	list    []*Object
}

func (obj *Object) addField(name, content string) *Object {
	obj.fields = append(obj.fields, &Field{
		name:    name,
		content: &content,
	})
	return obj
}

func (obj *Object) addObject(name string, child *Object) {
	obj.fields = append(obj.fields, &Field{
		name:   name,
		object: child,
	})
}

func (obj *Object) addListOfObjects(name string, list []*Object) {
	obj.fields = append(obj.fields, &Field{
		name: name,
		list: list,
	})
}

func (f *Field) string(indentLevel int) string {
	if f.content != nil {
		return fmt.Sprintf("%s%s: %s", genIndent(indentLevel), f.name, *f.content)
	}

	if f.object != nil {
		return fmt.Sprintf("%s%s: %s", genIndent(indentLevel), f.name, f.object.string(indentLevel))
	}

	if f.list != nil {
		var listContent []string
		for _, obj := range f.list {
			listContent = append(listContent, genIndent(indentLevel+1)+obj.string(indentLevel+1))
		}

		return fmt.Sprintf(`%s%s: [
%s
%s]`, genIndent(indentLevel), f.name, strings.Join(listContent, ",\n"), genIndent(indentLevel))
	}

	return ""
}

func (obj *Object) String() string {
	return obj.string(0)
}

func (obj *Object) string(indentLevel int) string {
	var fields []string
	for _, f := range obj.fields {
		fields = append(fields, f.string(indentLevel+1))
	}

	innerContent := strings.Join(fields, ",\n")

	return fmt.Sprintf(`{
%s
%s}`, innerContent, genIndent(indentLevel))
}

func newObject() *Object {
	return &Object{}
}
