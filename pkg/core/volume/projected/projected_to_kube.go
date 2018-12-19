package projected

import (
	"fmt"
	"strings"

	"mantle/pkg/core/volume/filemode"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ProjectedVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for ProjectedVolume: %s", version)
	}
}

func (s *ProjectedVolume) toKubeV1() (*v1.ProjectedVolumeSource, error) {
	sources, err := s.toKubeV1VolumeProjections()
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, "volume (%+v)", s)
	}
	return &v1.ProjectedVolumeSource{
		Sources:     sources,
		DefaultMode: filemode.ConvertFileModeToInt32Ptr(s.DefaultMode),
	}, nil
}

func (s *ProjectedVolume) toKubeV1VolumeProjections() ([]v1.VolumeProjection, error) {
	if len(s.Sources) == 0 {
		return nil, nil
	}

	projections := make([]v1.VolumeProjection, len(s.Sources))
	for _, projection := range s.Sources {
		vol, err := projection.ToKube("v1")
		if err != nil {
			return nil, err
		}

		projections[i] = vol
	}

	return projections, nil
}
