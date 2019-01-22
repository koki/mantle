package pvc

type PVCVolume struct {
	ClaimName string `json:"-"`
	ReadOnly  bool   `json:"-"`
}
