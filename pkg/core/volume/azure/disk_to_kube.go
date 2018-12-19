package azure

import (
	"fmt"
	"strings"

	"mantle/pkg/util"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *AzureDiskVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for AzureDiskVolume: %s", version)
	}
}

func (s *AzureDiskVolume) toKubeV1() (*v1.Volume, error) {
	source, err := s.toKubeV1AzureDiskVolumeSource()
	if err != nil {
		return nil, err
	}
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			AzureDisk: source,
		},
	}, nil
}

func (s *AzureDiskVolume) toKubeV1AzureDiskVolumeSource() (*v1.AzureDiskVolumeSource, error) {
	kind, err := s.toKubeV1AzureDataDiskKind()
	if err != nil {
		return nil, err
	}
	cachingMode, err := s.toKubeV1AzureDataDiskCachingMode()
	if err != nil {
		return nil, err
	}
	return &v1.AzureDiskVolumeSource{
		DiskName:    s.DiskName,
		DataDiskURI: s.DataDiskURI,
		FSType:      util.StringPtrOrNil(s.FSType),
		ReadOnly:    util.BoolPtrOrNil(s.ReadOnly),
		Kind:        kind,
		CachingMode: cachingMode,
	}, nil
}

func (s *AzureDiskVolume) toKubeV1AzureDataDiskKind() (*v1.AzureDataDiskKind, error) {
	if s.Kind == nil {
		return nil, nil
	}

	var kind v1.AzureDataDiskKind
	switch *s.Kind {
	case AzureDedicatedBlobDisk:
		kind = v1.AzureDedicatedBlobDisk
	case AzureSharedBlobDisk:
		kind = v1.AzureSharedBlobDisk
	case AzureManagedDisk:
		kind = v1.AzureManagedDisk
	default:
		return nil, serrors.InvalidValueErrorf(s.Kind, "unrecognized kind")
	}

	return &kind, nil
}

func (s *AzureDiskVolume) toKubeV1AzureDataDiskCachingMode() (*v1.AzureDataDiskCachingMode, error) {
	if s.CachingMode == nil {
		return nil, nil
	}

	var mode v1.AzureDataDiskCachingMode
	switch *s.CachingMode {
	case AzureDataDiskCachingNone:
		mode = v1.AzureDataDiskCachingNone
	case AzureDataDiskCachingReadOnly:
		mode = v1.AzureDataDiskCachingReadOnly
	case AzureDataDiskCachingReadWrite:
		mode = v1.AzureDataDiskCachingReadWrite
	default:
		return nil, serrors.InvalidValueErrorf(s.CachingMode, "unrecognized cache")
	}

	return &mode, nil
}
