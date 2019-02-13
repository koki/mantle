package fibrechannel

import (
	"fmt"
	"reflect"


	"k8s.io/api/core/v1"
)

// NewFibreChannelVolumeFromKubeFCVolumeSource will create a new
// FibreChannelVolume object with the data from a provided kubernetes
// FCVolumeSource object
func NewFibreChannelVolumeFromKubeFCVolumeSource(obj interface{}) (*FibreChannelVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.FCVolumeSource{}):
		o := obj.(v1.FCVolumeSource)
		return fromKubeFCVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.FCVolumeSource{}):
		return fromKubeFCVolumeSourceV1(obj.(*v1.FCVolumeSource))
	default:
		return nil, fmt.Errorf("unknown FCVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeFCVolumeSourceV1(vol *v1.FCVolumeSource) (*FibreChannelVolume, error) {
	return &FibreChannelVolume{
		TargetWWNs: vol.TargetWWNs,
		Lun:        vol.Lun,
		ReadOnly:   vol.ReadOnly,
		WWIDs:      vol.WWIDs,
		FSType:     vol.FSType,
	}, nil
}
