package pulumi

import (
	"fmt"
	"strconv"
	"strings"

	"shipa-gen/src/shipa"
)

func hasAppEnv(cfg shipa.Config) bool {
	return cfg.AppName != "" && len(cfg.Envs) > 0
}

func genAppEnv(cfg shipa.Config) string {
	return fmt.Sprintf(`
const %s = new shipa.AppEnv("%s", {
    app: "%s",
    appEnv: {
        envs: [
%s
        ],
        norestart: %s,
        private: %s
    }
}%s);
`, genAppEnvVarName(cfg.AppName), genAppEnvResourceName(cfg.AppName), cfg.AppName, getEnvs(cfg), strconv.FormatBool(cfg.Norestart), strconv.FormatBool(cfg.Private), genDepends(cfg.DependsOn, genAppEnvVarName))
}

func getEnvs(cfg shipa.Config) string {
	const indent = "           "
	var envs []string
	for _, e := range cfg.Envs {
		envs = append(envs, fmt.Sprintf(`%s {name: "%s", value: "%s"},`, indent, e.Name, e.Value))
	}
	return strings.Join(envs, "\n")
}

func genAppEnvVarName(name string) string {
	return genVarName("appEnv", name)
}

func genAppEnvResourceName(name string) string {
	return genResourceName("app-env-", name)
}
