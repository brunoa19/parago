package terraform

import (
	"fmt"
	"strconv"
	"strings"

	"shipa-gen/src/shipa"
)

func hasAppEnv(cfg shipa.Config) bool {
	return cfg.AppName != "" && len(cfg.Envs) > 0
}

func genAppEnvName(name string) string {
	return "shipa_app_env." + genAppVar(name)
}

func genAppEnv(cfg shipa.Config) string {
	return fmt.Sprintf(`
# Set app envs
resource "shipa_app_env" "%s" {
  app = %s
  app_env {
%s
   norestart = %s
   private = %s
  }
  %s
}
`, genAppVar(cfg.AppName), getAppName(cfg), getEnvs(cfg), strconv.FormatBool(cfg.Norestart), strconv.FormatBool(cfg.Private), genDepends(cfg.DependsOn, genAppEnvName))
}

func getEnvs(cfg shipa.Config) string {
	var envs []string
	for _, e := range cfg.Envs {
		envs = append(envs, fmt.Sprintf(`   envs {
     name = "%s"
     value = "%s"
   }`, e.Name, e.Value))
	}
	return strings.Join(envs, "\n")
}
