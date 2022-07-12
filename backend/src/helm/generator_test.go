package helm

import (
	"shipa-gen/src/shipa"
	"testing"
)

func Test_GetChartBuildPath(t *testing.T) {
	gen := NewChartGenerator("/tmp", "/out")
	patterns := []struct {
		cfg      shipa.Config
		expected string
	}{
		{
			cfg:      shipa.Config{AppName: "test_app"},
			expected: "/out/chart_test_app",
		},
	}
	for _, pattern := range patterns {
		genPath, _ := gen.GetChartBuildPath(&pattern.cfg)
		if genPath != pattern.expected {
			t.Errorf("Generated path %s is not as expected %s", genPath, pattern.expected)
		}
	}
}
