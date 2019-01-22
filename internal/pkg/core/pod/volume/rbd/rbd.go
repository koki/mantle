package rbd

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
