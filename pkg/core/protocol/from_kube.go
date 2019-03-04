package protocol

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewProtocolFromKubeProtocol will create a new
// Protocol object with the data from a provided kubernetes
// Protocol object
func NewProtocolFromKubeProtocol(obj interface{}) (*Protocol, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.Protocol("")):
		return fromKubeProtocolV1(obj.(v1.Protocol))
	default:
		return nil, fmt.Errorf("unknown Protocol version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeProtocolV1(kubeProtocol v1.Protocol) (*Protocol, error) {
	var protocol Protocol

	switch kubeProtocol {
	case v1.Protocol("TCP"):
		protocol = ProtocolTCP
	case v1.Protocol("UDP"):
		protocol = ProtocolUDP
	}

	return &protocol, nil
}
