package projected

import (
	"fmt"
	"reflect"

	"mantle/pkg/core/volume/filemode"

	"k8s.io/api/core/v1"
)

// NewProjectedVolumeFromKubeProjectedVolumeSource will create a new
// ProjectedVolume object with the data from a provided kubernetes
// ProjectedVolumeSource object
func NewProjectedVolumeFromKubeProjectedVolumeSource(obj interface{}) (*ProjectedVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ProjectedVolumeSource{}):
		o := obj.(v1.ProjectedVolumeSource)
		return fromKubeProjectedVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.ProjectedVolumeSource{}):
		return fromKubeProjectedVolumeSourceV1(obj.(*v1.ProjectedVolumeSource))
	default:
		return nil, fmt.Errorf("unknown ProjectedVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeProjectedVolumeSourceV1(vol *v1.ProjectedVolumeSource) (*ProjectedVolume, error) {
	if len(vol.Sources) == 0 {
		return nil, nil
	}

	sources := []VolumeProjection{}
	for _, kubeSource := range vol.Sources {
		source, err := NewVolumeProjectionFromKubeVolumeProjection(kubeSource)
		if err != nil {
			return nil, err
		}
		if source != nil {
			sources = append(sources, *source)
		}
	}

	return &ProjectedVolume{
		Sources:     sources,
		DefaultMode: filemode.NewFileModeFromKubeV1(vol.DefaultMode),
	}, nil
}
