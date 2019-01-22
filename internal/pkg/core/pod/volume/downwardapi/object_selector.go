package downwardapi

import (
	"fmt"
	"strings"

	"github.com/koki/json"
	serrors "github.com/koki/structurederrors"
)

type ObjectFieldSelector struct {
	// required
	FieldPath string `json:"-"`

	// optional
	APIVersion string `json:"-"`
}

func (s *ObjectFieldSelector) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return serrors.ContextualizeErrorf(err, "field selector should be written as a string")
	}

	segments := strings.Split(str, ":")
	if len(segments) > 2 {
		return serrors.InvalidValueErrorf(str, "field selector should contain one or two segments")
	}

	s.FieldPath = segments[0]
	if len(segments) > 1 {
		s.APIVersion = segments[1]
	}

	return nil
}

func (s ObjectFieldSelector) MarshalJSON() ([]byte, error) {
	if len(s.APIVersion) == 0 {
		b, err := json.Marshal(s.FieldPath)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "field selector path")
		}

		return b, nil
	}

	b, err := json.Marshal(fmt.Sprintf("%s:%s", s.FieldPath, s.APIVersion))
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, "field selector")
	}

	return b, nil
}
