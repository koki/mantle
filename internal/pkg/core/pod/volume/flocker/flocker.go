package flocker

import (
	"mantle/internal/marshal"

	serrors "github.com/koki/structurederrors"
)

type FlockerVolume struct {
	DatasetUUID string `json:"-"`
}

func (s *FlockerVolume) Unmarshal(selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected exactly one selector segment (dataset UUID) for %s", marshal.VolumeTypeFlocker)
	}

	s.DatasetUUID = selector[0]

	return nil
}

func (s FlockerVolume) Marshal() (*marshal.MarshalledVolume, error) {
	return &marshal.MarshalledVolume{
		Type:     marshal.VolumeTypeFlocker,
		Selector: []string{s.DatasetUUID},
	}, nil
}
