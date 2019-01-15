package configmap

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes volume object of the
// api version type defined in the object
func (cm *ConfigMap) ToKube() (runtime.Object, error) {
	switch strings.ToLower(cm.Version) {
	case "v1":
		return cm.toKubeV1()
	case "":
		return cm.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for config map: %s", cm.Version)
	}
}

func (cm *ConfigMap) toKubeV1() (*v1.ConfigMap, error) {
	kubeConfigMap := &v1.ConfigMap{}

	kubeConfigMap.Name = cm.Name
	kubeConfigMap.Namespace = cm.Namespace
	kubeConfigMap.APIVersion = cm.Version
	kubeConfigMap.ClusterName = cm.Cluster
	kubeConfigMap.Kind = "ConfigMap"
	kubeConfigMap.Labels = cm.Labels
	kubeConfigMap.Annotations = cm.Annotations
	kubeConfigMap.Data = cm.Data
	kubeConfigMap.BinaryData = cm.BinaryData

	return kubeConfigMap, nil
}