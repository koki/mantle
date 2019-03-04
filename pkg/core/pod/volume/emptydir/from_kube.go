package emptydir

import (
	"fmt"
	"reflect"


	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// NewEmptyDirVolumeFromKubeEmptyDirVolumeSource will create a new
// EmptyDirVolume object with the data from a provided kubernetes
// EmptyDirVolumeSource object
func NewEmptyDirVolumeFromKubeEmptyDirVolumeSource(obj interface{}) (*EmptyDirVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.EmptyDirVolumeSource{}):
		o := obj.(v1.EmptyDirVolumeSource)
		return fromKubeEmptyDirVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.EmptyDirVolumeSource{}):
		return fromKubeEmptyDirVolumeSourceV1(obj.(*v1.EmptyDirVolumeSource))
	default:
		return nil, fmt.Errorf("unknown EmptyDirVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeEmptyDirVolumeSourceV1(vol *v1.EmptyDirVolumeSource) (*EmptyDirVolume, error) {
	medium, err := fromKubeStorageMediumV1(vol.Medium)
	if err != nil {
		return nil, err
	}

	return &EmptyDirVolume{
		Medium:    medium,
		SizeLimit: vol.SizeLimit,
	}, nil
}

func fromKubeStorageMediumV1(medium v1.StorageMedium) (StorageMedium, error) {
	switch medium {
	case v1.StorageMediumDefault:
		return StorageMediumDefault, nil
	case v1.StorageMediumMemory:
		return StorageMediumMemory, nil
	case v1.StorageMediumHugePages:
		return StorageMediumHugePages, nil
	default:
		return StorageMediumDefault, serrors.InvalidValueErrorf(medium, "unrecognized storage medium")
	}
}
