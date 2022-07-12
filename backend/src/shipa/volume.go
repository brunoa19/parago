package shipa

type VolumesConfig struct {
	Provider string         `json:"provider"`
	Volumes  []VolumeConfig `json:"volumes"`
}

type VolumeConfig struct {
	// required
	Provider string `json:"-"`
	Name     string `json:"name"`
	Capacity string `json:"capacity"`
	Plan     string `json:"plan"`

	// optional
	AccessModes string         `json:"access_modes,omitempty"`
	Opts        *VolumeOptions `json:"opts,omitempty"`

	DependsOn []string `json:"dependsOn,omitempty"`
}
