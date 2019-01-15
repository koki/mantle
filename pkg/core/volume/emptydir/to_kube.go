package emptydir

import (
	"fmt"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *EmptyDirVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for EmptyDirVolume: %s", version)
	}
}

func (s *EmptyDirVolume) toKubeV1() (*v1.Volume, error) {
	medium, err := s.toKubeV1StorageMedium()
	if err != nil {
		return nil, err
	}
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			EmptyDir: &v1.EmptyDirVolumeSource{
				Medium:    medium,
				SizeLimit: s.SizeLimit,
			},
		},
	}, nil
}

func (s *EmptyDirVolume) toKubeV1StorageMedium() (v1.StorageMedium, error) {
	switch s.Medium {
	case StorageMediumDefault:
		return v1.StorageMediumDefault, nil
	case StorageMediumMemory:
		return v1.StorageMediumMemory, nil
	case StorageMediumHugePages:
		return v1.StorageMediumHugePages, nil
	default:
		return v1.StorageMediumDefault, serrors.InvalidValueErrorf(s.Medium, "unrecognized storage medium")
	}
}
