package git

import (
	"fmt"
	"strings"


	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (s *GitVolume) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return s.toKubeV1()
	case "":
		return s.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for GitVolume: %s", version)
	}
}

func (s *GitVolume) toKubeV1() (*v1.Volume, error) {
	return &v1.Volume{
		VolumeSource: v1.VolumeSource{
			GitRepo: &v1.GitRepoVolumeSource{
				Repository: s.Repository,
				Revision:   s.Revision,
				Directory:  s.Directory,
			},
		},
	}, nil
}
