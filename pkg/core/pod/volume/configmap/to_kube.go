package configmap

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"
	"mantle/pkg/core/pod/volume/util"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ConfigMapVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for ConfigMapVolume: %s", version)
	}
}

func (s *ConfigMapVolume) toKubeV1() (*v1.Volume, error) {
	ref := converterutils.NewKubeLocalObjectRefV1(s.Name)
	if ref == nil {
		return nil, serrors.InvalidInstanceErrorf(s, "config name is required")
	}

	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			ConfigMap: &v1.ConfigMapVolumeSource{
				LocalObjectReference: *ref,
				Items:                util.NewKubeKeyToPathV1(s.Items),
				DefaultMode:          util.ConvertFileModeToInt32Ptr(s.DefaultMode),
				Optional:             converterutils.RequiredToOptional(s.Required),
			},
		},
	}, nil
}
