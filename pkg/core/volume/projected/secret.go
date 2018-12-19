package projected

import (
	"mantle/pkg/core/volume/keyandmode"
)

type SecretProjection struct {
	Name string `json:"secret"`

	Items map[string]keyandmode.KeyAndMode `json:"items,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}
