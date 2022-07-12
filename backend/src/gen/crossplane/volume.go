package crossplane

type Header struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
}

type Volume struct {
	Header `yaml:",inline"`
	Spec   struct {
		ForProvider VolumeSpec `yaml:"forProvider"`
	} `yaml:"spec"`
}

type VolumeSpec struct {
	// required
	Name     string `yaml:"name"`
	Capacity string `yaml:"capacity"`
	Plan     string `yaml:"plan"`
	// optional
	AccessModes string         `yaml:"accessModes,omitempty"` // default: ReadWriteOnce
	Opts        *VolumeOptions `yaml:"opts,omitempty"`
}
