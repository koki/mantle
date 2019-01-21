package projected

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *DownwardAPIProjection) ToKube(version string) (interface{}, error) {
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
	items := []v1.DownwardAPIVolumeFile{}

	for path, vol := range s.Items {
		item, err := vol.ToKube("v1")
		if err != nil {
			return nil, err
		}
		v1Item := item.(*v1.DownwardAPIVolumeFile)
		v1Item.Path = path
		items = append(items, *v1Item)
	}

	return &v1.DownwardAPIProjection{
		Items: items,
	}, nil
}
