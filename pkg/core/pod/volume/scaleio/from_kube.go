package scaleio

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// NewScaleIOVolumeFromKubeScaleIOVolumeSource will create a new
// ScaleIOVolume object with the data from a provided kubernetes
// ScaleIOVolumeSource object
func NewScaleIOVolumeFromKubeScaleIOVolumeSource(obj interface{}) (*ScaleIOVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ScaleIOVolumeSource{}):
		o := obj.(v1.ScaleIOVolumeSource)
		return fromKubeScaleIOVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.ScaleIOVolumeSource{}):
		return fromKubeScaleIOVolumeSourceV1(obj.(*v1.ScaleIOVolumeSource))
	default:
		return nil, fmt.Errorf("unknown ScaleIOVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeScaleIOVolumeSourceV1(vol *v1.ScaleIOVolumeSource) (*ScaleIOVolume, error) {
	mode, err := fromScaleIOStorageMode(vol.StorageMode)
	if err != nil {
		return nil, err
	}
	return &ScaleIOVolume{
		Gateway:          vol.Gateway,
		System:           vol.System,
		SecretRef:        converterutils.FromKubeLocalObjectReferenceV1(vol.SecretRef),
		SSLEnabled:       vol.SSLEnabled,
		ProtectionDomain: vol.ProtectionDomain,
		StoragePool:      vol.StoragePool,
		StorageMode:      mode,
		VolumeName:       vol.VolumeName,
		FSType:           vol.FSType,
		ReadOnly:         vol.ReadOnly,
	}, nil
}

func fromScaleIOStorageMode(mode string) (ScaleIOStorageMode, error) {
	if len(mode) == 0 {
		return "", nil
	}

	switch mode {
	case "ThickProvisioned":
		return ScaleIOStorageModeThick, nil
	case "ThinProvisioned":
		return ScaleIOStorageModeThin, nil
	default:
		return "", serrors.InvalidValueErrorf(mode, "unrecognized ScaleIO storage mode")
	}
}
