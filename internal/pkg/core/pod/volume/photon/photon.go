package photon

import (
	"mantle/internal/marshal"

	serrors "github.com/koki/structurederrors"
)

type PhotonPDVolume struct {
	PdID   string `json:"-"`
	FSType string `json:"-"`
}

func (s *PhotonPDVolume) Unmarshal(selector []string) error {
	if len(selector) > 2 || len(selector) < 1 {
		return serrors.InvalidValueErrorf(selector, "expected one or two selector segments for %s", marshal.VolumeTypePhotonPD)
	}

	s.PdID = selector[0]

	if len(selector) > 1 {
		s.FSType = selector[1]
	}

	return nil
}

func (s PhotonPDVolume) Marshal() (*marshal.MarshalledVolume, error) {
	var selector []string
	if len(s.FSType) > 0 {
		selector = []string{s.PdID, s.FSType}
	} else {
		selector = []string{s.PdID}
	}
	return &marshal.MarshalledVolume{
		Type:     marshal.VolumeTypePhotonPD,
		Selector: selector,
	}, nil
}
