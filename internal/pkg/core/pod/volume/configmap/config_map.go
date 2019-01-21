package configmap

import (
	"mantle/internal/marshal"
	"mantle/internal/pkg/core/pod/volume/filemode"
	"mantle/internal/pkg/core/pod/volume/keyandmode"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"
)

type ConfigMapVolume struct {
	Name string `json:"-"`

	Items       map[string]keyandmode.KeyAndMode `json:"items,omitempty"`
	DefaultMode *filemode.FileMode               `json:"mode,omitempty"`

	// NOTE: opposite of Optional
	Required *bool `json:"required,omitempty"`
}

func (s *ConfigMapVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) != 1 {
		return serrors.InvalidValueErrorf(selector, "expected 1 selector segment (config name) for %s", marshal.VolumeTypeConfigMap)
	}
	s.Name = selector[0]

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeConfigMap)
	}

	return nil
}

func (s ConfigMapVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeConfigMap)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeConfigMap,
		Selector:    []string{s.Name},
		ExtraFields: obj,
	}, nil
}
