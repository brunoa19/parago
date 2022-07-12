package terraform

import (
	"fmt"
	"strconv"
	"strings"

	"shipa-gen/src/shipa"
)

func hasNetworkPolicy(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.NetworkPolicy != nil
}

func genNetworkPolicyName(name string) string {
	return "shipa_network_policy." + genAppVar(name)
}

func genNetworkPolicy(cfg shipa.Config) string {
	return fmt.Sprintf(`
# Set network-policy
resource "shipa_network_policy" "%s" {
  app = "%s"
  network_policy {
%s
  }
  %s
}
`, genAppVar(cfg.AppName), cfg.AppName, getNetworkPolicy(cfg), genDepends(cfg.DependsOn, genNetworkPolicyName))
}

func getNetworkPolicy(cfg shipa.Config) string {
	const indent = "   "
	fields := []string{
		fmt.Sprintf("%s restart_app = %s", indent, strconv.FormatBool(cfg.NetworkPolicy.RestartApp)),
	}

	if cfg.NetworkPolicy.Ingress != nil {
		fields = append(fields, fmt.Sprintf(`%s ingress {
%s
%s }`, indent, getNetworkPolicyConfig(cfg.NetworkPolicy.Ingress), indent))
	}

	if cfg.NetworkPolicy.Egress != nil {
		fields = append(fields, fmt.Sprintf(`%s egress {
%s
%s }`, indent, getNetworkPolicyConfig(cfg.NetworkPolicy.Egress), indent))
	}

	return strings.Join(fields, "\n")
}

func getNetworkPolicyConfig(config *shipa.NetworkPolicyConfig) string {
	const indent = "     "
	fields := []string{
		fmt.Sprintf(`%s policy_mode = "%s"`, indent, config.PolicyMode),
	}

	if len(config.CustomRules) > 0 {
		fields = append(fields, getNetworkPolicyConfigCustomRules(config.CustomRules))
	}

	return strings.Join(fields, "\n")
}

func getNetworkPolicyConfigCustomRules(rules []*shipa.NetworkPolicyRule) string {
	var items []string
	for _, r := range rules {
		items = append(items, getNetworkPolicyConfigCustomRule(r))
	}

	return strings.Join(items, ",\n")
}

func getNetworkPolicyConfigCustomRule(rule *shipa.NetworkPolicyRule) string {
	const indent = "       "
	fields := []string{
		fmt.Sprintf("%s enabled = %s", indent, strconv.FormatBool(rule.Enabled)),
	}

	if rule.ID != "" {
		fields = append(fields, fmt.Sprintf(`%s id = "%s"`, indent, rule.ID))
	}

	if rule.Description != "" {
		fields = append(fields, fmt.Sprintf(`%s description = "%s"`, indent, rule.Description))
	}

	if len(rule.AllowedFrameworks) > 0 {
		fields = append(fields, fmt.Sprintf(`%s allowed_frameworks = %s`, indent, genList(rule.AllowedFrameworks)))
	}

	if len(rule.AllowedApps) > 0 {
		fields = append(fields, fmt.Sprintf(`%s allowed_apps = %s`, indent, genList(rule.AllowedApps)))
	}

	if len(rule.Ports) > 0 {
		fields = append(fields, genPorts(rule.Ports))
	}

	return fmt.Sprintf(`      custom_rules {
%s
      }`, strings.Join(fields, "\n"))
}

func genPorts(ports []*shipa.NetworkPort) string {
	var items []string
	for _, p := range ports {
		items = append(items, genPort(p))
	}

	return strings.Join(items, "\n")
}

func genPort(p *shipa.NetworkPort) string {
	protocol := "TCP"
	if p.Protocol != "" {
		protocol = p.Protocol
	}
	return fmt.Sprintf(`        ports {
          port = %d
          protocol = "%s"
        }`, p.Port, protocol)
}

func genList(values []string) string {
	return fmt.Sprintf(`["%s"]`, strings.Join(values, `", "`))
}
