package projected

import (
	"fmt"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *VolumeProjection) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1(version)
	case "":
		return s.toKubeV1(version)
	default:
		return nil, fmt.Errorf("unsupported api version for VolumeProjection: %s", version)
	}
}

func (s *VolumeProjection) toKubeV1(version string) (*v1.VolumeProjection, error) {
	var secret *v1.SecretProjection
	var dAPI *v1.DownwardAPIProjection
	var config *v1.ConfigMapProjection

	if s.Secret != nil {
		v1Ptr, err := s.Secret.ToKube(version)
		secret = v1Ptr.(*v1.SecretProjection)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "secret volume-projection %+v", s.Secret)
		}
	}

	if s.ConfigMap != nil {
		v1Ptr, err := s.ConfigMap.ToKube(version)
		config = v1Ptr.(*v1.ConfigMapProjection)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "config-map volume-projection %+v", s.ConfigMap)
		}
	}

	if s.DownwardAPI != nil {
		v1Ptr, err := s.DownwardAPI.ToKube(version)
		dAPI = v1Ptr.(*v1.DownwardAPIProjection)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "downwardAPI volume-projection %+v", s.DownwardAPI)
		}
	}

	return &v1.VolumeProjection{
		Secret:      secret,
		DownwardAPI: dAPI,
		ConfigMap:   config,
	}, nil
}
