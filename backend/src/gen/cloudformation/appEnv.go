package cloudformation

type AppEnv struct {
	Type       string           `yaml:"Type"`
	Properties AppEnvProperties `yaml:"Properties"`
	DependsOn  interface{}      `yaml:"DependsOn,omitempty"`
}

type AppEnvProperties struct {
	App        string `yaml:"App"`
	Envs       []Env  `yaml:"Envs"`
	Norestart  bool   `yaml:"Norestart"`
	Private    bool   `yaml:"Private"`
	ShipaHost  string `yaml:"ShipaHost"`
	ShipaToken string `yaml:"ShipaToken"`
}

type Env struct {
	Name  string `yaml:"Name"`
	Value string `yaml:"Value"`
}
