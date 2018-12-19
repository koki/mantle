package portworx

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *PortworxVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for PortworxVolume: %s", version)
	}
}

func (s *PortworxVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			PortworxVolume: &v1.PortworxVolumeSource{
				VolumeID: s.VolumeID,
				FSType:   s.FSType,
				ReadOnly: s.ReadOnly,
			},
		},
	}, nil
}
