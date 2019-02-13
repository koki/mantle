package volumemount

type VolumeMount struct {
	MountPath   string            `json:"mount,omitempty"`
	Propagation *MountPropagation `json:"propagation,omitempty"`
	Store       string            `json:"store,omitempty"`
	ReadOnly    bool              `json:"readOnly,omitemty"`
}

type MountPropagation string

const (
	MountPropagationHostToContainer MountPropagation = "host-to-container"
	MountPropagationBidirectional   MountPropagation = "bidirectional"
	MountPropagationNone            MountPropagation = "none"
)
