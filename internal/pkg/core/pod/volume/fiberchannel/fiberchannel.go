package fiberchannel

type FibreChannelVolume struct {
	TargetWWNs []string `json:"wwn,omitempty"`
	Lun        *int32   `json:"lun,omitempty"`
	FSType     string   `json:"fs,omitempty"`
	ReadOnly   bool     `json:"ro,omitempty"`
	WWIDs      []string `json:"wwid,omitempty"`
}
