package configmap

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"
	"mantle/internal/pkg/core/pod/volume/filemode"
	"mantle/internal/pkg/core/pod/volume/keyandmode"

	"k8s.io/api/core/v1"
)

// NewConfigMapVolumeFromKubeConfigMapVolumeSource will create a new
// ConfigMapVolume object with the data from a provided kubernetes
// ConfigMapVolumeSource object
func NewConfigMapVolumeFromKubeConfigMapVolumeSource(obj interface{}) (*ConfigMapVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ConfigMapVolumeSource{}):
		o := obj.(v1.ConfigMapVolumeSource)
		return fromKubeConfigMapVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.ConfigMapVolumeSource{}):
		return fromKubeConfigMapVolumeSourceV1(obj.(*v1.ConfigMapVolumeSource))
	default:
		return nil, fmt.Errorf("unknown ConfigMapVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeConfigMapVolumeSourceV1(vol *v1.ConfigMapVolumeSource) (*ConfigMapVolume, error) {
	return &ConfigMapVolume{
		Name:        converterutils.FromKubeLocalObjectReferenceV1(&vol.LocalObjectReference),
		Items:       keyandmode.NewKeyToPathFromKubeKeyToPathV1(vol.Items),
		DefaultMode: filemode.NewFileModeFromKubeV1(vol.DefaultMode),
		Required:    converterutils.OptionalToRequired(vol.Optional),
	}, nil
}
