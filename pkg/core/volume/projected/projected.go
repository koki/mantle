package projected

import (
	"mantle/internal/marshal"
	"mantle/pkg/core/volume/filemode"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type ProjectedVolume struct {
	Sources     []VolumeProjection `json:"sources"`
	DefaultMode *filemode.FileMode `json:"mode,omitempty"`
}

func (s *ProjectedVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected zero selector segments for %s", marshal.VolumeTypeProjected)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeProjected)
	}

	return nil
}

func (s ProjectedVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeProjected)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeProjected,
		ExtraFields: obj,
	}, nil
}
