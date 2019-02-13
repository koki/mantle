package container

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewContainerFromKubeContainer will create a new Container object with
// the data from a provided kubernetes container object
func NewContainerFromKubeContainer(container interface{}) (*Container, error) {
	switch reflect.TypeOf(container) {
	case reflect.TypeOf(v1.Container{}):
		obj := container.(v1.Container)
		return fromKubeContainerV1(&obj)
	case reflect.TypeOf(&v1.Container{}):
		return fromKubeContainerV1(container.(*v1.Container))
	default:
		return nil, fmt.Errorf("unknown container version: %s", reflect.TypeOf(container))
	}
}

func fromKubeContainerV1(container *v1.Container) (*Container, error) {
	return nil, nil
}
