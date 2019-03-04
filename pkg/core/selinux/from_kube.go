package selinux

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NeNewSELinuxFromKubeSELinuxOptions will create a new
// SELinux object with the data from a provided kubernetes
// SELinuxOptions object
func NewSELinuxFromKubeSELinuxOptions(obj interface{}) (*SELinux, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.SELinuxOptions{}):
		o := obj.(v1.SELinuxOptions)
		return fromSELinuxOptionsV1(&o)
	case reflect.TypeOf(&v1.SELinuxOptions{}):
		return fromSELinuxOptionsV1(obj.(*v1.SELinuxOptions))
	default:
		return nil, fmt.Errorf("unknown SELinuxOptions version: %s", reflect.TypeOf(obj))
	}
}

func fromSELinuxOptionsV1(opts *v1.SELinuxOptions) (*SELinux, error) {
	return &SELinux{
		User:  opts.User,
		Level: opts.Level,
		Role:  opts.Role,
		Type:  opts.Type,
	}, nil
}
