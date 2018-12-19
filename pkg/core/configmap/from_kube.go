package configmap

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewConfigMapFromKubeConfigMap will create a new ConfigMap object with
// the data from a provided kubernetes config map object
func NewConfigMapFromKubeConfigMap(cm interface{}) (*ConfigMap, error) {
	switch reflect.TypeOf(cm) {
	case reflect.TypeOf(v1.ConfigMap{}):
		obj := cm.(v1.ConfigMap)
		return fromKubeConfigMapV1(&obj)
	case reflect.TypeOf(&v1.ConfigMap{}):
		return fromKubeConfigMapV1(cm.(*v1.ConfigMap))
	default:
		return nil, fmt.Errorf("unknown ConfigMap version: %s", reflect.TypeOf(cm))
	}
}

func fromKubeConfigMapV1(kubeConfigMap *v1.ConfigMap) (*ConfigMap, error) {
	cm := &ConfigMap{
		Name:        kubeConfigMap.Name,
		Namespace:   kubeConfigMap.Namespace,
		Version:     kubeConfigMap.APIVersion,
		Cluster:     kubeConfigMap.ClusterName,
		Labels:      kubeConfigMap.Labels,
		Annotations: kubeConfigMap.Annotations,
		Data:        kubeConfigMap.Data,
		BinaryData:  kubeConfigMap.BinaryData,
	}

	return cm, nil
}
