package ebs

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewAwsEBSVolumeFromKubeAWSElasticBlockStoreVolumeSource will create a new
// AwsEBSVolume object with the data from a provided kubernetes
// AWSElasticBlockStoreVolumeSource object
func NewAwsEBSVolumeFromKubeAWSElasticBlockStoreVolumeSource(obj interface{}) (*AwsEBSVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.AWSElasticBlockStoreVolumeSource{}):
		o := obj.(v1.AWSElasticBlockStoreVolumeSource)
		return fromKubeAWSElasticBlockStoreVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.AWSElasticBlockStoreVolumeSource{}):
		return fromKubeAWSElasticBlockStoreVolumeSourceV1(obj.(*v1.AWSElasticBlockStoreVolumeSource))
	default:
		return nil, fmt.Errorf("unknown AWSElasticBlockStoreVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeAWSElasticBlockStoreVolumeSourceV1(vol *v1.AWSElasticBlockStoreVolumeSource) (*AwsEBSVolume, error) {
	return &AwsEBSVolume{
		VolumeID:  vol.VolumeID,
		FSType:    vol.FSType,
		Partition: vol.Partition,
		ReadOnly:  vol.ReadOnly,
	}, nil
}
