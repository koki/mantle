package port

import (
	"fmt"
	"reflect"

	"mantle/pkg/core/protocol"

	"k8s.io/api/core/v1"
)

// NewPortFromKubeContainerPort will create a new
// Port object with the data from a provided kubernetes
// ContainerPort object
func NewPortFromKubeContainerPort(obj interface{}) (*Port, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.ContainerPort{}):
		return fromKubeContainerPortV1(obj.(v1.ContainerPort))
	case reflect.TypeOf(&v1.ContainerPort{}):
		o := obj.(*v1.ContainerPort)
		return fromKubeContainerPortV1(*o)
	default:
		return nil, fmt.Errorf("unknown ContainerPort version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeContainerPortV1(kubePort v1.ContainerPort) (*Port, error) {
	port := &Port{
		Name: kubePort.Name,
		IP:   kubePort.HostIP,
	}

	protocol, err := protocol.NewProtocolFromKubeProtocol(kubePort.Protocol)
	if err != nil {
		return nil, err
	}
	port.Protocol = *protocol

	if kubePort.HostPort != 0 {
		port.HostPort = fmt.Sprintf("%d", kubePort.HostPort)
	}

	if kubePort.ContainerPort != 0 {
		port.ContainerPort = fmt.Sprintf("%d", kubePort.ContainerPort)
	}

	return port, nil
}
