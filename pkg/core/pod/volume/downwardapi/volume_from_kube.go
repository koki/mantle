package downwardapi

import (
	"fmt"
	"reflect"

	"mantle/pkg/core/pod/volume/util"

	"k8s.io/api/core/v1"
)

// NewDownwardAPIVolumeFromKubeDownwardAPIVolumeSource will create a new
// DownwardAPIVolume object with the data from a provided kubernetes
// DownwardAPIVolumeSource object
func NewDownwardAPIVolumeFromKubeDownwardAPIVolumeSource(obj interface{}) (*DownwardAPIVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.DownwardAPIVolumeSource{}):
		o := obj.(v1.DownwardAPIVolumeSource)
		return fromKubeDownwardAPIVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.DownwardAPIVolumeSource{}):
		return fromKubeDownwardAPIVolumeSourceV1(obj.(*v1.DownwardAPIVolumeSource))
	default:
		return nil, fmt.Errorf("unknown DownwardAPIVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeDownwardAPIVolumeSourceV1(vol *v1.DownwardAPIVolumeSource) (*DownwardAPIVolume, error) {
	items, err := FromKubeDownwardAPIVolumeFileV1(vol.Items)
	if err != nil {
		return nil, err
	}

	return &DownwardAPIVolume{
		Items:       items,
		DefaultMode: util.NewFileModeFromKubeV1(vol.DefaultMode),
	}, nil
}

func FromKubeDownwardAPIVolumeFileV1(vols []v1.DownwardAPIVolumeFile) (map[string]DownwardAPIVolumeFile, error) {
	if len(vols) == 0 {
		return nil, nil
	}

	items := map[string]DownwardAPIVolumeFile{}
	for _, vol := range vols {
		vf, err := NewDownwardAPIVolumeFileFromKubeDownwardAPIVolumeFile(vol)
		if err != nil {
			return nil, err
		}
		if vf != nil {
			items[vol.Path] = *vf
		}
	}

	return items, nil
}
