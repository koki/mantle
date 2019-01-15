package projected

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *DownwardAPIProjection) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for DownwardAPIProjection: %s", version)
	}
}

func (s *DownwardAPIProjection) toKubeV1() (*v1.DownwardAPIProjection, error) {
	items := []*v1.DownwardAPIVolumeFile{}

	for path, vol := range s.Items {
		item, err := vol.ToKube("v1")
		if err != nil {
			return nil, err
		}
		item.Path = path
		items = append(items, item)
	}

	return &v1.DownwardAPIProjection{
		Items: items,
	}, nil
}
