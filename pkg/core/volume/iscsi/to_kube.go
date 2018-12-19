package iscsi

import (
	"fmt"
	"strings"

	"mantle/internal/converterutils"
	"mantle/pkg/util"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *ISCSIVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for ISCSIVolume: %s", version)
	}
}

func (s *ISCSIVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			ISCSI: &v1.ISCSIVolumeSource{
				TargetPortal:      s.TargetPortal,
				IQN:               s.IQN,
				Lun:               s.Lun,
				ISCSIInterface:    s.ISCSIInterface,
				FSType:            s.FSType,
				ReadOnly:          s.ReadOnly,
				Portals:           s.Portals,
				DiscoveryCHAPAuth: s.DiscoveryCHAPAuth,
				SessionCHAPAuth:   s.SessionCHAPAuth,
				SecretRef:         converterutils.NewKubeV1LocalObjectRef(s.SecretRef),
				InitiatorName:     util.StringPtrOrNil(s.InitiatorName),
			},
		},
	}, nil
}
