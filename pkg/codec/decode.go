package codec

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"mantle/internal/yaml"
	"mantle/pkg/core/configmap"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func Decode(input io.Reader) (io.Reader, error) {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	obj := map[string]interface{}{}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	kubeObj, err := ParseKubeNativeType(obj)
	if err != nil {
		//TBD: parse koki type
		return nil, err
	}

	switch kubeTypedObj := kubeObj.(type) {
	case *v1.ConfigMap:
		cm, err := configmap.NewConfigMapFromKubeConfigMap(kubeTypedObj)
		if err != nil {
			return nil, err
		}
		buf := &bytes.Buffer{}
		encoder := json.NewEncoder(buf)
		err = encoder.Encode(cm)
		return buf, err
	}
	return nil, errors.New("unreachable")
}

func ParseKubeNativeType(obj map[string]interface{}) (runtime.Object, error) {
	u := &unstructured.Unstructured{
		Object: obj,
	}

	typedObj, err := creator.New(u.GetObjectKind().GroupVersionKind())
	if err != nil {
		return nil, err
	}

	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj, typedObj); err != nil {
		return nil, err
	}
	return typedObj, nil
}
