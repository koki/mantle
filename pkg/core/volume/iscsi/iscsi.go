package iscsi

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type ISCSIVolume struct {
	TargetPortal   string   `json:"target_portal"`
	IQN            string   `json:"iqn"`
	Lun            int32    `json:"lun"`
	ISCSIInterface string   `json:"iscsi_interface,omitempty"`
	FSType         string   `json:"fs,omitempty"`
	ReadOnly       bool     `json:"ro,omitempty"`
	Portals        []string `json:"portals,omitempty"`
	// TODO: should this actually be "chap_auth"?
	DiscoveryCHAPAuth bool   `json:"chap_discovery,omitempty"`
	SessionCHAPAuth   bool   `json:"chap_session,omitempty"`
	SecretRef         string `json:"secret,omitempty"`
	// NOTE: InitiatorName is a pointer in k8s
	InitiatorName string `json:"initiator,omitempty"`
}

func (s *ISCSIVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected zero selector segments for %s", marshal.VolumeTypeISCSI)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeISCSI)
	}

	return nil
}

func (s ISCSIVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeISCSI)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeISCSI,
		ExtraFields: obj,
	}, nil
}
