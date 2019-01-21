package git

import (
	"strings"

	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type GitVolume struct {
	Repository string `json:"-"`
	Revision   string `json:"rev,omitempty"`
	Directory  string `json:"dir,omitempty"`
}

func (s *GitVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	s.Repository = strings.Join(selector, ":")

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeGit)
	}

	return nil
}

func (s GitVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeGit)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeGit,
		Selector:    []string{s.Repository},
		ExtraFields: obj,
	}, nil
}
