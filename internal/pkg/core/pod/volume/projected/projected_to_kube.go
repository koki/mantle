package projected

import (
	"fmt"
	"strings"

	"mantle/internal/pkg/core/pod/volume/filemode"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ProjectedVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for ProjectedVolume: %s", version)
	}
}

func (s *ProjectedVolume) toKubeV1() (*v1.Volume, error) {
	sources, err := s.toKubeVolumeProjectionsV1()
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, "volume (%+v)", s)
	}
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			Projected: &v1.ProjectedVolumeSource{
				Sources:     sources,
				DefaultMode: filemode.ConvertFileModeToInt32Ptr(s.DefaultMode),
			},
		},
	}, nil
}

func (s *ProjectedVolume) toKubeVolumeProjectionsV1() ([]v1.VolumeProjection, error) {
	if len(s.Sources) == 0 {
		return nil, nil
	}

	projections := []v1.VolumeProjection{}
	for _, projection := range s.Sources {
		vol, err := projection.ToKube("v1")
		if err != nil {
			return nil, err
		}

		projections = append(projections, vol.(v1.VolumeProjection))
	}

	return projections, nil
}
