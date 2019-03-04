package vsphere

import (
	"fmt"
	"strings"


	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *VsphereVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for VsphereVolume: %s", version)
	}
}

func (s *VsphereVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			VsphereVolume: s.toKubeV1VsphereVirtualDiskVolumeSource(),
		},
	}, nil
}

func (s *VsphereVolume) toKubeV1VsphereVirtualDiskVolumeSource() *v1.VsphereVirtualDiskVolumeSource {
	storagePolicyName, storagePolicyID := s.toKubeV1VSphereStoragePolicy()
	return &v1.VsphereVirtualDiskVolumeSource{
		VolumePath:        s.VolumePath,
		FSType:            s.FSType,
		StoragePolicyName: storagePolicyName,
		StoragePolicyID:   storagePolicyID,
	}
}

func (s *VsphereVolume) toKubeV1VSphereStoragePolicy() (name, ID string) {
	if s.StoragePolicy == nil {
		return "", ""
	}

	return s.StoragePolicy.Name, s.StoragePolicy.ID
}
