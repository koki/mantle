package port

import (
	"mantle/pkg/core/protocol"
)

type Port struct {
	Name          string
	Protocol      protocol.Protocol
	IP            string
	HostPort      string
	ContainerPort string
}
