package scaleio

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ScaleIOVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for ScaleIOVolume: %s", version)
	}
}

func (s *ScaleIOVolume) toKubeV1() (*v1.Volume, error) {
	mode, err := s.toKubeV1ScaleIOStorageMode()
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, "ScaleIO storage mode")
	}
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			ScaleIO: &v1.ScaleIOVolumeSource{
				Gateway:          s.Gateway,
				System:           s.System,
				SecretRef:        converterutils.NewKubeLocalObjectRefV1(s.SecretRef),
				SSLEnabled:       s.SSLEnabled,
				ProtectionDomain: s.ProtectionDomain,
				StoragePool:      s.StoragePool,
				StorageMode:      mode,
				VolumeName:       s.VolumeName,
				FSType:           s.FSType,
				ReadOnly:         s.ReadOnly,
			},
		},
	}, nil
}

func (s *ScaleIOVolume) toKubeV1ScaleIOStorageMode() (string, error) {
	if len(s.StorageMode) == 0 {
		return "", nil
	}

	switch s.StorageMode {
	case ScaleIOStorageModeThick:
		return "ThickProvisioned", nil
	case ScaleIOStorageModeThin:
		return "ThinProvisioned", nil
	default:
		return "", serrors.InvalidInstanceError(s.StorageMode)
	}
}
