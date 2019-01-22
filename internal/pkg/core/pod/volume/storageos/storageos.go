package storageos

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type StorageOSVolume struct {
	VolumeName      string `json:"-"`
	VolumeNamespace string `json:"vol_ns,omitempty"`
	FSType          string `json:"fs,omitempty"`
	ReadOnly        bool   `json:"ro,omitempty"`
	SecretRef       string `json:"secret,omitempty"`
}

func (s *StorageOSVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (volume name) for %s", marshal.VolumeTypeStorageOS)
	}
	s.VolumeName = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeStorageOS)
	}

	return nil
}

func (s StorageOSVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeStorageOS)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeStorageOS,
		Selector:    []string{s.VolumeName},
		ExtraFields: obj,
	}, nil
}
