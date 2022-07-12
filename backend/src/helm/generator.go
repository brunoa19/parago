package helm

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"shipa-gen/src/shipa"
	"sigs.k8s.io/yaml"
	"text/template"
)

type ChartGenerator struct {
	templatesBasePath  string
	generationBasePath string
}

func NewChartGenerator(templatesPath string, generationPath string) ChartGeneratorInterface {
	return &ChartGenerator{
		templatesBasePath:  templatesPath,
		generationBasePath: generationPath,
	}
}

func (g *ChartGenerator) PrepareChart(cfg *shipa.Config) (string, error) {
	// clear destination path if exists
	dstPath, err := g.GetChartBuildPath(cfg)
	if err != nil {
		return "", err
	}
	// clear potentially generated chart
	err = os.RemoveAll(dstPath)
	if err != nil {
		return "", err
	}
	dstTemplatesPath := dstPath + "/templates"
	err = os.MkdirAll(dstTemplatesPath, os.FileMode(0775))
	if err != nil {
		return "", err
	}
	srcPathBase := "backend/src/helm/template/"
	srcTemplateChunks := []string{"common"}
	// copy template files from chunks
	for _, chunkName := range srcTemplateChunks {
		srcSource := srcPathBase + chunkName + "/yamls"
		files, err := ioutil.ReadDir(srcSource)
		if err != nil {
			return "", err
		}
		for _, fileInfo := range files {
			srcFilePath := srcSource + "/" + fileInfo.Name()
			dstFilePath := dstTemplatesPath + "/" + fileInfo.Name()
			log.Printf("copy %s -> %s", srcFilePath, dstFilePath)
			_, err = copyFile(srcFilePath, dstFilePath)
			if err != nil {
				return "", err
			}
		}
	}

	// generate chart file
	if err = generateChartFile(dstPath+"/Chart.yaml", cfg); err != nil {
		return "", err
	}

	// generate values file
	if err = generateValuesFile(dstPath+"/values.yaml", cfg); err != nil {
		return "", err
	}

	return dstPath, nil
}

func (g *ChartGenerator) BuildChart(chartPath string) (string, error) {
	// build chart file
	// return chart path
	return "", nil
}

func (g *ChartGenerator) GetChartBuildPath(cfg *shipa.Config) (string, error) {
	appName := getAppName(cfg)

	return g.generationBasePath + "/chart_" + appName, nil
}

func getAppName(cfg *shipa.Config) string {
	if cfg == nil {
		return ""
	}

	return url.PathEscape(cfg.AppName)
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	return io.Copy(destination, source)
}

const (
	chartYaml = `apiVersion: v2
name: {{ .AppName }}
description: {{ .Description }}
type: application
version: {{ .Version }}
{{- if .AppVersion }}
appVersion: {{ .AppVersion }}
{{- end }}
`
)

func generateChartFile(dstPath string, cfg *shipa.Config) error {
	buf := bytes.Buffer{}
	t := template.Must(template.New("chart.yaml").Parse(chartYaml))
	err := t.Execute(&buf, map[string]interface{}{
		"AppName":     cfg.AppName,
		"Description": cfg.Cname,
		"Version":     "0.0.1",
	})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dstPath, buf.Bytes(), 0644)
}

func generateValuesFile(dstPath string, cfg *shipa.Config) error {
	values := values{
		App: &app{
			Name: cfg.AppName,
			Type: DeploymentAppType,
		},
		IngressController: nil,
	}
	valuesBytes, err := yaml.Marshal(values)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dstPath, valuesBytes, 0644)
}
