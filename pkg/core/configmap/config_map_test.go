package configmap

import (
	"reflect"
	"testing"

	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var mappings = map[string]string{
	"Name":        "Name",
	"Namespace":   "Namespace",
	"Version":     "APIVersion",
	"Cluster":     "ClusterName",
	"Labels":      "Labels",
	"Annotations": "Annotations",
	"Data":        "Data",
	"BinaryData":  "BinaryData",
}

func TestNewConfigMapFromKubeConfigMap(t *testing.T) {
	testcases := []struct {
		description string
		obj         interface{}
	}{
		{
			description: "v1 config map object",
			obj:         v1.ConfigMap{},
		},
		{
			description: "v1 config map pointer",
			obj:         &v1.ConfigMap{},
		},
	}

	for _, tc := range testcases {
		obj, _ := NewConfigMapFromKubeConfigMap(tc.obj)
		expectedObj := reflect.TypeOf(&ConfigMap{})
		objType := reflect.TypeOf(obj)
		if expectedObj != objType {
			t.Errorf("expected %s got %s", expectedObj, objType)
		}
	}
}

func TestFromKubeV1(t *testing.T) {
	v1CM := v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        "testCM",
			Namespace:   "testNS",
			ClusterName: "testCluster",
			Labels:      map[string]string{"label1": "test1", "label2": "test2"},
			Annotations: map[string]string{"ann1": "test1", "ann2": "test2"},
		},
		Data:       map[string]string{"field1": "data1", "field2": "data2"},
		BinaryData: map[string][]byte{"bfield1": []byte("bdata1")},
	}

	cm, _ := fromKubeV1(&v1CM)
	if cm.Name != v1CM.Name {
		t.Errorf("incorrect name, expected %s got %s", v1CM.Name, cm.Name)
	}

	for name, v1Name := range mappings {
		value := reflect.ValueOf(cm).Elem().FieldByName(name).Interface()
		v1Value := reflect.ValueOf(v1CM).FieldByName(v1Name).Interface()
		if !reflect.DeepEqual(value, v1Value) {
			t.Errorf("incorrect %s, expected %s, got %s", name, v1Value, value)
		}
	}
}

func TestToKube(t *testing.T) {
	testcases := []struct {
		description string
		version     string
		expectedObj interface{}
	}{
		{
			description: "v1 api version",
			version:     "v1",
			expectedObj: &v1.ConfigMap{},
		},
		{
			description: "empty api version",
			version:     "",
			expectedObj: &v1.ConfigMap{},
		},
		{
			description: "unknown api version",
			version:     "unknown",
			expectedObj: nil,
		},
	}

	for _, tc := range testcases {
		cm := ConfigMap{
			Version: tc.version,
		}
		kubeObj, err := cm.ToKube()
		kubeType := reflect.TypeOf(kubeObj)
		expectedType := reflect.TypeOf(tc.expectedObj)
		if kubeType != expectedType {
			t.Errorf("wrong api version, got %s expected %s", kubeType, expectedType)
		}
		if tc.expectedObj == nil && err == nil {
			t.Errorf("no error returned")
		}
	}
}

func TestToKubeV1(t *testing.T) {
	cm := ConfigMap{
		Version:     "v1",
		Name:        "testCM",
		Namespace:   "testNS",
		Cluster:     "testCluster",
		Labels:      map[string]string{"label1": "test1", "label2": "test2"},
		Annotations: map[string]string{"ann1": "test1", "ann2": "test2"},
		Data:        map[string]string{"field1": "data1", "field2": "data2"},
		BinaryData:  map[string][]byte{"bfield1": []byte("bdata1")},
	}

	kubeObj, _ := cm.toKubeV1()
	for name, v1Name := range mappings {
		value := reflect.ValueOf(cm).FieldByName(name).Interface()
		v1Value := reflect.ValueOf(kubeObj).Elem().FieldByName(v1Name).Interface()
		if !reflect.DeepEqual(value, v1Value) {
			t.Errorf("incorrect %s, expected %s, got %s", v1Name, value, v1Value)
		}
	}
}
