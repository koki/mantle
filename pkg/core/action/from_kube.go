package action

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"
)

// NewActionFromKubeHandler will create a new Action object with
// the data from a provided kubernetes Handler object
func NewActionFromKubeHandler(handler interface{}) (*Action, error) {
	switch reflect.TypeOf(handler) {
	case reflect.TypeOf(v1.Handler{}):
		obj := handler.(v1.Handler)
		return fromKubeHandlerV1(&obj)
	case reflect.TypeOf(&v1.Handler{}):
		return fromKubeHandlerV1(handler.(*v1.Handler))
	default:
		return nil, fmt.Errorf("unknown Handler version: %s", reflect.TypeOf(handler))
	}
}

func fromKubeHandlerV1(handler *v1.Handler) (*Action, error) {
	if handler == nil {
		return nil, nil
	}

	if handler.Exec != nil {
		return &Action{
			ActionType: ActionTypeCommand,
			Command:    handler.Exec.Command,
		}, nil
	}

	if handler.HTTPGet != nil {
		headers := []string{}

		for _, inHeader := range handler.HTTPGet.HTTPHeaders {
			outHeader := fmt.Sprintf("%s:%s", inHeader.Name, inHeader.Value)
			headers = append(headers, outHeader)
		}

		var actionType ActionType
		if handler.HTTPGet.Scheme == v1.URISchemeHTTP {
			actionType = ActionTypeHTTP
		} else {
			actionType = ActionTypeHTTPS
		}

		return &Action{
			ActionType: actionType,
			Path:       handler.HTTPGet.Path,
			Port:       handler.HTTPGet.Port.String(),
			Host:       handler.HTTPGet.Host,
			Headers:    headers,
		}, nil
	}

	if handler.TCPSocket != nil {
		return &Action{
			ActionType: ActionTypeTCP,
			Host:       handler.TCPSocket.Host,
			Port:       handler.TCPSocket.Port.String(),
		}, nil
	}

	return nil, nil
}
