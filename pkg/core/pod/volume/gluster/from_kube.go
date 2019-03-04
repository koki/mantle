package gluster

import (
	"fmt"
	"reflect"


	"k8s.io/api/core/v1"
)

// NewGlusterfsVolumeFromKubeGlusterfsVolumeSource will create a new
// GlusterfsVolume object with the data from a provided kubernetes
// GlusterfsVolumeSource object
func NewGlusterfsVolumeFromKubeGlusterfsVolumeSource(obj interface{}) (*GlusterfsVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.GlusterfsVolumeSource{}):
		o := obj.(v1.GlusterfsVolumeSource)
		return fromKubeGlusterfsVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.GlusterfsVolumeSource{}):
		return fromKubeGlusterfsVolumeSourceV1(obj.(*v1.GlusterfsVolumeSource))
	default:
		return nil, fmt.Errorf("unknown GlusterfsVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeGlusterfsVolumeSourceV1(vol *v1.GlusterfsVolumeSource) (*GlusterfsVolume, error) {
	return &GlusterfsVolume{
		EndpointsName: vol.EndpointsName,
		Path:          vol.Path,
		ReadOnly:      vol.ReadOnly,
	}, nil

}
