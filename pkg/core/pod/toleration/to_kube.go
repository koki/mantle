package toleration

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes toleration object of the api version provided
func (t *Toleration) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return t.toKubeV1()
	case "":
		return t.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Toleration: %s", version)
	}
}

func (t *Toleration) toKubeV1() (*v1.Toleration, error) {
	toleration := &v1.Toleration{
		Key:               t.Key,
		Value:             t.Value,
		TolerationSeconds: t.ExpirationSeconds,
	}

	switch t.Op {
	case TolerationOperatorExists:
		toleration.Operator = v1.TolerationOpExists
	case TolerationOperatorEqual:
		toleration.Operator = v1.TolerationOpEqual
	default:
		return nil, fmt.Errorf("unrecognized op in toleration: %v", t)
	}

	switch t.Effect {
	case TaintEffectNoSchedule:
		toleration.Effect = v1.TaintEffectNoSchedule
	case TaintEffectPreferNoSchedule:
		toleration.Effect = v1.TaintEffectPreferNoSchedule
	case TaintEffectNoExecute:
		toleration.Effect = v1.TaintEffectNoExecute
	default:
		return nil, fmt.Errorf("unrecognized effect in toleration: %v", t)
	}

	return toleration, nil
}
