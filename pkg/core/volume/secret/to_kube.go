package secret

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"
	"mantle/pkg/core/volume/filemode"
	"mantle/pkg/core/volume/keyandmode"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *SecretVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for SecretVolume: %s", version)
	}
}

func (s *SecretVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			Secret: &v1.SecretVolumeSource{
				SecretName:  s.SecretName,
				Items:       keyandmode.NewKubeV1KeyToPath(s.Items),
				DefaultMode: filemode.ConvertFileModeToInt32Ptr(s.DefaultMode),
				Optional:    converterutils.RequiredToOptional(s.Required),
			},
		},
	}, nil
}
