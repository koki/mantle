package projected

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"
	"mantle/pkg/core/pod/volume/util"

	"k8s.io/api/core/v1"
)

// NewConfigMapProjectionFromKubeConfigMapProjection will create a new
// ConfigMapProjection object with the data from a provided kubernetes
// ConfigMapProjection object
func NewConfigMapProjectionFromKubeConfigMapProjection(obj interface{}) (*ConfigMapProjection, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ConfigMapProjection{}):
		o := obj.(v1.ConfigMapProjection)
		return fromKubeConfigMapProjectionV1(&o)
	case reflect.TypeOf(&v1.ConfigMapProjection{}):
		return fromKubeConfigMapProjectionV1(obj.(*v1.ConfigMapProjection))
	default:
		return nil, fmt.Errorf("unknown ConfigMapProjection version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeConfigMapProjectionV1(vol *v1.ConfigMapProjection) (*ConfigMapProjection, error) {
	return &ConfigMapProjection{
		Name:  converterutils.FromKubeLocalObjectReferenceV1(&vol.LocalObjectReference),
		Items: util.NewKeyToPathFromKubeKeyToPathV1(vol.Items),
	}, nil
}
