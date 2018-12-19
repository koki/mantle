package keyandmode

import (
	"fmt"
	"regexp"
	"strconv"

	"mantle/pkg/core/volume/filemode"

	"github.com/koki/json"
	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"github.com/golang/glog"
)

type KeyAndMode struct {
	Key  string             `json:"-"`
	Mode *filemode.FileMode `json:"-"`
}

var keyAndModeRegexp = regexp.MustCompile(`^(.*):(0[0-7][0-7][0-7])$`)

func (k *KeyAndMode) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return serrors.InvalidValueErrorf(string(data), "expected string for key:mode")
	}

	matches := keyAndModeRegexp.FindStringSubmatch(str)
	if len(matches) == 0 {
		k.Key = str
		return nil
	}

	k.Key = matches[1]

	// The regexp should ensure that this always succeeds.
	i, err := strconv.ParseInt(matches[2], 8, 32)
	if err != nil {
		glog.V(0).Info("KeyAndMode regexp is matching non-integer file modes.")
		return serrors.InvalidValueErrorf(str, "expected integer for file mode in key:mode")
	}
	mode := filemode.FileMode(int32(i))
	k.Mode = &mode
	return nil
}

func (k KeyAndMode) MarshalJSON() ([]byte, error) {
	if k.Mode != nil {
		return json.Marshal(fmt.Sprintf("%s:0%o", k.Key, *k.Mode))
	}

	return json.Marshal(k.Key)
}

func NewKubeV1KeyToPath(items map[string]KeyAndMode) []v1.KeyToPath {
	if len(items) == 0 {
		return nil
	}

	kubeItems := []v1.KeyToPath{}
	for path, item := range items {
		kubeItems = append(kubeItems, v1.KeyToPath{
			Path: path,
			Key:  item.Key,
			Mode: filemode.ConvertFileModeToInt32Ptr(item.Mode),
		})
	}

	return kubeItems
}

func NewKeyToPathFromKubeVKeyToPathV1(kubeItems []v1.KeyToPath) map[string]KeyAndMode {
	if len(kubeItems) == 0 {
		return nil
	}

	items := map[string]KeyAndMode{}
	for _, item := range kubeItems {
		items[item.Path] = KeyAndMode{
			Key:  item.Key,
			Mode: filemode.NewFileModeFromKubeV1(item.Mode),
		}
	}

	return items
}
