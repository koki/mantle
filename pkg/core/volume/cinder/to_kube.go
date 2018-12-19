package cinder

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *CinderVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for CinderVolume: %s", version)
	}
}

func (s *CinderVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			Cinder: &v1.CinderVolumeSource{
				VolumeID: s.VolumeID,
				FSType:   s.FSType,
				ReadOnly: s.ReadOnly,
			},
		},
	}, nil
}
