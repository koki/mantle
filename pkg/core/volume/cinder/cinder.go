package cinder

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type CinderVolume struct {
	VolumeID string `json:"-"`
	FSType   string `json:"fs,omitempty"`
	ReadOnly bool   `json:"ro,omitempty"`
}

func (s *CinderVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 (volume ID) selector segment for %s", marshal.VolumeTypeCinder)
	}

	s.VolumeID = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeCinder)
	}

	return nil
}

func (s CinderVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeCinder)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeCinder,
		Selector:    []string{s.VolumeID},
		ExtraFields: obj,
	}, nil
}
