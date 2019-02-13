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
func (s *DownwardAPIVolumeFile) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1(version)
	case "":
		return s.toKubeV1(version)
	default:
		return nil, fmt.Errorf("unsupported api version for DownwardAPIVolumeFile: %s", version)
	}
}

func (s *DownwardAPIVolumeFile) toKubeV1(version string) (*v1.DownwardAPIVolumeFile, error) {
	var fieldRef *v1.ObjectFieldSelector
	var resourceRef *v1.ResourceFieldSelector

	if s.FieldRef == nil {
		fieldRef = nil
	} else {
		ref, err := s.FieldRef.ToKube(version)
		if err != nil {
			return nil, err
		}
		fieldRef = reflect.ValueOf(ref).Interface().(*v1.ObjectFieldSelector)
	}

	if s.ResourceFieldRef == nil {
		resourceRef = nil
	} else {
		ref, err := s.ResourceFieldRef.ToKube(version)
		if err != nil {
			return nil, err
		}
		resourceRef = reflect.ValueOf(ref).Interface().(*v1.ResourceFieldSelector)
	}

	return &v1.DownwardAPIVolumeFile{
		FieldRef:         fieldRef,
		ResourceFieldRef: resourceRef,
		Mode:             util.ConvertFileModeToInt32Ptr(s.Mode),
	}, nil
}
