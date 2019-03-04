package ceph

type CephFSVolume struct {
	Monitors        []string               `json:"monitors"`
	Path            string                 `json:"path, omitempty"`
	User            string                 `json:"user,omitempty"`
	SecretFileOrRef *CephFSSecretFileOrRef `json:"secret,omitempty"`
	ReadOnly        bool                   `json:"ro,omitempty"`
}

type CephFSSecretFileOrRef struct {
	File string `json:"-"`
	Ref  string `json:"-"`
}
