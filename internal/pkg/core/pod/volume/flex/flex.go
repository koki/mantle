package flex

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type FlexVolume struct {
	Driver    string            `json:"-"`
	FSType    string            `json:"fs,omitempty"`
	SecretRef string            `json:"secret,omitempty"`
	ReadOnly  bool              `json:"ro,omitempty"`
	Options   map[string]string `json:"options,omitempty"`
}

func (s *FlexVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (driver) for %s", marshal.VolumeTypeFlex)
	}
	s.Driver = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeFlex)
	}

	return nil
}

func (s FlexVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeFlex)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeFlex,
		Selector:    []string{s.Driver},
		ExtraFields: obj,
	}, nil
}
