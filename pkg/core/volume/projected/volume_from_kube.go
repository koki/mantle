package projected

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewVolumeProjectionFromKubeVolumeProjection will create a new
// ProjectedVolume object with the data from a provided kubernetes
// VolumeProjection object
func NewVolumeProjectionFromKubeVolumeProjection(obj interface{}) (*VolumeProjection, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.VolumeProjection{}):
		o := obj.(v1.VolumeProjection)
		return fromKubeVolumeProjectionV1(&o)
	case reflect.TypeOf(&v1.VolumeProjection{}):
		return fromKubeVolumeProjectionV1(obj.(*v1.VolumeProjection))
	default:
		return nil, fmt.Errorf("unknown VolumeProjection version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeVolumeProjectionV1(vol *v1.VolumeProjection) (*VolumeProjection, error) {
	var s *SecretProjection
	var dapi *DownwardAPIProjection
	var cm *ConfigMapProjection
	var err error

	if vol.Secret != nil {
		s, err = NewSecretProjectionFromKubeSecretProjection(vol.Secret)
		if err != nil {
			return nil, err
		}
	}

	if vol.ConfigMap != nil {
		cm, err = NewConfigMapProjectionFromKubeConfigMapProjection(vol.ConfigMap)
		if err != nil {
			return nil, err
		}
	}

	if vol.DownwardAPI != nil {
		dapi, err = NewDownwardAPIProjectionFromKubeDownwardAPIProjection(vol.DownardAPI)
		if err != nil {
			return nil, err
		}
	}

	return &VolumeProjection{
		Secret:      s,
		DownwardAPI: dapi,
		ConfigMap:   cm,
	}, nil
}
