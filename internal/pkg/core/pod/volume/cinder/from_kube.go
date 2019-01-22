package cinder

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewCinderVolumeFromKubeCinderVolumeSource will create a new
// CinderVolume object with the data from a provided kubernetes
// CinderVolumeSource object
func NewCinderVolumeFromKubeCinderVolumeSource(obj interface{}) (*CinderVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.CinderVolumeSource{}):
		o := obj.(v1.CinderVolumeSource)
		return fromKubeCinderVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.CinderVolumeSource{}):
		return fromKubeCinderVolumeSourceV1(obj.(*v1.CinderVolumeSource))
	default:
		return nil, fmt.Errorf("unknown CinderVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeCinderVolumeSourceV1(vol *v1.CinderVolumeSource) (*CinderVolume, error) {
	return &CinderVolume{
		VolumeID: vol.VolumeID,
		FSType:   vol.FSType,
		ReadOnly: vol.ReadOnly,
	}, nil
}
