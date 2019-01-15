package projected

import (
	"fmt"
	"reflect"

	"mantle/pkg/core/volume/downwardapi"

	"k8s.io/api/core/v1"
)

// NewDownwardAPIProjectionFromKubeDownwardAPIProjection will create a new
// DownwardAPIProjection object with the data from a provided kubernetes
// DownwardAPIProjection object
func NewDownwardAPIProjectionFromKubeDownwardAPIProjection(obj interface{}) (*DownwardAPIProjection, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.DownwardAPIProjection{}):
		o := obj.(v1.DownwardAPIProjection)
		return fromKubeDownwardAPIProjectionV1(&o)
	case reflect.TypeOf(&v1.DownwardAPIProjection{}):
		return fromKubeDownwardAPIProjectionV1(obj.(*v1.DownwardAPIProjection))
	default:
		return nil, fmt.Errorf("unknown ConfigMapProjection version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeDownwardAPIProjectionV1(vol *v1.DownwardAPIProjection) (*DownwardAPIProjection, error) {
	v, err := downwardapi.NewDownwardAPIVolumeFileFromKubeDownwardAPIVolumeFile(vol.Items)
	if err != nil {
		return nil, err
	}
	return &DownwardAPIProjection{
		Items: v,
	}, nil
}
