package ebs

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type AwsEBSVolume struct {
	VolumeID  string `json:"-"`
	FSType    string `json:"fs,omitempty"`
	Partition int32  `json:"partition,omitempty"`
	ReadOnly  bool   `json:"ro,omitempty"`
}

func (s *AwsEBSVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (ebs uuid) for %s", marshal.VolumeTypeAwsEBS)
	}
	s.VolumeID = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeAwsEBS)
	}

	return nil
}

func (s AwsEBSVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeAwsEBS)
	}

	if len(s.VolumeID) == 0 {
		return nil, serrors.InvalidInstanceErrorf(&s, "selector must contain ebs uuid")
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeAwsEBS,
		Selector:    []string{s.VolumeID},
		ExtraFields: obj,
	}, nil
}
