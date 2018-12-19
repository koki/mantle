package vsphere

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type VsphereVolume struct {
	VolumePath    string                `json:"-"`
	FSType        string                `json:"fs,omitempty"`
	StoragePolicy *VsphereStoragePolicy `json:"policy,omitempty"`
}

type VsphereStoragePolicy struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

func (s *VsphereVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (volume path) for %s", marshal.VolumeTypeVsphere)
	}
	s.VolumePath = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeVsphere)
	}

	return nil
}

func (s VsphereVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeVsphere)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeVsphere,
		Selector:    []string{s.VolumePath},
		ExtraFields: obj,
	}, nil
}
