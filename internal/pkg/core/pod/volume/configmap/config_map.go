package configmap

import (
	"mantle/internal/pkg/core/pod/volume/filemode"
	"mantle/internal/pkg/core/pod/volume/keyandmode"
)

type ConfigMapVolume struct {
	Name string `json:"-"`

	Items       map[string]keyandmode.KeyAndMode `json:"items,omitempty"`
	DefaultMode *filemode.FileMode               `json:"mode,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}
