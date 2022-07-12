package crossplane

import "shipa-gen/src/shipa"

type Framework struct {
	Header `yaml:",inline"`
	Spec   struct {
		ForProvider shipa.FrameworkConfig `yaml:"forProvider"`
	} `yaml:"spec"`
}
