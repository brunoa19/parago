package helm

import "shipa-gen/src/shipa"

type ChartGeneratorInterface interface {
	GetChartBuildPath(cfg *shipa.Config) (string, error)
	PrepareChart(cfg *shipa.Config) (string, error)
	BuildChart(chartPath string) (string, error)
}
