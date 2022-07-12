package crossplane

type AppCname struct {
	ApiVersion string       `yaml:"apiVersion"`
	Kind       string       `yaml:"kind"`
	Metadata   Metadata     `yaml:"metadata"`
	Spec       AppCnameSpec `yaml:"spec"`
}

type AppCnameSpec struct {
	ForProvider AppCnameForProvider `yaml:"forProvider"`
}

type AppCnameForProvider struct {
	App     string `yaml:"app"`
	Cname   string `yaml:"cname"`
	Encrypt bool   `yaml:"encrypt,omitempty"`
}
