package probe

import (
	"fmt"
	"reflect"

	"mantle/pkg/core/action"

	"k8s.io/api/core/v1"
)

// NewProbeFromKubeProbe will create a new
// Probe object with the data from a provided kubernetes
// Probe object
func NewProbeFromKubeProbe(obj interface{}) (*Probe, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.Probe{}):
		o := obj.(v1.Probe)
		return fromKubeProbeV1(&o)
	case reflect.TypeOf(&v1.Probe{}):
		return fromKubeProbeV1(obj.(*v1.Probe))
	default:
		return nil, fmt.Errorf("unknown Probe version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeProbeV1(probe *v1.Probe) (*Probe, error) {
	if probe == nil {
		return nil, nil
	}

	act, err := action.NewActionFromKubeHandler(&probe.Handler)
	if err != nil {
		return nil, err
	}

	var p *Probe
	if act != nil {
		p = &Probe{
			Action: *act,
		}
	} else {
		p = &Probe{}
	}

	p.Delay = probe.InitialDelaySeconds
	p.MinCountSuccess = probe.SuccessThreshold
	p.MinCountFailure = probe.FailureThreshold
	p.Interval = probe.PeriodSeconds
	p.Timeout = probe.TimeoutSeconds

	return p, nil
}
