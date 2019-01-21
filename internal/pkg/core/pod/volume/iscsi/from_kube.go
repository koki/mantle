package iscsi

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"
	"mantle/pkg/util"

	"k8s.io/api/core/v1"
)

// NewAwsEBSVolumeFromKubeISCSIVolumeSource will create a new
// AwsEBSVolume object with the data from a provided kubernetes
// ISCSIVolumeSource object
func NewAwsEBSVolumeFromKubeISCSIVolumeSource(obj interface{}) (*ISCSIVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ISCSIVolumeSource{}):
		o := obj.(v1.ISCSIVolumeSource)
		return fromKubeISCSIVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.AWSElasticBlockStoreVolumeSource{}):
		return fromKubeISCSIVolumeSourceV1(obj.(*v1.ISCSIVolumeSource))
	default:
		return nil, fmt.Errorf("unknown ISCSIVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeISCSIVolumeSourceV1(vol *v1.ISCSIVolumeSource) (*ISCSIVolume, error) {
	return &ISCSIVolume{
		TargetPortal:      vol.TargetPortal,
		IQN:               vol.IQN,
		Lun:               vol.Lun,
		ISCSIInterface:    vol.ISCSIInterface,
		FSType:            vol.FSType,
		ReadOnly:          vol.ReadOnly,
		Portals:           vol.Portals,
		DiscoveryCHAPAuth: vol.DiscoveryCHAPAuth,
		SessionCHAPAuth:   vol.SessionCHAPAuth,
		SecretRef:         converterutils.FromKubeLocalObjectReferenceV1(vol.SecretRef),
		InitiatorName:     util.FromStringPtr(vol.InitiatorName),
	}, nil
}
