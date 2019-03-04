package secret

import (
	"mantle/pkg/core/pod/volume/util"
)

type SecretVolume struct {
	SecretName string `json:"-"`

	Items       map[string]util.KeyAndMode `json:"items,omitempty"`
	DefaultMode *util.FileMode             `json:"mode,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}
