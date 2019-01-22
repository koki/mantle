package azure

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

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

func (s *AzureDiskVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected zero selector segments for %s", marshal.VolumeTypeAzureDisk)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeAzureDisk)
	}

	return nil
}

func (s AzureDiskVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeAzureDisk)
	}

	if len(s.DiskName) == 0 {
		return nil, serrors.InvalidInstanceErrorf(&s, "disk_name is required for %s", marshal.VolumeTypeAzureDisk)
	}

	if len(s.DataDiskURI) == 0 {
		return nil, serrors.InvalidInstanceErrorf(&s, "disk_uri is required for %s", marshal.VolumeTypeAzureDisk)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeAzureDisk,
		ExtraFields: obj,
	}, nil
}
