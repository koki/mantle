package gcepd

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type GcePDVolume struct {
	PDName    string `json:"-"`
	FSType    string `json:"fs,omitempty"`
	Partition int32  `json:"partition,omitempty"`
	ReadOnly  bool   `json:"ro,omitempty"`
}

func (s *GcePDVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (disk name) for %s", marshal.VolumeTypeGcePD)
	}
	s.PDName = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeGcePD)
	}

	return nil
}

func (s GcePDVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeGcePD)
	}

	if len(s.PDName) == 0 {
		return nil, serrors.InvalidInstanceErrorf(&s, "selector must contain disk name")
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeGcePD,
		Selector:    []string{s.PDName},
		ExtraFields: obj,
	}, nil
}
