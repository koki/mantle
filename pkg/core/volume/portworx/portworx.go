package portworx

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type PortworxVolume struct {
	VolumeID string `json:"-"`
	FSType   string `json:"fs,omitempty"`
	ReadOnly bool   `json:"ro,omitempty"`
}

func (s *PortworxVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (volume ID) for %s", marshal.VolumeTypePortworx)
	}
	s.VolumeID = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypePortworx)
	}

	return nil
}

func (s PortworxVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypePortworx)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypePortworx,
		Selector:    []string{s.VolumeID},
		ExtraFields: obj,
	}, nil
}
