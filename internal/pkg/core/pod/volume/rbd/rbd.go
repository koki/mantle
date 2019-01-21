package rbd

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type RBDVolume struct {
	CephMonitors []string `json:"monitors"`
	RBDImage     string   `json:"image"`
	FSType       string   `json:"fs,omitempty"`
	RBDPool      string   `json:"pool,omitempty"`
	RadosUser    string   `json:"user,omitempty"`
	Keyring      string   `json:"keyring,omitempty"`
	SecretRef    string   `json:"secret,omitempty"`
	ReadOnly     bool     `json:"ro,omitempty"`
}

func (s *RBDVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected zero selector segments for %s", marshal.VolumeTypeRBD)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeRBD)
	}

	return nil
}

func (s RBDVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeRBD)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeRBD,
		ExtraFields: obj,
	}, nil
}
