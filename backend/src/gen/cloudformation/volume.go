package cloudformation

type Volume struct {
	Type       string           `yaml:"Type"`
	Properties VolumeProperties `yaml:"Properties"`
	DependsOn  interface{}      `yaml:"DependsOn,omitempty"`
}

type Shipa struct {
	ShipaHost  string `yaml:"ShipaHost"`
	ShipaToken string `yaml:"ShipaToken"`
}

type VolumeProperties struct {
	Shipa       `yaml:",inline"`
	Name        string         `yaml:"Name"`
	Capacity    string         `yaml:"Capacity"`
	Plan        string         `yaml:"Plan"`
	AccessModes string         `yaml:"AccessModes,omitempty"`
	Opts        *VolumeOptions `yaml:"Opts,omitempty"`
}
