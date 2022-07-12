package crossplane

type App struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       AppSpec  `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type AppSpec struct {
	ForProvider AppForProvider `yaml:"forProvider"`
}

type AppForProvider struct {
	Name      string   `yaml:"name"`
	Framework string   `yaml:"framework"`
	TeamOwner string   `yaml:"teamOwner"`
	Plan      string   `yaml:"plan"`
	Tags      []string `yaml:"tags"`
}
