package downwardapi

import (
	"fmt"
	"reflect"
	"strings"

	"mantle/pkg/core/pod/volume/util"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *DownwardAPIVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for DownwardAPIVolume: %s", version)
	}
}

func (s *DownwardAPIVolume) toKubeV1() (*v1.Volume, error) {
	items, err := s.toKubeV1DownwardAPIVolumeFile()
	if err != nil {
		return nil, err
	}
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			DownwardAPI: &v1.DownwardAPIVolumeSource{
				Items:       items,
				DefaultMode: util.ConvertFileModeToInt32Ptr(s.DefaultMode),
			},
		},
	}, nil
}

func (s *DownwardAPIVolume) toKubeV1DownwardAPIVolumeFile() ([]v1.DownwardAPIVolumeFile, error) {
	if len(s.Items) == 0 {
		return nil, nil
	}

	items := []v1.DownwardAPIVolumeFile{}
	for path, vol := range s.Items {
		item, err := vol.ToKube("v1")
		if err != nil {
			return nil, err
		}
		vf := s.getReflectObjectFromInterface(item).Interface().(v1.DownwardAPIVolumeFile)
		vf.Path = path
		items = append(items, vf)
	}

	return items, nil
}

func (s *DownwardAPIVolume) getReflectObjectFromInterface(iface interface{}) reflect.Value {
	v := reflect.ValueOf(iface)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}
