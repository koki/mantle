package storageos

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *StorageOSVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for StorageOSVolume: %s", version)
	}
}

func (s *StorageOSVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			StorageOS: &v1.StorageOSVolumeSource{
				VolumeName:      s.VolumeName,
				VolumeNamespace: s.VolumeNamespace,
				FSType:          s.FSType,
				ReadOnly:        s.ReadOnly,
				SecretRef:       converterutils.NewKubeLocalObjectRefV1(s.SecretRef),
			},
		},
	}, nil
}
