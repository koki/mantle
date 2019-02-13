package vsphere

import (
	"fmt"
	"reflect"


	"k8s.io/api/core/v1"
)

// NewVsphereVolumeFromKubeVsphereVirtualDiskVolumeSource will create a new
// VsphereVolume object with the data from a provided kubernetes
// VsphereVirtualDiskVolumeSource object
func NewVsphereVolumeFromKubeVsphereVirtualDiskVolumeSource(obj interface{}) (*VsphereVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.VsphereVirtualDiskVolumeSource{}):
		o := obj.(v1.VsphereVirtualDiskVolumeSource)
		return fromKubeVsphereVirtualDiskVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.VsphereVirtualDiskVolumeSource{}):
		return fromKubeVsphereVirtualDiskVolumeSourceV1(obj.(*v1.VsphereVirtualDiskVolumeSource))
	default:
		return nil, fmt.Errorf("unknown VsphereVirtualDiskVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeVsphereVirtualDiskVolumeSourceV1(vol *v1.VsphereVirtualDiskVolumeSource) (*VsphereVolume, error) {
	return &VsphereVolume{
		VolumePath:    vol.VolumePath,
		FSType:        vol.FSType,
		StoragePolicy: fromVsphereStoragePolicy(vol.StoragePolicyName, vol.StoragePolicyID),
	}, nil

}

func fromVsphereStoragePolicy(kubeName, kubeID string) *VsphereStoragePolicy {
	if len(kubeName) > 0 {
		return &VsphereStoragePolicy{
			Name: kubeName,
			ID:   kubeID,
		}
	}

	return nil
}
