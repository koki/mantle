package ebs

type AwsEBSVolume struct {
	VolumeID  string `json:"-"`
	FSType    string `json:"fs,omitempty"`
	Partition int32  `json:"partition,omitempty"`
	ReadOnly  bool   `json:"ro,omitempty"`
}
