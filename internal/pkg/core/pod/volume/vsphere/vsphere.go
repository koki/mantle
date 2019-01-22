package vsphere

type VsphereVolume struct {
	VolumePath    string                `json:"-"`
	FSType        string                `json:"fs,omitempty"`
	StoragePolicy *VsphereStoragePolicy `json:"policy,omitempty"`
}

type VsphereStoragePolicy struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
