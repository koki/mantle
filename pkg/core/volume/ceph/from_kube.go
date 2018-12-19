package ceph

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewCephFSVolumeFromKubeCephFSVolumeSource will create a new
// CephFSVolume object with the data from a provided kubernetes
// CephFSVolumeSource object
func NewCephFSVolumeFromKubeCephFSVolumeSource(obj interface{}) (*CephFSVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.CephFSVolumeSource{}):
		o := obj.(v1.CephFSVolumeSource)
		return fromKubeCephFSVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.CephFSVolumeSource{}):
		return fromKubeCephFSVolumeSourceV1(obj.(*v1.CephFSVolumeSource))
	default:
		return nil, fmt.Errorf("unknown CephFSVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeCephFSVolumeSourceV1(vol *v1.CephFSVolumeSource) (*CephFSVolume, error) {
	secretFileOrRef := createSecretFileOrRef(vol.SecretFile, vol.SecretRef)

	return &CephFSVolume{
		Monitors:        vol.Monitors,
		Path:            vol.Path,
		User:            vol.User,
		SecretFileOrRef: secretFileOrRef,
		ReadOnly:        vol.ReadOnly,
	}, nil
}

func createSecretFileOrRef(sFile string, sRef *v1.LocalObjectReference) *CephFSSecretFileOrRef {
	if len(sFile) > 0 {
		return &CephFSSecretFileOrRef{
			File: sFile,
		}
	}

	if sRef != nil {
		return &CephFSSecretFileOrRef{
			Ref: sRef.Name,
		}
	}

	return nil
}
