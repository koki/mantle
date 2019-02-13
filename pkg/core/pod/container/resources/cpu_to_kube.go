package resources

import (
	"fmt"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/resource"
)

// ToKube will return a kubernetes ResourceRequirements object of the api version provided
func (c *CPU) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return c.toKubeV1()
	case "":
		return c.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for CPU: %s", version)
	}
}

func (c *CPU) toKubeV1() (*v1.ResourceRequirements, error) {
	limits := v1.ResourceList{}
	requests := v1.ResourceList{}

	if c.Min != "" {
		q, err := resource.ParseQuantity(c.Min)
		if err != nil {
			return nil, serrors.InvalidInstanceErrorf(c, "couldn't parse cpu min quantity: %s", err)
		}
		requests[v1.ResourceCPU] = q
	}

	if c.Max != "" {
		q, err := resource.ParseQuantity(c.Max)
		if err != nil {
			return nil, serrors.InvalidInstanceErrorf(c, "couldn't parse cpu max quantity: %s", err)
		}
		limits[v1.ResourceCPU] = q
	}

	return &v1.ResourceRequirements{Limits: limits, Requests: requests}, nil
}
