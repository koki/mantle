package rbd

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *RBDVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for RBDVolume: %s", version)
	}
}

func (s *RBDVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			RBD: &v1.RBDVolumeSource{
				CephMonitors: s.CephMonitors,
				RBDImage:     s.RBDImage,
				FSType:       s.FSType,
				RBDPool:      s.RBDPool,
				RadosUser:    s.RadosUser,
				Keyring:      s.Keyring,
				SecretRef:    converterutils.NewKubeLocalObjectRefV1(s.SecretRef),
				ReadOnly:     s.ReadOnly,
			},
		},
	}, nil
}
