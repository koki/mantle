package resources

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewMemUFromKubeResourceRequirements will create a new
// MemU object with the data from a provided kubernetes
// ResourceRequirements object
func NewMemFromKubeResourceRequirements(obj interface{}) (*Mem, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ResourceRequirements{}):
		return fromMemResourceRequirementsV1(obj.(v1.ResourceRequirements))
	case reflect.TypeOf(&v1.ResourceRequirements{}):
		o := obj.(*v1.ResourceRequirements)
		return fromMemResourceRequirementsV1(*o)
	default:
		return nil, fmt.Errorf("unknown ResourceRequirements version: %s", reflect.TypeOf(obj))
	}
}

func fromMemResourceRequirementsV1(resources v1.ResourceRequirements) (*Mem, error) {
	mem := &Mem{}

	if resources.Limits != nil {
		max := ""
		if q, ok := resources.Limits["memory"]; ok {
			max = q.String()
		}
		mem.Max = max
	}

	if resources.Requests != nil {
		min := ""
		if q, ok := resources.Requests["memory"]; ok {
			min = q.String()
		}
		mem.Min = min
	}

	if !mem.IsEmpty() {
		return mem, nil
	}

	return nil, nil
}
