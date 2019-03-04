package probe

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes Probe object of the api version provided
func (p *Probe) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return p.toKubeV1()
	case "":
		return p.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Probe: %s", version)
	}
}

func (p *Probe) toKubeV1() (*v1.Probe, error) {
	handler, err := p.Action.ToKube("v1")
	if err != nil {
		return nil, err
	}

	h := handler.(*v1.Handler)
	return &v1.Probe{
		Handler:             *h,
		InitialDelaySeconds: p.Delay,
		TimeoutSeconds:      p.Timeout,
		PeriodSeconds:       p.Interval,
		SuccessThreshold:    p.MinCountSuccess,
		FailureThreshold:    p.MinCountFailure,
	}, nil
}
