package cloudformation

type AppCname struct {
	Type       string             `yaml:"Type"`
	Properties AppCnameProperties `yaml:"Properties"`
	DependsOn  interface{}        `yaml:"DependsOn,omitempty"`
}

type AppCnameProperties struct {
	App        string `yaml:"App"`
	Cname      string `yaml:"Cname"`
	Encrypt    bool   `yaml:"Encrypt,omitempty"`
	ShipaHost  string `yaml:"ShipaHost"`
	ShipaToken string `yaml:"ShipaToken"`
}
