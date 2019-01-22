package downwardapi

import (
	"mantle/internal/pkg/core/pod/volume/filemode"
)

type DownwardAPIVolume struct {
	Items       map[string]DownwardAPIVolumeFile `json:"items,omitempty"`
	DefaultMode *filemode.FileMode               `json:"mode,omitempty"`
}
