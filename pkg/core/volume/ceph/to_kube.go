package ceph

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *CephFSVolume) ToKube(version string) (runtime.Object, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for CephFSVolume: %s", version)
	}
}

func (s *CephFSVolume) toKubeV1() (*v1.Volume, error) {
	secretFile, secretRef := s.toKubeV1LocalObjectReference()
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			CephFS: &v1.CephFSVolumeSource{
				Monitors:   s.Monitors,
				Path:       s.Path,
				User:       s.User,
				SecretFile: secretFile,
				SecretRef:  secretRef,
				ReadOnly:   s.ReadOnly,
			},
		},
	}, nil
}

func (s *CephFSVolume) toKubeV1LocalObjectReference() (string, *v1.LocalObjectReference) {
	if s.SecretFileOrRef == nil {
		return "", nil
	}

	if len(s.SecretFileOrRef.File) > 0 {
		return s.SecretFileOrRef.File, nil
	}

	return "", &v1.LocalObjectReference{
		Name: s.SecretFileOrRef.Ref,
	}
}
