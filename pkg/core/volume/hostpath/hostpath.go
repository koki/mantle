package hostpath

import (
	"mantle/internal/marshal"

	serrors "github.com/koki/structurederrors"
)

type HostPathType string

const (
	HostPathUnset             HostPathType = ""
	HostPathDirectoryOrCreate HostPathType = "dir-or-create"
	HostPathDirectory         HostPathType = "dir"
	HostPathFileOrCreate      HostPathType = "file-or-create"
	HostPathFile              HostPathType = "file"
	HostPathSocket            HostPathType = "socket"
	HostPathCharDev           HostPathType = "char-dev"
	HostPathBlockDev          HostPathType = "block-dev"
)

type HostPathVolume struct {
	Path string       `json:"-"`
	Type HostPathType `json:"-"`
}

func (s *HostPathVolume) Unmarshal(selector []string) error {
	if len(selector) > 2 || len(selector) == 0 {
		return serrors.InvalidValueErrorf(selector, "expected one or two selector segments for %s", marshal.VolumeTypeHostPath)
	}

	s.Path = selector[0]

	if len(selector) > 1 {
		hostPathType := HostPathType(selector[1])
		switch hostPathType {
		case HostPathUnset:
		case HostPathDirectoryOrCreate:
		case HostPathDirectory:
		case HostPathFileOrCreate:
		case HostPathFile:
		case HostPathSocket:
		case HostPathCharDev:
		case HostPathBlockDev:
		default:
			return serrors.InvalidValueErrorf(hostPathType, "invalid 'vol_type' selector for %s", marshal.VolumeTypeHostPath)
		}

		s.Type = hostPathType
	}

	return nil
}

func (s HostPathVolume) Marshal() (*marshal.MarshalledVolume, error) {
	var selector []string
	if len(s.Type) > 0 {
		selector = []string{s.Path, string(s.Type)}
	} else {
		selector = []string{s.Path}
	}
	return &marshal.MarshalledVolume{
		Type:     marshal.VolumeTypeHostPath,
		Selector: selector,
	}, nil
}
