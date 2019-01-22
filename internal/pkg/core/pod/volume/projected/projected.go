package projected

import (
	"mantle/internal/pkg/core/pod/volume/filemode"
)

type ProjectedVolume struct {
	Sources     []VolumeProjection `json:"sources"`
	DefaultMode *filemode.FileMode `json:"mode,omitempty"`
}
