package git

type GitVolume struct {
	Repository string `json:"-"`
	Revision   string `json:"rev,omitempty"`
	Directory  string `json:"dir,omitempty"`
}
