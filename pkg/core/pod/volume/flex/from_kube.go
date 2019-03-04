package flex

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"

	"k8s.io/api/core/v1"
)

// NewFlexVolumeFromKubeFlexVolumeSource will create a new
// FlexVolume object with the data from a provided kubernetes
// FlexVolumeSource object
func NewFlexVolumeFromKubeFlexVolumeSource(obj interface{}) (*FlexVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.FlexVolumeSource{}):
		o := obj.(v1.FlexVolumeSource)
		return fromKubeFlexVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.FlexVolumeSource{}):
		return fromKubeFlexVolumeSourceV1(obj.(*v1.FlexVolumeSource))
	default:
		return nil, fmt.Errorf("unknown FlexVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeFlexVolumeSourceV1(vol *v1.FlexVolumeSource) (*FlexVolume, error) {
	return &FlexVolume{
		Driver:    vol.Driver,
		FSType:    vol.FSType,
		SecretRef: converterutils.FromKubeLocalObjectReferenceV1(vol.SecretRef),
		ReadOnly:  vol.ReadOnly,
		Options:   vol.Options,
	}, nil
}
