package terraform

import (
	"fmt"
	"strconv"

	"shipa-gen/src/shipa"
)

func hasAppCname(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Cname != ""
}

func genAppCnameName(name string) string {
	return "shipa_app_cname." + genAppVar(name)
}

func genAppCname(cfg shipa.Config) string {
	return fmt.Sprintf(`
# Set app cname
resource "shipa_app_cname" "%s" {
  app = %s
  cname = "%s"
  encrypt = %s
  %s
}
`, genAppVar(cfg.AppName), getAppName(cfg), cfg.Cname, strconv.FormatBool(cfg.Encrypt), genDepends(cfg.DependsOn, genAppCnameName))
}

func getAppName(cfg shipa.Config) string {
	return fmt.Sprintf(`"%s"`, cfg.AppName)
}
