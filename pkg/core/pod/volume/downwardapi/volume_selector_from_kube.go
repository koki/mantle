package downwardapi

import (
	"fmt"
	"reflect"


	"k8s.io/api/core/v1"
)

// NewVolumeResourceFieldSelectorFromKubeResourceFieldSelector will create a new
// VolumeResourceFieldSelector object with the data from a provided kubernetes
// ResourceFieldSelector object
func NewVolumeResourceFieldSelectorFromKubeResourceFieldSelector(obj interface{}) (*VolumeResourceFieldSelector, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ResourceFieldSelector{}):
		o := obj.(v1.ResourceFieldSelector)
		return fromKubeResourceFieldSelectorV1(&o)
	case reflect.TypeOf(&v1.ResourceFieldSelector{}):
		return fromKubeResourceFieldSelectorV1(obj.(*v1.ResourceFieldSelector))
	default:
		return nil, fmt.Errorf("unknown ResourceFieldSelector version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeResourceFieldSelectorV1(vol *v1.ResourceFieldSelector) (*VolumeResourceFieldSelector, error) {
	if vol == nil {
		return nil, nil
	}

	return &VolumeResourceFieldSelector{
		ContainerName: vol.ContainerName,
		Resource:      vol.Resource,
		Divisor:       vol.Divisor,
	}, nil
}
