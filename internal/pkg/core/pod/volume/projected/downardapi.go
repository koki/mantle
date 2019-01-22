package projected

import (
	"mantle/internal/pkg/core/pod/volume/downwardapi"
)

type DownwardAPIProjection struct {
	Items map[string]downwardapi.DownwardAPIVolumeFile `json:"items,omitempty"`
}
