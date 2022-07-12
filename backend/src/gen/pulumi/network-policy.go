package pulumi

import (
	"fmt"
	"strconv"
	"strings"

	"shipa-gen/src/shipa"
)

func genNetworkPolicyVarName(name string) string {
	return genVarName("netpolicy", name)
}

func genNetworkPolicyResourceName(name string) string {
	return genResourceName("netpolicy-", name)
}

func hasNetworkPolicy(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.NetworkPolicy != nil
}

func genNetworkPolicy(cfg shipa.Config) string {
	return fmt.Sprintf(`
const %s = new shipa.NetworkPolicy("%s", {
    networkPolicy: {
        app: "%s",
        networkPolicy: {
%s
        }
    }
}%s);
`, genNetworkPolicyVarName(cfg.AppName), genNetworkPolicyResourceName(cfg.AppName), cfg.AppName, getNetworkPolicy(cfg), genDepends(cfg.DependsOn, genNetworkPolicyVarName))
}

func getNetworkPolicy(cfg shipa.Config) string {
	const indent = "           "
	fields := []string{
		fmt.Sprintf("%s restartApp: %s", indent, strconv.FormatBool(cfg.NetworkPolicy.RestartApp)),
	}

	if cfg.NetworkPolicy.Ingress != nil {
		fields = append(fields, fmt.Sprintf(`%s ingress: {
%s
%s }`, indent, getNetworkPolicyConfig(cfg.NetworkPolicy.Ingress), indent))
	}

	if cfg.NetworkPolicy.Egress != nil {
		fields = append(fields, fmt.Sprintf(`%s egress: {
%s
%s }`, indent, getNetworkPolicyConfig(cfg.NetworkPolicy.Egress), indent))
	}

	return strings.Join(fields, ",\n")
}

func getNetworkPolicyConfig(config *shipa.NetworkPolicyConfig) string {
	const indent = "               "
	fields := []string{
		fmt.Sprintf(`%s policyMode: "%s"`, indent, config.PolicyMode),
	}

	if len(config.CustomRules) > 0 {
		fields = append(fields, fmt.Sprintf(`%s customRules: [
%s
%s ]
`, indent, getNetworkPolicyConfigCustomRules(config.CustomRules), indent))
	}

	return strings.Join(fields, ",\n")
}

func getNetworkPolicyConfigCustomRules(rules []*shipa.NetworkPolicyRule) string {
	var items []string
	for _, r := range rules {
		items = append(items, getNetworkPolicyConfigCustomRule(r))
	}

	return strings.Join(items, ",\n")
}

func getNetworkPolicyConfigCustomRule(rule *shipa.NetworkPolicyRule) string {
	const indent = "                       "
	fields := []string{
		fmt.Sprintf("%s enabled: %s", indent, strconv.FormatBool(rule.Enabled)),
	}

	if rule.ID != "" {
		fields = append(fields, fmt.Sprintf(`%s id: "%s"`, indent, rule.ID))
	}

	if rule.Description != "" {
		fields = append(fields, fmt.Sprintf(`%s description: "%s"`, indent, rule.Description))
	}

	if len(rule.AllowedFrameworks) > 0 {
		fields = append(fields, fmt.Sprintf(`%s allowedFrameworks: %s`, indent, genList(rule.AllowedFrameworks)))
	}

	if len(rule.AllowedApps) > 0 {
		fields = append(fields, fmt.Sprintf(`%s allowedApps: %s`, indent, genList(rule.AllowedApps)))
	}

	if len(rule.Ports) > 0 {
		fields = append(fields, fmt.Sprintf(`%s ports: [
%s
%s ]`, indent, genPorts(rule.Ports), indent))
	}

	return fmt.Sprintf(`                    {
%s
                    }`, strings.Join(fields, ",\n"))
}

func genPorts(ports []*shipa.NetworkPort) string {
	var items []string
	for _, p := range ports {
		items = append(items, genPort(p))
	}

	return strings.Join(items, ",\n")
}

func genPort(p *shipa.NetworkPort) string {
	protocol := "TCP"
	if p.Protocol != "" {
		protocol = p.Protocol
	}
	return fmt.Sprintf(`                            {
                                port: %d,
                                protocol: "%s"
                            }`, p.Port, protocol)
}

func genList(values []string) string {
	return fmt.Sprintf(`["%s"]`, strings.Join(values, `", "`))
}
