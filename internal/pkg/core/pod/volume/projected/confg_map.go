package projected

import (
	"mantle/internal/pkg/core/pod/volume/keyandmode"
)

type ConfigMapProjection struct {
	Name string `json:"config"`

	Items map[string]keyandmode.KeyAndMode `json:"items,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}
