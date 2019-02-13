package portworx

type PortworxVolume struct {
	VolumeID string `json:"-"`
	FSType   string `json:"fs,omitempty"`
	ReadOnly bool   `json:"ro,omitempty"`
}
