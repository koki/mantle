package fiberchannel

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type FibreChannelVolume struct {
	TargetWWNs []string `json:"wwn,omitempty"`
	Lun        *int32   `json:"lun,omitempty"`
	FSType     string   `json:"fs,omitempty"`
	ReadOnly   bool     `json:"ro,omitempty"`
	WWIDs      []string `json:"wwid,omitempty"`
}

func (s *FibreChannelVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected 0 selector segments for %s", marshal.VolumeTypeFibreChannel)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeFibreChannel)
	}

	return nil
}

func (s FibreChannelVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeFibreChannel)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeFibreChannel,
		ExtraFields: obj,
	}, nil
}
