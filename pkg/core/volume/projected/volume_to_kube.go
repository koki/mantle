package projected

import (
	"fmt"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *VolumeProjection) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1(version)
	case "":
		return s.toKubeV1(version)
	default:
		return nil, fmt.Errorf("unsupported api version for VolumeProjection: %s", version)
	}
}

func (s *VolumeProjection) toKubeV1(version string) (v1.VolumeProjection, error) {
	var secret *v1.SecretProjection
	var dAPI *v1.DownwardAPIProjection
	var config *v1.ConfigMapProjection
	var err error

	if s.Secret != nil {
		secret, err = s.Secret.ToKube(version)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "secret volume-projection %+v", s.Secret)
		}
	}

	if s.ConfigMap != nil {
		config, err = s.ConfigMap.ToKube(version)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "config-map volume-projection %+v", s.ConfigMap)
		}
	}

	if s.DownwardAPI != nil {
		dAPI, err = s.DownwardAPI.ToKube(version)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "downwardAPI volume-projection %+v", s.DownwardAPI)
		}
	}

	return v1.VolumeProjection{
		Secret:      secret,
		DownwardAPI: dAPI,
		ConfigMap:   config,
	}, nil
}
