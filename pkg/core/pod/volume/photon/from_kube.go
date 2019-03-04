package photon

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewPhotonPDVolumeFromKubePhotonPersistentDiskVolumeSource will create a new
// PhotonPDVolume object with the data from a provided kubernetes
// PhotonPersistentDiskVolumeSource object
func NewPhotonPDVolumeFromKubePhotonPersistentDiskVolumeSource(obj interface{}) (*PhotonPDVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.PhotonPersistentDiskVolumeSource{}):
		o := obj.(v1.PhotonPersistentDiskVolumeSource)
		return fromKubePhotonPersistentDiskVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.PhotonPersistentDiskVolumeSource{}):
		return fromKubePhotonPersistentDiskVolumeSourceV1(obj.(*v1.PhotonPersistentDiskVolumeSource))
	default:
		return nil, fmt.Errorf("unknown PhotonPersistentDiskVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubePhotonPersistentDiskVolumeSourceV1(vol *v1.PhotonPersistentDiskVolumeSource) (*PhotonPDVolume, error) {
	return &PhotonPDVolume{
		PdID:   vol.PdID,
		FSType: vol.FSType,
	}, nil
}
