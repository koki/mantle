package downwardapi

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/koki/json"
	serrors "github.com/koki/structurederrors"

	"k8s.io/apimachinery/pkg/api/resource"
)

type VolumeResourceFieldSelector struct {
	// required
	ContainerName string `json:"-"`

	// required
	Resource string `json:"-"`

	// optional
	Divisor resource.Quantity `json:"-"`
}

func (s *VolumeResourceFieldSelector) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return serrors.ContextualizeErrorf(err, "resource selector should be written as a string")
	}

	segments := strings.Split(str, ":")
	if len(segments) > 3 || len(segments) < 2 {
		return serrors.InvalidValueErrorf(str, "resource selector should contain two or three segments")
	}

	s.ContainerName = segments[0]
	s.Resource = segments[1]
	if len(segments) > 2 {
		divisor, err := resource.ParseQuantity(segments[2])
		if err != nil {
			return serrors.ContextualizeErrorf(err, "resource selector divisor")
		}
		s.Divisor = divisor
	}

	return nil
}

func (s VolumeResourceFieldSelector) MarshalJSON() ([]byte, error) {
	if reflect.DeepEqual(s.Divisor, resource.Quantity{}) {
		b, err := json.Marshal(fmt.Sprintf("%s:%s", s.ContainerName, s.Resource))
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "resource selector")
		}

		return b, nil
	}

	b, err := json.Marshal(fmt.Sprintf("%s:%s:%s", s.ContainerName, s.Resource, s.Divisor.String()))
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, "resource selector")
	}

	return b, nil
}
