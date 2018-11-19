package codec

import (
	"bytes"
	"encoding/json"
	"io"

	"mantle/pkg/core/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func Decode(input io.Reader) (io.Writer, error) {
	obj := map[string]interface{}{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	kubeObj, err := ParseKubeNativeType(obj)
	if err != nil {
		//TBD: parse koki type
		return nil, err
	}

	if kubeObj.GetObjectKind == "ConfigMap" {
		cm, err := v1.NewConfigMapFromKubeConfigMap(kubeObj)
		if err != nil {
			return nil, err
		}
		buf := &bytes.Buffer{}
		encoder := json.NewEncoder(buf)
		err := encoder.Encode(cm)
		return buf, err
	}
}

func ParseKubeNativeType(obj interface{}) (runtime.Object, error) {
	u := &unstructured.Unstructured{
		Object: obj,
	}

	typedObj, err := creator.New(u.GetObjectKind().GroupVersionKind())
	if err != nil {
		return nil, serrors.InvalidValueContextErrorf(err, u, "unsupported apiVersion/kind (is the manifest kube-native format?)")
	}

	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj, typedObj); err != nil {
		return nil, serrors.InvalidValueForTypeContextErrorf(err, obj, typedObj, "couldn't convert to typed kube obj")
	}
	return typedObj, nil
}
