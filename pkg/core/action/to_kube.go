package action

import (
	"fmt"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// ToKube will return a kubernetes Handler object of the api version provided
func (a *Action) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return a.toKubeV1()
	case "":
		return a.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Action: %s", version)
	}
}

func (a *Action) toKubeV1() (*v1.Handler, error) {
	handler := &v1.Handler{}

	switch a.ActionType {
	case ActionTypeCommand:
		handler.Exec = &v1.ExecAction{
			Command: a.Command,
		}

	case ActionTypeHTTP, ActionTypeHTTPS:
		var scheme v1.URIScheme

		if a.ActionType == ActionTypeHTTP {
			scheme = v1.URISchemeHTTP
		} else {
			scheme = v1.URISchemeHTTPS
		}

		port := intstr.FromString("80")
		if len(a.Port) > 0 {
			port = intstr.Parse(a.Port)
		}

		var headers []v1.HTTPHeader
		for _, header := range a.Headers {
			fields := strings.Split(header, ":")
			if len(fields) != 2 {
				return nil, serrors.InvalidInstanceErrorf(a, "unexpected HTTP Header %s", header)
			}
			kubeHeader := v1.HTTPHeader{
				Name:  fields[0],
				Value: fields[1],
			}
			headers = append(headers, kubeHeader)
		}

		handler.HTTPGet = &v1.HTTPGetAction{
			Scheme:      scheme,
			Path:        a.Path,
			Port:        port,
			Host:        a.Host,
			HTTPHeaders: headers,
		}

	case ActionTypeTCP:
		port := intstr.FromString("80")
		if len(a.Port) > 0 {
			port = intstr.Parse(a.Port)
		}
		handler.TCPSocket = &v1.TCPSocketAction{
			Host: a.Host,
			Port: port,
		}

	default:
		return nil, serrors.InvalidInstanceErrorf(a.ActionType, "unrecognized ActionType")
	}

	return handler, nil
}
