package protocol

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes protocol object of the api version provided
func (p *Protocol) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return p.toKubeV1()
	case "":
		return p.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Protocol: %s", version)
	}
}

func (p *Protocol) toKubeV1() (v1.Protocol, error) {
	var protocol v1.Protocol

	switch *p {
	case ProtocolTCP:
		protocol = v1.Protocol("TCP")
	case ProtocolUDP:
		protocol = v1.Protocol("UDP")
	}

	return protocol, nil
}
