package resources

import (
	"fmt"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/resource"
)

// ToKube will return a kubernetes ResourceRequirements object of the api version provided
func (m *Mem) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return m.toKubeV1()
	case "":
		return m.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Mem: %s", version)
	}
}

func (m *Mem) toKubeV1() (*v1.ResourceRequirements, error) {
	limits := v1.ResourceList{}
	requests := v1.ResourceList{}

	if m.Min != "" {
		q, err := resource.ParseQuantity(m.Min)
		if err != nil {
			return nil, serrors.InvalidInstanceErrorf(m, "couldn't parse mem min quantity: %s", err)
		}
		requests[v1.ResourceMemory] = q
	}

	if m.Max != "" {
		q, err := resource.ParseQuantity(m.Max)
		if err != nil {
			return nil, serrors.InvalidInstanceErrorf(m, "couldn't parse mem max quantity: %s", err)
		}
		limits[v1.ResourceMemory] = q
	}

	return &v1.ResourceRequirements{Limits: limits, Requests: requests}, nil
}
