package rbd

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"

	"k8s.io/api/core/v1"
)

// NewRBDVolumeFromKubeRBDVolumeSource will create a new
// RBDVolume object with the data from a provided kubernetes
// RBDVolumeSource object
func NewRBDVolumeFromKubeRBDVolumeSource(obj interface{}) (*RBDVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.RBDVolumeSource{}):
		o := obj.(v1.RBDVolumeSource)
		return fromKubeRBDVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.RBDVolumeSource{}):
		return fromKubeRBDVolumeSourceV1(obj.(*v1.RBDVolumeSource))
	default:
		return nil, fmt.Errorf("unknown RBDVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeRBDVolumeSourceV1(vol *v1.RBDVolumeSource) (*RBDVolume, error) {
	return &RBDVolume{
		CephMonitors: vol.CephMonitors,
		RBDImage:     vol.RBDImage,
		FSType:       vol.FSType,
		RBDPool:      vol.RBDPool,
		RadosUser:    vol.RadosUser,
		Keyring:      vol.Keyring,
		SecretRef:    converterutils.FromKubeLocalObjectReferenceV1(vol.SecretRef),
		ReadOnly:     vol.ReadOnly,
	}, nil
}
