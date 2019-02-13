package quobyte

type QuobyteVolume struct {
	Registry string `json:"registry"`
	Volume   string `json:"-"`
	ReadOnly bool   `json:"ro,omitempty"`
	User     string `json:"user,omitempty"`
	Group    string `json:"group,omitempty"`
}
