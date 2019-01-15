package configmap

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"
	"mantle/pkg/core/volume/filemode"
	"mantle/pkg/core/volume/keyandmode"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ConfigMapVolume) ToKube(version string) (runtime.Object, error) {
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
	ref := converterutils.NewKubeV1LocalObjectRef(s.Name)
	if ref == nil {
		return nil, serrors.InvalidInstanceErrorf(s, "config name is required")
	}

	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			ConfigMap: &v1.ConfigMapVolumeSource{
				LocalObjectReference: *ref,
				Items:                keyandmode.NewKubeV1KeyToPath(s.Items),
				DefaultMode:          filemode.ConvertFileModeToInt32Ptr(s.DefaultMode),
				Optional:             converterutils.RequiredToOptional(s.Required),
			},
		},
	}, nil
}
