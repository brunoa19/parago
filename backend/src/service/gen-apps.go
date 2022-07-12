package service

import (
	"errors"
	"fmt"
	"shipa-gen/src/gen/helm_chart"
	"strings"

	"shipa-gen/src/gen/ansible"
	"shipa-gen/src/gen/cloudformation"
	"shipa-gen/src/gen/crossplane"
	"shipa-gen/src/gen/github"
	"shipa-gen/src/gen/pulumi"
	"shipa-gen/src/gen/terraform"
	"shipa-gen/src/gen/topsort"
	"shipa-gen/src/models"
	"shipa-gen/src/shipa"
)

func GenerateApps(cfg shipa.AppsConfig) (*models.Payload, error) {
	out := &models.Payload{}
	var results []*shipa.Result
	for _, app := range cfg.Apps {
		app.Provider = cfg.Provider
		file, err := generateApp(app)
		if err != nil {
			out.Errors = append(out.Errors, models.Error{
				Name:  app.AppName,
				Error: err.Error(),
			})
			continue
		}

		if file != nil {
			results = append(results, file)
		}
	}

	file, err := combineResultFiles(results)
	if err != nil {
		return nil, err
	}
	out.File = file

	return out, nil
}

func combineResultFiles(results []*shipa.Result) (*models.FileData, error) {
	path, err := findPath(results)
	if err != nil {
		return nil, err
	}

	lookup := make(map[string]string)
	for _, r := range results {
		lookup[r.Name] = r.Content
	}

	f := results[0]
	var contents []string
	for _, p := range path {
		contents = append(contents, lookup[p])
	}

	separator := "\n"
	if f.Separator != "" {
		separator = f.Separator
	}

	content := strings.Join(contents, separator)
	if f.Header != "" {
		content = fmt.Sprintf("%s\n%s", f.Header, content)
	}

	return &models.FileData{
		Name:    f.Filename,
		Content: content,
	}, nil
}

func hasDependencies(results []*shipa.Result) bool {
	for _, r := range results {
		if len(r.DependsOn) > 0 {
			return true
		}
	}
	return false
}

func findPath(results []*shipa.Result) ([]string, error) {
	if len(results) == 0 {
		return nil, errors.New("empty result data")
	}

	var path []string
	if !hasDependencies(results) {
		for _, r := range results {
			path = append(path, r.Name)
		}
		return path, nil
	}

	g := topsort.NewGraph()
	for _, r := range results {
		if r.DependsOn == nil {
			g.AddNode(r.Name)
		}
		for _, dep := range r.DependsOn {
			if err := g.AddEdge(r.Name, dep); err != nil {
				return nil, err
			}
		}
	}

	for _, r := range results {
		order, err := g.TopSort(r.Name)
		if err == nil && len(order) > len(path) {
			path = order
		}
	}

	if path == nil {
		return nil, errors.New("failed put resources in order")
	}

	inPath := make(map[string]bool)
	for _, p := range path {
		inPath[p] = true
	}

	for _, r := range results {
		if inPath[r.Name] {
			continue
		}
		path = append(path, r.Name)
	}

	return path, nil
}

func generateApp(cfg shipa.Config) (*shipa.Result, error) {
	var data *shipa.Result
	switch cfg.Provider {
	case models.ProviderCrossplane:
		data = crossplane.Generate(cfg)
	case models.ProviderCloudformation:
		data = cloudformation.Generate(cfg)
	case models.ProviderGithub, models.ProviderGitlab:
		data = github.Generate(cfg)
	case models.ProviderAnsible:
		data = ansible.Generate(cfg)
	case models.ProviderTerraform:
		data = terraform.Generate(cfg)
	case models.ProviderPulumi:
		data = pulumi.Generate(cfg)
	case models.ProviderHelmChart:
		data = helm_chart.Generate(cfg)
	default:
		return nil, fmt.Errorf("not supported provider: %s", cfg.Provider)
	}

	if data == nil {
		return nil, errors.New("not data was generated")
	}

	data.Name = cfg.AppName
	data.DependsOn = cfg.DependsOn
	return data, nil
}
