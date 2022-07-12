package pulumi

import (
	"shipa-gen/src/shipa"
)

func Generate(cfg shipa.Config) *shipa.Result {
	header := genMain()

	var content string

	if hasAppDeploy(cfg) {
		content += genAppDeploy(cfg)
	} else {
		if hasApp(cfg) {
			content += genApp(cfg)
		}

		if hasAppEnv(cfg) {
			content += genAppEnv(cfg)
		}
	}

	if hasAppCname(cfg) {
		content += genAppCname(cfg)
	}

	if hasNetworkPolicy(cfg) {
		content += genNetworkPolicy(cfg)
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
