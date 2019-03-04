package selinux

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes SELinuxOptions object of the api version provided
func (s *SELinux) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for SELinux: %s", version)
	}
}

func (s *SELinux) toKubeV1() (*v1.SELinuxOptions, error) {
	return &v1.SELinuxOptions{
		User:  s.User,
		Role:  s.Role,
		Type:  s.Type,
		Level: s.Level,
	}, nil
}
