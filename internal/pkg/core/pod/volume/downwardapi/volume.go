package downwardapi

import (
	"mantle/internal/marshal"
	"mantle/internal/pkg/core/pod/volume/filemode"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type DownwardAPIVolume struct {
	Items       map[string]DownwardAPIVolumeFile `json:"items,omitempty"`
	DefaultMode *filemode.FileMode               `json:"mode,omitempty"`
}

func (s *DownwardAPIVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected zero selector segments for %s", marshal.VolumeTypeDownwardAPI)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeDownwardAPI)
	}

	return nil
}

func (s DownwardAPIVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeDownwardAPI)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeDownwardAPI,
		ExtraFields: obj,
	}, nil
}
