package emptydir

import (
	"mantle/internal/marshal"

	"github.com/koki/json/jsonutil"
	serrors "github.com/koki/structurederrors"

	"k8s.io/apimachinery/pkg/api/resource"
)

type StorageMedium string

const (
	StorageMediumDefault   StorageMedium = ""           // use whatever the default is for the node
	StorageMediumMemory    StorageMedium = "memory"     // use memory (tmpfs)
	StorageMediumHugePages StorageMedium = "huge-pages" // use hugepages
)

type EmptyDirVolume struct {
	Medium    StorageMedium      `json:"medium,omitempty"`
	SizeLimit *resource.Quantity `json:"max_size,omitempty"`
}

func (s *EmptyDirVolume) Unmarshal(obj map[string]interface{}, selector []string) error {
	if len(selector) > 0 {
		return serrors.InvalidValueErrorf(selector, "expected zero selector segments for %s", marshal.VolumeTypeEmptyDir)
	}

	err := jsonutil.UnmarshalMap(obj, &s)
	if err != nil {
		return serrors.ContextualizeErrorf(err, marshal.VolumeTypeEmptyDir)
	}

	return nil
}

func (s EmptyDirVolume) Marshal() (*marshal.MarshalledVolume, error) {
	obj, err := jsonutil.MarshalMap(&s)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, marshal.VolumeTypeEmptyDir)
	}

	return &marshal.MarshalledVolume{
		Type:        marshal.VolumeTypeEmptyDir,
		ExtraFields: obj,
	}, nil
}
