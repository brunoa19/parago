package pulumi

import (
	"fmt"
	"strconv"

	"shipa-gen/src/shipa"
)

func hasAppCname(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Cname != ""
}

func genAppCname(cfg shipa.Config) string {
	return fmt.Sprintf(`
const %s = new shipa.AppCname("%s", {
    app: "%s",
    cname: "%s",
    encrypt: %s
}%s);
`, genAppCnameVarName(cfg.AppName), genAppCnameResourceName(cfg.AppName), cfg.AppName, cfg.Cname, strconv.FormatBool(cfg.Encrypt), genDepends(cfg.DependsOn, genAppCnameVarName))
}

func genAppCnameVarName(name string) string {
	return genVarName("appCname", name)
}

func genAppCnameResourceName(name string) string {
	return genResourceName("app-cname-", name)
}
