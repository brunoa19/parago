package helm_chart

import (
	"gopkg.in/yaml.v2"
	"shipa-gen/src/shipa"
)

func Generate(cfg shipa.Config) *shipa.Result {

	data, _ := yaml.Marshal(cfg)

	//helmGenerator := helm.NewChartGenerator()

	return &shipa.Result{
		Filename: "helm_chart.txt",
		Content:  string(data),
	}
}
