package azure

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewAzureFileVolumeFromAzureFileVolumeSource will create a new
// AzureFileVolume object with the data from a provided kubernetes
// AzureFileVolumeSource object
func NewAzureFileVolumeFromAzureFileVolumeSource(obj interface{}) (*AzureFileVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.AzureFileVolumeSource{}):
		o := obj.(v1.AzureFileVolumeSource)
		return fromKubeAzureFileVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.AzureFileVolumeSource{}):
		return fromKubeAzureFileVolumeSourceV1(obj.(*v1.AzureFileVolumeSource))
	default:
		return nil, fmt.Errorf("unknown AzureFileVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeAzureFileVolumeSourceV1(vol *v1.AzureFileVolumeSource) (*AzureFileVolume, error) {
	return &AzureFileVolume{
		SecretName: vol.SecretName,
		ShareName:  vol.ShareName,
		ReadOnly:   vol.ReadOnly,
	}, nil
}
