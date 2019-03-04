package quobyte

import (
	"fmt"
	"strings"


	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *QuobyteVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for QuobyteVolume: %s", version)
	}
}

func (s *QuobyteVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			Quobyte: &v1.QuobyteVolumeSource{
				Registry: s.Registry,
				Volume:   s.Volume,
				ReadOnly: s.ReadOnly,
				User:     s.User,
				Group:    s.Group,
			},
		},
	}, nil
}
