package volumemount

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewVolumeMountFromKubeVolumeMount will create a new
// VolumeMount object with the data from a provided kubernetes
// VolumeMount object
func NewVolumeMountFromKubeVolumeMount(obj interface{}) (*VolumeMount, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.VolumeMount{}):
		return fromKubeVolumeMountV1(obj.(v1.VolumeMount))
	case reflect.TypeOf(&v1.VolumeMount{}):
		o := obj.(*v1.VolumeMount)
		return fromKubeVolumeMountV1(*o)
	default:
		return nil, fmt.Errorf("unknown VolumeMount version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeVolumeMountV1(mount v1.VolumeMount) (*VolumeMount, error) {
	m := VolumeMount{
		MountPath: mount.MountPath,
		ReadOnly:  mount.ReadOnly,
	}

	if mount.MountPropagation != nil {
		var mode MountPropagation

		switch *mount.MountPropagation {
		case v1.MountPropagationHostToContainer:
			mode = MountPropagationHostToContainer

		case v1.MountPropagationBidirectional:
			mode = MountPropagationBidirectional

		case v1.MountPropagationNone:
			mode = MountPropagationNone

		default:
			mode = MountPropagationDefault
		}

		m.Propagation = &mode
	}

	store := mount.Name
	if len(mount.SubPath) > 0 {
		store += fmt.Sprintf(":%s", mount.SubPath)
	}
	m.Store = store

	return &m, nil
}
