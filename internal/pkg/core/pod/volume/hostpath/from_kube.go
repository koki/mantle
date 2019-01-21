package hostpath

import (
	"fmt"
	"reflect"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// NewHostPathVolumeFromKubeHostPathVolumeSource will create a new
// HostPathVolume object with the data from a provided kubernetes
// HostPathVolumeSource object
func NewHostPathVolumeFromKubeHostPathVolumeSource(obj interface{}) (*HostPathVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.HostPathVolumeSource{}):
		o := obj.(v1.HostPathVolumeSource)
		return fromKubeHostPathVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.HostPathVolumeSource{}):
		return fromKubeHostPathVolumeSourceV1(obj.(*v1.HostPathVolumeSource))
	default:
		return nil, fmt.Errorf("unknown HostPathVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeHostPathVolumeSourceV1(vol *v1.HostPathVolumeSource) (*HostPathVolume, error) {
	kokiType, err := fromKubeHostPathTypeV1(vol.Type)
	if err != nil {
		return nil, err
	}

	return &HostPathVolume{
		Path: vol.Path,
		Type: kokiType,
	}, nil
}

func fromKubeHostPathTypeV1(kubeType *v1.HostPathType) (HostPathType, error) {
	if kubeType == nil {
		return HostPathUnset, nil
	}

	switch *kubeType {
	case v1.HostPathUnset:
		return HostPathUnset, nil
	case v1.HostPathDirectoryOrCreate:
		return HostPathDirectoryOrCreate, nil
	case v1.HostPathDirectory:
		return HostPathDirectory, nil
	case v1.HostPathFileOrCreate:
		return HostPathFileOrCreate, nil
	case v1.HostPathFile:
		return HostPathFile, nil
	case v1.HostPathSocket:
		return HostPathSocket, nil
	case v1.HostPathCharDev:
		return HostPathCharDev, nil
	case v1.HostPathBlockDev:
		return HostPathBlockDev, nil
	default:
		return HostPathUnset, serrors.InvalidValueErrorf(kubeType, "unrecognized host_path type")
	}
}
