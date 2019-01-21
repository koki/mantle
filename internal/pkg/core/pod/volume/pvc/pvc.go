package pvc

import (
	"mantle/internal/marshal"

	serrors "github.com/koki/structurederrors"
)

type PVCVolume struct {
	ClaimName string `json:"-"`
	ReadOnly  bool   `json:"-"`
}

func (s *PVCVolume) Unmarshal(selector []string) error {
	if len(selector) > 2 || len(selector) < 1 {
		return serrors.InvalidValueErrorf(selector, "expected one or two selector segments for %s", marshal.VolumeTypePVC)
	}

	s.ClaimName = selector[0]

	if len(selector) > 1 {
		switch selector[1] {
		case marshal.SelectorSegmentReadOnly:
			s.ReadOnly = true
		default:
			return serrors.InvalidValueErrorf(selector[2], "invalid selector segment for %s", marshal.VolumeTypePVC)
		}
	}

	return nil
}

func (s PVCVolume) Marshal() (*marshal.MarshalledVolume, error) {
	var selector []string
	if s.ReadOnly {
		selector = []string{s.ClaimName, marshal.SelectorSegmentReadOnly}
	} else {
		selector = []string{s.ClaimName}
	}
	return &marshal.MarshalledVolume{
		Type:     marshal.VolumeTypePVC,
		Selector: selector,
	}, nil
}
