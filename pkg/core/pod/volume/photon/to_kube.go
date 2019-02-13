package photon

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api versiono type defined in the object
func (s *PhotonPDVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for PhotonPDVolume: %s", version)
	}
}

func (s *PhotonPDVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			PhotonPersistentDisk: &v1.PhotonPersistentDiskVolumeSource{
				PdID:   s.PdID,
				FSType: s.FSType,
			},
		},
	}, nil
}
