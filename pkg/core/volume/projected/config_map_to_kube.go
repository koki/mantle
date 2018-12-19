package projected

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"
	"mantle/pkg/core/volume/keyandmode"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ConfigMapProjection) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for ConfigMapProjection: %s", version)
	}
}

func (s *ConfigMapProjection) toKubeV1() (*v1.ConfigMapProjection, error) {
	ref := converterutils.NewKubeV1LocalObjectRef(s.Name)
	if ref == nil {
		return nil, serrors.InvalidInstanceErrorf(s, "config-map name is missing")
	}
	return &v1.ConfigMapProjection{
		LocalObjectReference: *ref,
		Items:                keyandmode.NewKubeV1KeyToPath(s.Items),
	}, nil
}
