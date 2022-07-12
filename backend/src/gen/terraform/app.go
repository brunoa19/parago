package terraform

import (
	"fmt"
	"strings"

	"shipa-gen/src/gen/utils"
	"shipa-gen/src/shipa"

	"github.com/iancoleman/strcase"
)

func hasApp(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Team != "" && cfg.Framework != ""
}

func genDepends(dependsOn []string, genVar func(string) string) string {
	if dependsOn == nil {
		return ""
	}

	var deps []string
	for _, app := range dependsOn {
		deps = append(deps, genVar(app))
	}

	return fmt.Sprintf("depends_on: [%s]", strings.Join(deps, ", "))
}

func genAppVar(name string) string {
	return "app_" + strcase.ToSnake(name)
}

func genAppName(name string) string {
	return "shipa_app." + genAppVar(name)
}

func genApp(cfg shipa.Config) string {
	app := fmt.Sprintf(`
# Create app
resource "shipa_app" "%s" {
  app {
    name = "%s"
    teamowner = "%s"
    framework = "%s"`, genAppVar(cfg.AppName), cfg.AppName, cfg.Team, cfg.Framework)

	tags := genTags(cfg)
	if tags != "" {
		app = fmt.Sprintf(`%s
    %s`, app, tags)
	}

	app = fmt.Sprintf(`%s
  }
  %s
}
`, app, genDepends(cfg.DependsOn, genAppName))

	return app
}

func genTags(cfg shipa.Config) string {
	tags := utils.ParseValues(cfg.Tags)
	if len(tags) == 0 {
		return ""
	}

	return fmt.Sprintf(`tags = ["%s"]`, strings.Join(tags, `", "`))
}
