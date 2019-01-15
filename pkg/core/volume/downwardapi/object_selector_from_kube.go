package downwardapi

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewObjectFieldSelectorFromKubeObjectFieldSelector will create a new
// ObjectFieldSelector object with the data from a provided kubernetes
// ObjectFieldSelector object
func NewObjectFieldSelectorFromKubeObjectFieldSelector(obj interface{}) (*ObjectFieldSelector, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ObjectFieldSelector{}):
		o := obj.(v1.ObjectFieldSelector)
		return fromKubeObjectFieldSelectorV1(&o)
	case reflect.TypeOf(&v1.ObjectFieldSelector{}):
		return fromKubeObjectFieldSelectorV1(obj.(*v1.ObjectFieldSelector))
	default:
		return nil, fmt.Errorf("unknown ObjectFieldSelector version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeObjectFieldSelectorV1(vol *v1.ObjectFieldSelector) (*ObjectFieldSelector, error) {
	if vol == nil {
		return nil, nil
	}

	return &ObjectFieldSelector{
		FieldPath:  vol.FieldPath,
		APIVersion: vol.APIVersion,
	}, nil
}
