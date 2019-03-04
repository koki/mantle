package flex

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *FlexVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for FlexVolume: %s", version)
	}
}

func (s *FlexVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			FlexVolume: &v1.FlexVolumeSource{
				Driver:    s.Driver,
				FSType:    s.FSType,
				SecretRef: converterutils.NewKubeLocalObjectRefV1(s.SecretRef),
				ReadOnly:  s.ReadOnly,
				Options:   s.Options,
			},
		},
	}, nil
}
