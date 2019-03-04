package nfs

type NFSVolume struct {
	Server   string `json:"-"`
	Path     string `json:"-"`
	ReadOnly bool   `json:"-"`
}
