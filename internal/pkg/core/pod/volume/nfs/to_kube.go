package nfs

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *NFSVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for NFSVolume: %s", version)
	}
}

func (s *NFSVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			NFS: &v1.NFSVolumeSource{
				Server:   s.Server,
				Path:     s.Path,
				ReadOnly: s.ReadOnly,
			},
		},
	}, nil
}
