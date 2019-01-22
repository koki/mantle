package gcepd

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewGcePDVolumeFromKubeGCEPersistentDiskVolumeSource will create a new
// GcePDVolume object with the data from a provided kubernetes
// GCEPersistentDiskVolumeSource object
func NewGcePDVolumeFromKubeGCEPersistentDiskVolumeSource(obj interface{}) (*GcePDVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.GCEPersistentDiskVolumeSource{}):
		o := obj.(v1.GCEPersistentDiskVolumeSource)
		return fromKubeGCEPersistentDiskVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.GCEPersistentDiskVolumeSource{}):
		return fromKubeGCEPersistentDiskVolumeSourceV1(obj.(*v1.GCEPersistentDiskVolumeSource))
	default:
		return nil, fmt.Errorf("unknown GCEPersistentDiskVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeGCEPersistentDiskVolumeSourceV1(vol *v1.GCEPersistentDiskVolumeSource) (*GcePDVolume, error) {
	return &GcePDVolume{
		PDName:    vol.PDName,
		FSType:    vol.FSType,
		Partition: vol.Partition,
		ReadOnly:  vol.ReadOnly,
	}, nil
}
