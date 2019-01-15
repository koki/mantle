package secret

import (
	"mantle/internal/marshal"
	"mantle/pkg/core/volume/filemode"
	"mantle/pkg/core/volume/keyandmode"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type SecretVolume struct {
	SecretName string `json:"-"`

	Items       map[string]keyandmode.KeyAndMode `json:"items,omitempty"`
	DefaultMode *filemode.FileMode               `json:"mode,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}

func (s *SecretVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (secret name) for %s", marshal.VolumeTypeSecret)
	}
	s.SecretName = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeSecret)
	}

	return nil
}

func (s SecretVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeSecret)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeSecret,
		Selector:    []string{s.SecretName},
		ExtraFields: obj,
	}, nil
}
