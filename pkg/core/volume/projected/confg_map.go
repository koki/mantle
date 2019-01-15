package projected

import (
	"mantle/pkg/core/volume/keyandmode"
)

type ConfigMapProjection struct {
	Name string `json:"config"`

	Items map[string]keyandmode.KeyAndMode `json:"items,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}
