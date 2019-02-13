package projected

import (
	"mantle/pkg/core/pod/volume/downwardapi"
	"mantle/pkg/core/pod/volume/util"
)

type ProjectedVolume struct {
	Sources     []VolumeProjection `json:"sources"`
	DefaultMode *util.FileMode     `json:"mode,omitempty"`
}

type VolumeProjection struct {
	Secret      *SecretProjection      `json:"-"`
	DownwardAPI *DownwardAPIProjection `json:"-"`
	ConfigMap   *ConfigMapProjection   `json:"-"`
}

type SecretProjection struct {
	Name string `json:"secret"`

	Items map[string]util.KeyAndMode `json:"items,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}

type DownwardAPIProjection struct {
	Items map[string]downwardapi.DownwardAPIVolumeFile `json:"items,omitempty"`
}

type ConfigMapProjection struct {
	Name string `json:"config"`

	Items map[string]util.KeyAndMode `json:"items,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}
