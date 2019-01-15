package marshal

const (
	VolumeTypeHostPath     = "host_path"
	VolumeTypeEmptyDir     = "empty_dir"
	VolumeTypeGcePD        = "gce_pd"
	VolumeTypeAwsEBS       = "aws_ebs"
	VolumeTypeAzureDisk    = "azure_disk"
	VolumeTypeAzureFile    = "azure_file"
	VolumeTypeCephFS       = "cephfs"
	VolumeTypeCinder       = "cinder"
	VolumeTypeFibreChannel = "fc"
	VolumeTypeFlex         = "flex"
	VolumeTypeFlocker      = "flocker"
	VolumeTypeGlusterfs    = "glusterfs"
	VolumeTypeISCSI        = "iscsi"
	VolumeTypeNFS          = "nfs"
	VolumeTypePhotonPD     = "photon"
	VolumeTypePortworx     = "portworx"
	VolumeTypePVC          = "pvc"
	VolumeTypeQuobyte      = "quobyte"
	VolumeTypeScaleIO      = "scaleio"
	VolumeTypeVsphere      = "vsphere"
	VolumeTypeConfigMap    = "config-map"
	VolumeTypeSecret       = "secret"
	VolumeTypeDownwardAPI  = "downward_api"
	VolumeTypeProjected    = "projected"
	VolumeTypeGit          = "git"
	VolumeTypeRBD          = "rbd"
	VolumeTypeStorageOS    = "storageos"
	VolumeTypeAny          = "*"

	SelectorSegmentReadOnly = "ro"
)

type MarshalledVolume struct {
	Type        string
	Selector    []string
	ExtraFields map[string]interface{}
}
