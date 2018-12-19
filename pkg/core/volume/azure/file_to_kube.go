package azure

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *AzureFileVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for AzureFileVolume: %s", version)
	}
}

func (s *AzureFileVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			AzureFile: &v1.AzureFileVolumeSource{
				SecretName: s.SecretName,
				ShareName:  s.ShareName,
				ReadOnly:   s.ReadOnly,
			},
		},
	}, nil
}
