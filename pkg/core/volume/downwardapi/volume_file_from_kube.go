package downwardapi

import (
	"fmt"
	"reflect"

	"mantle/pkg/core/volume/filemode"

	"k8s.io/api/core/v1"
)

// NewDownwardAPIVolumeFileFromKubeDownwardAPIVolumeFile will create a new
// DownwardAPIVolumeFile object with the data from a provided kubernetes
// DownwardAPIVolumeFile object
func NewDownwardAPIVolumeFileFromKubeDownwardAPIVolumeFile(obj interface{}) (*DownwardAPIVolumeFile, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.DownwardAPIVolumeFile{}):
		o := obj.(v1.DownwardAPIVolumeFile)
		return fromKubeDownwardAPIVolumeFileV1(&o)
	case reflect.TypeOf(&v1.DownwardAPIVolumeFile{}):
		return fromKubeDownwardAPIVolumeFileV1(obj.(*v1.DownwardAPIVolumeFile))
	default:
		return nil, fmt.Errorf("unknown DownwardAPIVolumeFile version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeDownwardAPIVolumeFileV1(vol *v1.DownwardAPIVolumeFile) (*DownwardAPIVolumeFile, error) {
	objRef, err := NewObjectFieldSelectorFromKubeObjectFieldSelector(vol.FieldRef)
	if err != nil {
		return nil, err
	}

	volRef, err := NewVolumeResourceFieldSelectorFromKubeResourceFieldSelector(vol.ResourceFieldRef)
	if err != nil {
		return nil, err
	}

	return &DownwardAPIVolumeFile{
		FieldRef:         objRef,
		ResourceFieldRef: volRef,
		Mode:             filemode.NewFileModeFromKubeV1(vol.Mode),
	}, nil
}
