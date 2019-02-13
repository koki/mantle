package scaleio

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
