package resources

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewCPUFromKubeResourceRequirements will create a new
// CPU object with the data from a provided kubernetes
// ResourceRequirements object
func NewCPUFromKubeResourceRequirements(obj interface{}) (*CPU, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ResourceRequirements{}):
		return fromCPUResourceRequirementsV1(obj.(v1.ResourceRequirements))
	case reflect.TypeOf(&v1.ResourceRequirements{}):
		o := obj.(*v1.ResourceRequirements)
		return fromCPUResourceRequirementsV1(*o)
	default:
		return nil, fmt.Errorf("unknown ResourceRequirements version: %s", reflect.TypeOf(obj))
	}
}

func fromCPUResourceRequirementsV1(resources v1.ResourceRequirements) (*CPU, error) {
	cpu := &CPU{}

	if resources.Limits != nil {
		max := ""
		if q, ok := resources.Limits["cpu"]; ok {
			max = q.String()
		}
		cpu.Max = max
	}

	if resources.Requests != nil {
		min := ""
		if q, ok := resources.Requests["cpu"]; ok {
			min = q.String()
		}
		cpu.Min = min
	}

	if !cpu.IsEmpty() {
		return cpu, nil
	}

	return nil, nil
}
