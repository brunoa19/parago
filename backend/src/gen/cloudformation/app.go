package cloudformation

type App struct {
	Type       string        `yaml:"Type"`
	Properties AppProperties `yaml:"Properties"`
	DependsOn  interface{}   `yaml:"DependsOn,omitempty"`
}

type AppProperties struct {
	Name       string   `yaml:"Name"`
	Teamowner  string   `yaml:"Teamowner"`
	Framework  string   `yaml:"Framework"`
	Plan       string   `yaml:"Plan"`
	Tags       []string `yaml:"Tags"`
	ShipaHost  string   `yaml:"ShipaHost"`
	ShipaToken string   `yaml:"ShipaToken"`
}
