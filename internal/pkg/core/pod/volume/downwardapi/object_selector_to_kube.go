package downwardapi

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ObjectFieldSelector) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for ObjectFieldSelector: %s", version)
	}
}

func (s *ObjectFieldSelector) toKubeV1() (*v1.ObjectFieldSelector, error) {
	return &v1.ObjectFieldSelector{
		FieldPath:  s.FieldPath,
		APIVersion: s.APIVersion,
	}, nil
}
