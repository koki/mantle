package gluster

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type GlusterfsVolume struct {
	EndpointsName string `json:"endpoints"`

	// Path is the Glusterfs volume name.
	Path     string `json:"path"`
	ReadOnly bool   `json:"ro,omitempty"`
}

func (s *GlusterfsVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected zero selector segments for %s", marshal.VolumeTypeGlusterfs)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeGlusterfs)
	}

	return nil
}

func (s GlusterfsVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeGlusterfs)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeGlusterfs,
		ExtraFields: obj,
	}, nil
}
