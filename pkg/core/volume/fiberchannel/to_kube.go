package fiberchannel

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *FibreChannelVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for FibreChannelVolume: %s", version)
	}
}

func (s *FibreChannelVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			FC: &v1.FCVolumeSource{
				TargetWWNs: s.TargetWWNs,
				Lun:        s.Lun,
				FSType:     s.FSType,
				ReadOnly:   s.ReadOnly,
				WWIDs:      s.WWIDs,
			},
		},
	}, nil
}
