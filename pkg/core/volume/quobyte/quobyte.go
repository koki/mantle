package quobyte

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type QuobyteVolume struct {
	Registry string `json:"registry"`
	Volume   string `json:"-"`
	ReadOnly bool   `json:"ro,omitempty"`
	User     string `json:"user,omitempty"`
	Group    string `json:"group,omitempty"`
}

func (s *QuobyteVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (volume ID) for %s", marshal.VolumeTypeQuobyte)
	}
	s.Volume = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeQuobyte)
	}

	return nil
}

func (s QuobyteVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeQuobyte)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeQuobyte,
		Selector:    []string{s.Volume},
		ExtraFields: obj,
	}, nil
}
