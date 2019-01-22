package iscsi

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
