package gluster

type GlusterfsVolume struct {
	EndpointsName string `json:"endpoints"`

	// Path is the Glusterfs volume name.
	Path     string `json:"path"`
	ReadOnly bool   `json:"ro,omitempty"`
}
