package crossplane

type AppEnv struct {
	ApiVersion string     `yaml:"apiVersion"`
	Kind       string     `yaml:"kind"`
	Metadata   Metadata   `yaml:"metadata"`
	Spec       AppEnvSpec `yaml:"spec"`
}

type AppEnvSpec struct {
	ForProvider AppEnvForProvider `yaml:"forProvider"`
}

type AppEnvForProvider struct {
	App    string     `yaml:"app"`
	AppEnv AppEnvData `yaml:"app_env"`
}

type AppEnvData struct {
	Envs      []Env `yaml:"envs"`
	Norestart bool  `yaml:"norestart"`
	Private   bool  `yaml:"private"`
}

type Env struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
