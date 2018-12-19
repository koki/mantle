package pvc

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewPVCVolumeFromKubePersistentVolumeClaimVolumeSource will create a new
// PVCVolume object with the data from a provided kubernetes
// PersistentVolumeClaimVolumeSource object
func NewPVCVolumeFromKubePersistentVolumeClaimVolumeSource(obj interface{}) (*PVCVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.PersistentVolumeClaimVolumeSource{}):
		o := obj.(v1.PersistentVolumeClaimVolumeSource)
		return fromKubePersistentVolumeClaimVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.PersistentVolumeClaimVolumeSource{}):
		return fromKubePersistentVolumeClaimVolumeSourceV1(obj.(*v1.PersistentVolumeClaimVolumeSource))
	default:
		return nil, fmt.Errorf("unknown PersistentVolumeClaimVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubePersistentVolumeClaimVolumeSourceV1(vol *v1.PersistentVolumeClaimVolumeSource) (*PVCVolume, error) {
	return &PVCVolume{
		ClaimName: vol.ClaimName,
		ReadOnly:  vol.ReadOnly,
	}, nil
}
