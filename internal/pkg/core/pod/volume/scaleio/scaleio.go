package scaleio

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type ScaleIOStorageMode string

const (
	ScaleIOStorageModeThick ScaleIOStorageMode = "thick"
	ScaleIOStorageModeThin  ScaleIOStorageMode = "thin"
)

type ScaleIOVolume struct {
	Gateway          string             `json:"gateway"`
	System           string             `json:"system"`
	SecretRef        string             `json:"secret"`
	SSLEnabled       bool               `json:"ssl,omitempty"`
	ProtectionDomain string             `json:"protection_domain,omitempty"`
	StoragePool      string             `json:"storage_pool,omitempty"`
	StorageMode      ScaleIOStorageMode `json:"storage_mode,omitempty"`
	VolumeName       string             `json:"-"`
	FSType           string             `json:"fs,omitempty"`
	ReadOnly         bool               `json:"ro,omitempty"`
}

func (s *ScaleIOVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (volume name) for %s", marshal.VolumeTypeScaleIO)
	}
	s.VolumeName = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeScaleIO)
	}

	return nil
}

func (s ScaleIOVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeScaleIO)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeScaleIO,
		Selector:    []string{s.VolumeName},
		ExtraFields: obj,
	}, nil
}
