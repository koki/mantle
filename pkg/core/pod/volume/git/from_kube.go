package git

import (
	"fmt"
	"reflect"


	"k8s.io/api/core/v1"
)

// NewGitVolumeFromKubeGitRepoVolumeSource will create a new
// GitVolume object with the data from a provided kubernetes
// GitRepoVolumeSource object
func NewGitVolumeFromKubeGitRepoVolumeSource(obj interface{}) (*GitVolume, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.GitRepoVolumeSource{}):
		o := obj.(v1.GitRepoVolumeSource)
		return fromKubeGitRepoVolumeSourceV1(&o)
	case reflect.TypeOf(&v1.GitRepoVolumeSource{}):
		return fromKubeGitRepoVolumeSourceV1(obj.(*v1.GitRepoVolumeSource))
	default:
		return nil, fmt.Errorf("unknown GitRepoVolumeSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeGitRepoVolumeSourceV1(vol *v1.GitRepoVolumeSource) (*GitVolume, error) {
	return &GitVolume{
		Repository: vol.Repository,
		Revision:   vol.Revision,
		Directory:  vol.Directory,
	}, nil
}
