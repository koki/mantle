package downwardapi

import (
	"mantle/pkg/core/pod/volume/util"

	"k8s.io/apimachinery/pkg/api/resource"
)

type DownwardAPIVolume struct {
	Items       map[string]DownwardAPIVolumeFile `json:"items,omitempty"`
	DefaultMode *util.FileMode                   `json:"mode,omitempty"`
}

type DownwardAPIVolumeFile struct {
	FieldRef         *ObjectFieldSelector         `json:"field,omitempty"`
	ResourceFieldRef *VolumeResourceFieldSelector `json:"resource,omitempty"`
	Mode             *util.FileMode               `json:"mode,omitempty"`
}

type ObjectFieldSelector struct {
	// required
	FieldPath string `json:"-"`

	// optional
	APIVersion string `json:"-"`
}

type VolumeResourceFieldSelector struct {
	// required
	ContainerName string `json:"-"`

	// required
	Resource string `json:"-"`

	// optional
	Divisor resource.Quantity `json:"-"`
}
