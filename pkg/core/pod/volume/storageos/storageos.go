package storageos

type StorageOSVolume struct {
	VolumeName      string `json:"-"`
	VolumeNamespace string `json:"vol_ns,omitempty"`
	FSType          string `json:"fs,omitempty"`
	ReadOnly        bool   `json:"ro,omitempty"`
	SecretRef       string `json:"secret,omitempty"`
}
