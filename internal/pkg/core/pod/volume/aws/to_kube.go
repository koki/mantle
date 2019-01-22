package ebs

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *AwsEBSVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for AwsEBSVolume: %s", version)
	}
}

func (s *AwsEBSVolume) toKubeV1() (*v1.Volume, error) {
	store := &v1.AWSElasticBlockStoreVolumeSource{
		VolumeID:  s.VolumeID,
		FSType:    s.FSType,
		Partition: s.Partition,
		ReadOnly:  s.ReadOnly,
	}

	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			AWSElasticBlockStore: store,
		},
	}, nil
}
