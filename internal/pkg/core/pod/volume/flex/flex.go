package flex

type FlexVolume struct {
	Driver    string            `json:"-"`
	FSType    string            `json:"fs,omitempty"`
	SecretRef string            `json:"secret,omitempty"`
	ReadOnly  bool              `json:"ro,omitempty"`
	Options   map[string]string `json:"options,omitempty"`
}
