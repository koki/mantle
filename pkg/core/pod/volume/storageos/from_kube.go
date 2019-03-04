package storageos

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"

	"k8s.io/api/core/v1"
)

// NewStorageOSVolumeFromKubeStorageOSVolumeSource will create a new
// StorageOSVolume object with the data from a provided kubernetes
// StorageOSVolumeSource object
func NewStorageOSVolumeFromKubeStorageOSVolumeSource(obj interface{}) (*StorageOSVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.StorageOSVolumeSource{}):
		o := obj.(v1.StorageOSVolumeSource)
		return fromKubeStorageOSVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.StorageOSVolumeSource{}):
		return fromKubeStorageOSVolumeSourceV1(obj.(*v1.StorageOSVolumeSource))
	default:
		return nil, fmt.Errorf("unknown StorageOSVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeStorageOSVolumeSourceV1(vol *v1.StorageOSVolumeSource) (*StorageOSVolume, error) {
	return &StorageOSVolume{
		VolumeName:      vol.VolumeName,
		VolumeNamespace: vol.VolumeNamespace,
		FSType:          vol.FSType,
		ReadOnly:        vol.ReadOnly,
		SecretRef:       converterutils.FromKubeLocalObjectReferenceV1(vol.SecretRef),
	}, nil
}
