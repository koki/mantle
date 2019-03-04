package flocker

import (
	"fmt"
	"strings"


	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *FlockerVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for FlockerVolume: %s", version)
	}
}

func (s *FlockerVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			Flocker: &v1.FlockerVolumeSource{
				DatasetUUID: s.DatasetUUID,
			},
		},
	}, nil
}
