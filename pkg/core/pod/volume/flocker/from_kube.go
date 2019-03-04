package flocker

import (
	"fmt"
	"reflect"


	"k8s.io/api/core/v1"
)

// NewFlockerVolumeFromKubeFlockerVolumeSource will create a new
// FlockerVolume object with the data from a provided kubernetes
// FlockerVolumeSource object
func NewFlockerVolumeFromKubeFlockerVolumeSource(obj interface{}) (*FlockerVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.FlockerVolumeSource{}):
		o := obj.(v1.FlockerVolumeSource)
		return fromKubeFlockerVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.FlockerVolumeSource{}):
		return fromKubeFlockerVolumeSourceV1(obj.(*v1.FlockerVolumeSource))
	default:
		return nil, fmt.Errorf("unknown FlockerVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeFlockerVolumeSourceV1(vol *v1.FlockerVolumeSource) (*FlockerVolume, error) {
	var dataset string

	if len(vol.DatasetUUID) > 0 {
		dataset = vol.DatasetUUID
	} else {
		dataset = vol.DatasetName
	}
	return &FlockerVolume{
		DatasetUUID: dataset,
	}, nil
}
