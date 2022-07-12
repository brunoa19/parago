package pulumi

import (
	"fmt"
	"strings"

	"shipa-gen/src/gen/utils"
	"shipa-gen/src/shipa"
)

func genAppVarName(name string) string {
	return genVarName("app", name)
}

func genAppResourceName(name string) string {
	return genResourceName("app-", name)
}

func hasApp(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Team != "" && cfg.Framework != "" && cfg.Plan != ""
}

func genApp(cfg shipa.Config) string {
	return fmt.Sprintf(`
const %s = new shipa.App("%s", {
    app: {
%s
    }
}%s);

export const appName = app.app.name;
`, genAppVarName(cfg.AppName), genAppResourceName(cfg.AppName), genAppParams(cfg), genDepends(cfg.DependsOn, genAppVarName))
}

func genAppParams(cfg shipa.Config) string {
	const indent = "       "
	params := []string{
		fmt.Sprintf(`%s name: "%s"`, indent, cfg.AppName),
		fmt.Sprintf(`%s framework: "%s"`, indent, cfg.Framework),
		fmt.Sprintf(`%s teamowner: "%s"`, indent, cfg.Team),
		fmt.Sprintf(`%s plan: "%s"`, indent, cfg.Plan),
	}

	tags := genTags(cfg)
	if tags != "" {
		params = append(params, fmt.Sprintf(`%s %s`, indent, tags))
	}

	return strings.Join(params, ",\n")
}

func genTags(cfg shipa.Config) string {
	tags := utils.ParseValues(cfg.Tags)
	if len(tags) == 0 {
		return ""
	}

	return fmt.Sprintf(`tags: ["%s"]`, strings.Join(tags, `", "`))
}
