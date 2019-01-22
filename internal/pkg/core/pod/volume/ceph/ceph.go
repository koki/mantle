package ceph

import (
	"encoding/json"
	"fmt"
	"regexp"

	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type CephFSVolume struct {
	Monitors        []string               `json:"monitors"`
	Path            string                 `json:"path, omitempty"`
	User            string                 `json:"user,omitempty"`
	SecretFileOrRef *CephFSSecretFileOrRef `json:"secret,omitempty"`
	ReadOnly        bool                   `json:"ro,omitempty"`
}

func (s *CephFSVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 0 {
		return serrors.InvalidValueErrorf(selector, "expected 0 selector segments for %s", marshal.VolumeTypeCephFS)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeCephFS)
	}

	return nil
}

func (s CephFSVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeCephFS)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeCephFS,
		ExtraFields: obj,
	}, nil
}

type CephFSSecretFileOrRef struct {
	File string `json:"-"`
	Ref  string `json:"-"`
}

var fileOrRefRegexp = regexp.MustCompile(`^(file|ref):(.*)$`)

func (s *CephFSSecretFileOrRef) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return serrors.ContextualizeErrorf(err, "cephfs secret should be a string")
	}

	matches := fileOrRefRegexp.FindStringSubmatch(str)
	if len(matches) > 0 {
		if matches[1] == "file" {
			s.File = matches[2]
		} else {
			s.Ref = matches[2]
		}
	} else {
		return serrors.InvalidValueErrorf(string(data), "unrecognized format for cephfs secret")
	}

	return nil
}

func (s CephFSSecretFileOrRef) MarshalJSON() ([]byte, error) {
	if len(s.Ref) > 0 {
		return json.Marshal(fmt.Sprintf("ref:%s", s.Ref))
	}

	return json.Marshal(fmt.Sprintf("file:%s", s.File))
}
