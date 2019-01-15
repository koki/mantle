package projected

import (
	"mantle/pkg/core/volume/downwardapi"
)

type DownwardAPIProjection struct {
	Items map[string]downwardapi.DownwardAPIVolumeFile `json:"items,omitempty"`
}
