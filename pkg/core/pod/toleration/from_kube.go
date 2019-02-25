package toleration

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewTolerationFromKubeToleration will create a new Toleration object with
// the data from a provided kubernetes Toleration object
func NewTolerationFromKubeToleration(obj interface{}) (*Toleration, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.Toleration{}):
		return fromKubeTolerationV1(obj.(v1.Toleration))
	case reflect.TypeOf(&v1.Toleration{}):
		o := obj.(*v1.Toleration)
		return fromKubeTolerationV1(*o)
	default:
		return nil, fmt.Errorf("unknown Toleration version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeTolerationV1(tol v1.Toleration) (*Toleration, error) {
	toleration := &Toleration{
		Key:               tol.Key,
		Value:             tol.Value,
		ExpirationSeconds: tol.TolerationSeconds,
	}

	switch tol.Operator {
	case v1.TolerationOpExists:
		toleration.Op = TolerationOperatorExists

	case v1.TolerationOpEqual:
		toleration.Op = TolerationOperatorEqual

	default:
		return nil, fmt.Errorf("unrecognized operator in toleration: %v", tol)
	}

	switch tol.Effect {
	case v1.TaintEffectNoSchedule:
		toleration.Effect = TaintEffectNoSchedule

	case v1.TaintEffectPreferNoSchedule:
		toleration.Effect = TaintEffectPreferNoSchedule

	case v1.TaintEffectNoExecute:
		toleration.Effect = TaintEffectNoExecute

	default:
		return nil, fmt.Errorf("unrecognized effect in toleration: %v", tol)
	}

	return toleration, nil
}
