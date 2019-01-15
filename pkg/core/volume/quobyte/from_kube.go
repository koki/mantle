package quobyte

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewQuobyteVolumeFromKubeQuobyteVolumeSource will create a new
// QuobyteVolume object with the data from a provided kubernetes
// QuobyteVolumeSource object
func NewQuobyteVolumeFromKubeQuobyteVolumeSource(obj interface{}) (*QuobyteVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.QuobyteVolumeSource{}):
		o := obj.(v1.QuobyteVolumeSource)
		return fromKubeQuobyteVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.QuobyteVolumeSource{}):
		return fromKubeQuobyteVolumeSourceV1(obj.(*v1.QuobyteVolumeSource))
	default:
		return nil, fmt.Errorf("unknown QuobyteVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeQuobyteVolumeSourceV1(vol *v1.QuobyteVolumeSource) (*QuobyteVolume, error) {
	return &QuobyteVolume{
		Registry: vol.Registry,
		Volume:   vol.Volume,
		ReadOnly: vol.ReadOnly,
		User:     vol.User,
		Group:    vol.Group,
	}, nil
}
