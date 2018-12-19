package secret

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"
	"mantle/pkg/core/volume/filemode"
	"mantle/pkg/core/volume/keyandmode"

	"k8s.io/api/core/v1"
)

// NewSecretVolumeFromKubeSecretVolumeSource will create a new
// SecretVolume object with the data from a provided kubernetes
// SecretVolumeSource object
func NewSecretVolumeFromKubeSecretVolumeSource(obj interface{}) (*SecretVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.SecretVolumeSource{}):
		o := obj.(v1.SecretVolumeSource)
		return fromKubeSecretVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.SecretVolumeSource{}):
		return fromKubeSecretVolumeSourceV1(obj.(*v1.SecretVolumeSource))
	default:
		return nil, fmt.Errorf("unknown SecretVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeSecretVolumeSourceV1(vol *v1.SecretVolumeSource) (*SecretVolume, error) {
	return &SecretVolume{
		SecretName:  vol.SecretName,
		Items:       keyandmode.NewKeyToPathFromKubeVKeyToPathV1(vol.Items),
		DefaultMode: filemode.NewFileModeFromKubeV1(vol.DefaultMode),
		Required:    converterutils.OptionalToRequired(vol.Optional),
	}, nil
}
