package portworx

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewPortworxVolumeVolumeFromKubePortworxVolumeSource will create a new
// PortworxVolumeVolume object with the data from a provided kubernetes
// PortworxVolumeSource object
func NewPortworxVolumeVolumeFromKubePortworxVolumeSource(obj interface{}) (*PortworxVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.PortworxVolumeSource{}):
		o := obj.(v1.PortworxVolumeSource)
		return fromKubePortworxVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.PortworxVolumeSource{}):
		return fromKubePortworxVolumeSourceV1(obj.(*v1.PortworxVolumeSource))
	default:
		return nil, fmt.Errorf("unknown PortworxVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubePortworxVolumeSourceV1(vol *v1.PortworxVolumeSource) (*PortworxVolume, error) {
	return &PortworxVolume{
		VolumeID: vol.VolumeID,
		FSType:   vol.FSType,
		ReadOnly: vol.ReadOnly,
	}, nil
}
