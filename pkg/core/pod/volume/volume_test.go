package volume

import (
	"reflect"
	"testing"

	. "mantle/internal/pkg/core/pod/volume/pvc"

	"k8s.io/api/core/v1"
)

func TestVolumeToKube(t *testing.T) {
	v := Volume{
		PVC: &PVCVolume{},
	}

	testcases := []struct {
		name     string
		version  string
		expected interface{}
		pass     bool
	}{
		{
			name:     "v1 volume",
			version:  "v1",
			expected: &v1.Volume{},
			pass:     true,
		},
		{
			name:     "unknown volume version",
			version:  "invalid",
			expected: nil,
			pass:     false,
		},
	}

	for _, tc := range testcases {
		obj, err := v.ToKube(tc.version)
		if (err != nil) != !tc.pass {
			t.Errorf("%s: ToKube failed unexpectedly with %v", tc.name, err)
		}

		got := reflect.TypeOf(obj)
		exp := reflect.TypeOf(tc.expected)
		if got != exp {
			t.Errorf("%s: wrong object type.  Got %v expected %v", tc.name, got, exp)
		}
	}
}
