package azure

import (
	"fmt"
	"reflect"

	"mantle/pkg/util"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// NewAzureDiskVolumeFromAzureDiskVolumeSource will create a new
// AzureDiskVolume object with the data from a provided kubernetes
// AzureDiskVolumeSource object
func NewAzureDiskVolumeFromAzureDiskVolumeSource(obj interface{}) (*AzureDiskVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.AzureDiskVolumeSource{}):
		obj := obj.(v1.AzureDiskVolumeSource)
		return fromKubeAzureDiskVolumeSourceV1(&obj)
	case reflect.TypeOf(&v1.AzureDiskVolumeSource{}):
		return fromKubeAzureDiskVolumeSourceV1(obj.(*v1.AzureDiskVolumeSource))
	default:
		return nil, fmt.Errorf("unknown AzureDiskVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeAzureDiskVolumeSourceV1(vol *v1.AzureDiskVolumeSource) (*AzureDiskVolume, error) {
	var kind *AzureDataDiskKind
	var mode *AzureDataDiskCachingMode

	fstype := util.FromStringPtr(vol.FSType)
	readOnly := util.FromBoolPtr(vol.ReadOnly)

	if vol.Kind != nil {
		var azureKind AzureDataDiskKind

		switch *vol.Kind {
		case v1.AzureDedicatedBlobDisk:
			azureKind = AzureDedicatedBlobDisk
		case v1.AzureSharedBlobDisk:
			azureKind = AzureSharedBlobDisk
		case v1.AzureManagedDisk:
			azureKind = AzureManagedDisk
		default:
			return nil, serrors.InvalidValueErrorf(*vol.Kind, "unrecognized kind")
		}
		kind = &azureKind
	}

	if vol.CachingMode != nil {
		var cachingMode AzureDataDiskCachingMode

		switch *vol.CachingMode {
		case v1.AzureDataDiskCachingNone:
			cachingMode = AzureDataDiskCachingNone
		case v1.AzureDataDiskCachingReadOnly:
			cachingMode = AzureDataDiskCachingReadOnly
		case v1.AzureDataDiskCachingReadWrite:
			cachingMode = AzureDataDiskCachingReadWrite
		default:
			return nil, serrors.InvalidValueErrorf(*vol.CachingMode, "unrecognized cache")
		}
		mode = &cachingMode
	}

	return &AzureDiskVolume{
		DiskName:    vol.DiskName,
		DataDiskURI: vol.DataDiskURI,
		FSType:      fstype,
		ReadOnly:    readOnly,
		Kind:        kind,
		CachingMode: mode,
	}, nil
}
