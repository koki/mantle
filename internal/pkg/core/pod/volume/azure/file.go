package azure

import (
	"mantle/internal/marshal"

	serrors "github.com/koki/structurederrors"
)

type AzureFileVolume struct {
	SecretName string `json:"-"`
	ShareName  string `json:"-"`
	ReadOnly   bool   `json:"-"`
}

func (s *AzureFileVolume) Unmarshal(selector []string) error {
	if len(selector) > 3 || len(selector) < 2 {
		return serrors.InvalidValueErrorf(selector, "expected two or three selector segments for %s", marshal.VolumeTypeAzureFile)
	}

	s.SecretName = selector[0]
	s.ShareName = selector[1]

	if len(selector) > 2 {
		switch selector[2] {
		case marshal.SelectorSegmentReadOnly:
			s.ReadOnly = true
		default:
			return serrors.InvalidValueErrorf(selector[2], "invalid selector segment for %s", marshal.VolumeTypeAzureFile)
		}
	}

	return nil
}

func (s AzureFileVolume) Marshal() (*marshal.MarshalledVolume, error) {
	var selector []string
	if s.ReadOnly {
		selector = []string{s.SecretName, s.ShareName, marshal.SelectorSegmentReadOnly}
	} else {
		selector = []string{s.SecretName, s.ShareName}
	}
	return &marshal.MarshalledVolume{
		Type:     marshal.VolumeTypeAzureFile,
		Selector: selector,
	}, nil
}
