package projected

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"
	"mantle/internal/pkg/core/pod/volume/keyandmode"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *SecretProjection) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for SecretProjection: %s", version)
	}
}

func (s *SecretProjection) toKubeV1() (*v1.SecretProjection, error) {
	ref := converterutils.NewKubeLocalObjectRefV1(s.Name)
	if ref == nil {
		return nil, serrors.InvalidInstanceErrorf(s, "secret name is missing")
	}
	return &v1.SecretProjection{
		LocalObjectReference: *ref,
		Items:                keyandmode.NewKubeKeyToPathV1(s.Items),
	}, nil
}
