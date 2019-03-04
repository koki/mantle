package hostalias

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewHostAliasFromKubeHostAlias will create a new HostAlias object with
// the data from a provided kubernetes HostAlias object
func NewHostAliasFromKubeHostAlias(obj interface{}) (*HostAlias, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.HostAlias{}):
		return fromKubeHostAliasV1(obj.(v1.HostAlias))
	case reflect.TypeOf(&v1.HostAlias{}):
		o := obj.(*v1.HostAlias)
		return fromKubeHostAliasV1(*o)
	default:
		return nil, fmt.Errorf("unknown HostAlias version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeHostAliasV1(alias v1.HostAlias) (*HostAlias, error) {
	return &HostAlias{
		IP:        alias.IP,
		Hostnames: alias.Hostnames,
	}, nil
}
