package util

import (
	"k8s.io/api/core/v1"
)

type KeyAndMode struct {
	Key  string    `json:"-"`
	Mode *FileMode `json:"-"`
}

func NewKubeKeyToPathV1(items map[string]KeyAndMode) []v1.KeyToPath {
	if len(items) == 0 {
		return nil
	}

	kubeItems := []v1.KeyToPath{}
	for path, item := range items {
		kubeItems = append(kubeItems, v1.KeyToPath{
			Path: path,
			Key:  item.Key,
			Mode: ConvertFileModeToInt32Ptr(item.Mode),
		})
	}

	return kubeItems
}

func NewKeyToPathFromKubeKeyToPathV1(kubeItems []v1.KeyToPath) map[string]KeyAndMode {
	if len(kubeItems) == 0 {
		return nil
	}

	items := map[string]KeyAndMode{}
	for _, item := range kubeItems {
		items[item.Path] = KeyAndMode{
			Key:  item.Key,
			Mode: NewFileModeFromKubeV1(item.Mode),
		}
	}

	return items
}
