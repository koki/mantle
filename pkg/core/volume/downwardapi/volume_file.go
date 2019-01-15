package downwardapi

import (
	"mantle/pkg/core/volume/filemode"
)

type DownwardAPIVolumeFile struct {
	FieldRef         *ObjectFieldSelector         `json:"field,omitempty"`
	ResourceFieldRef *VolumeResourceFieldSelector `json:"resource,omitempty"`
	Mode             *filemode.FileMode           `json:"mode,omitempty"`
}
