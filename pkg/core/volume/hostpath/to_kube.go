package hostpath

import (
	"fmt"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *HostPathVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for HostPathVolume: %s", version)
	}
}

func (s *HostPathVolume) toKubeV1() (*v1.Volume, error) {
	kubeType, err := s.toKubeV1HostPathType()
	if err != nil {
		return nil, err
	}

	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			HostPath: &v1.HostPathVolumeSource{
				Path: s.Path,
				Type: &kubeType,
			},
		},
	}, nil
}

func (s *HostPathVolume) toKubeV1HostPathType() (v1.HostPathType, error) {
	switch s.Type {
	case HostPathUnset:
		return v1.HostPathUnset, nil
	case HostPathDirectoryOrCreate:
		return v1.HostPathDirectoryOrCreate, nil
	case HostPathDirectory:
		return v1.HostPathDirectory, nil
	case HostPathFileOrCreate:
		return v1.HostPathFileOrCreate, nil
	case HostPathFile:
		return v1.HostPathFile, nil
	case HostPathSocket:
		return v1.HostPathSocket, nil
	case HostPathCharDev:
		return v1.HostPathCharDev, nil
	case HostPathBlockDev:
		return v1.HostPathBlockDev, nil
	default:
		return v1.HostPathUnset, serrors.InvalidValueErrorf(s.Type, "unrecognized host_path type")
	}
}
