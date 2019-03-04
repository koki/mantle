package azure

type AzureDataDiskCachingMode string
type AzureDataDiskKind string

const (
	AzureDataDiskCachingNone      AzureDataDiskCachingMode = "none"
	AzureDataDiskCachingReadOnly  AzureDataDiskCachingMode = "ro"
	AzureDataDiskCachingReadWrite AzureDataDiskCachingMode = "rw"

	AzureSharedBlobDisk    AzureDataDiskKind = "shared"
	AzureDedicatedBlobDisk AzureDataDiskKind = "dedicated"
	AzureManagedDisk       AzureDataDiskKind = "managed"
)

type AzureDiskVolume struct {
	DiskName    string                    `json:"disk_name"`
	DataDiskURI string                    `json:"disk_uri"`
	CachingMode *AzureDataDiskCachingMode `json:"cache,omitempty"`
	FSType      string                    `json:"fs,omitempty"`
	ReadOnly    bool                      `json:"ro,omitempty"`
	Kind        *AzureDataDiskKind        `json:"kind,omitempty"`
}

type AzureFileVolume struct {
	SecretName string `json:"-"`
	ShareName  string `json:"-"`
	ReadOnly   bool   `json:"-"`
}
