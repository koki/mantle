package downwardapi

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *VolumeResourceFieldSelector) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for VolumeResourceFieldSelector: %s", version)
	}
}

func (s *VolumeResourceFieldSelector) toKubeV1() (*v1.ResourceFieldSelector, error) {
	return &v1.ResourceFieldSelector{
		ContainerName: s.ContainerName,
		Resource:      s.Resource,
		Divisor:       s.Divisor,
	}, nil
}
