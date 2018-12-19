package nfs

import (
	"mantle/internal/marshal"

	serrors "github.com/koki/structurederrors"
)

type NFSVolume struct {
	Server   string `json:"-"`
	Path     string `json:"-"`
	ReadOnly bool   `json:"-"`
}

func (s *NFSVolume) Unmarshal(selector []string) error {
	if len(selector) > 3 || len(selector) < 2 {
		return serrors.InvalidValueErrorf(selector, "expected two or three selector segments for %s", marshal.VolumeTypeNFS)
	}

	s.Server = selector[0]
	s.Path = selector[1]

	if len(selector) > 2 {
		switch selector[2] {
		case marshal.SelectorSegmentReadOnly:
			s.ReadOnly = true
		default:
			return serrors.InvalidValueErrorf(selector[2], "invalid selector segment for %s", marshal.VolumeTypeNFS)
		}
	}

	return nil
}

func (s NFSVolume) Marshal() (*marshal.MarshalledVolume, error) {
	var selector []string
	if s.ReadOnly {
		selector = []string{s.Server, s.Path, marshal.SelectorSegmentReadOnly}
	} else {
		selector = []string{s.Server, s.Path}
	}
	return &marshal.MarshalledVolume{
		Type:     marshal.VolumeTypeNFS,
		Selector: selector,
	}, nil
}
