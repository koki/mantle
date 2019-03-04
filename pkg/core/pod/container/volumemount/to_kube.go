package volumemount

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes VolumeMount object of the api version provided
func (vm *VolumeMount) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return vm.toKubeV1()
	case "":
		return vm.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for VolumeMount: %s", version)
	}
}

func (vm *VolumeMount) toKubeV1() (*v1.VolumeMount, error) {
	kubeMount := v1.VolumeMount{}

	if vm.Propagation != nil {
		var mode v1.MountPropagationMode

		switch *vm.Propagation {
		case MountPropagationHostToContainer:
			mode = v1.MountPropagationHostToContainer

		case MountPropagationBidirectional:
			mode = v1.MountPropagationBidirectional

		case MountPropagationNone:
			mode = v1.MountPropagationNone

		default:
			mode = ""
		}

		kubeMount.MountPropagation = &mode
	}

	kubeMount.MountPath = vm.MountPath

	fields := strings.SplitN(vm.Store, ":", 2)
	if len(fields) == 1 {
		kubeMount.Name = vm.Store
	} else {
		kubeMount.Name = fields[0]
		kubeMount.SubPath = fields[1]
	}
	kubeMount.ReadOnly = vm.ReadOnly

	return &kubeMount, nil
}
