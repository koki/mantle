package nfs

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewNFSVolumeFromKubeNFSVolumeSource will create a new
// NFSVolume object with the data from a provided kubernetes
// NFSVolumeSource object
func NewNFSVolumeFromNFSVolumeSource(obj interface{}) (*NFSVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.NFSVolumeSource{}):
		o := obj.(v1.NFSVolumeSource)
		return fromKubeNFSVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.NFSVolumeSource{}):
		return fromKubeNFSVolumeSourceV1(obj.(*v1.NFSVolumeSource))
	default:
		return nil, fmt.Errorf("unknown NFSVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeNFSVolumeSourceV1(vol *v1.NFSVolumeSource) (*NFSVolume, error) {
	return &NFSVolume{
		Server:   vol.Server,
		Path:     vol.Path,
		ReadOnly: vol.ReadOnly,
	}, nil
}
