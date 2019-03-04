package port

import (
	"fmt"
	"strconv"
	"strings"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes ContainerPort object of the api version provided
func (p *Port) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return p.toKubeV1()
	case "":
		return p.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Port: %s", version)
	}
}

func (p *Port) toKubeV1() (*v1.ContainerPort, error) {
	var err error
	kubePort := v1.ContainerPort{}

	kubePort.Name = p.Name
	protocol, err := p.Protocol.ToKube("v1")
	if err != nil {
		return nil, err
	}
	kubePort.Protocol = protocol.(v1.Protocol)

	kubePort.HostPort, err = p.hostPortInt()
	if err != nil {
		return nil, err
	}

	kubePort.ContainerPort, err = p.containerPortInt()
	if err != nil {
		return nil, err
	}

	return &kubePort, nil
}

func (p *Port) hostPortInt() (int32, error) {
	if len(p.HostPort) > 0 {
		hostPort, err := strconv.ParseInt(p.HostPort, 10, 32)
		if err != nil {
			return 0, serrors.InvalidInstanceContextErrorf(err, p, "HostPort should be an int")
		}

		return int32(hostPort), nil
	}

	return 0, nil
}

func (p *Port) containerPortInt() (int32, error) {
	if len(p.ContainerPort) > 0 {
		containerPort, err := strconv.ParseInt(p.ContainerPort, 10, 32)
		if err != nil {
			return 0, serrors.InvalidInstanceContextErrorf(err, p, "ContainerPort should be an int")
		}

		return int32(containerPort), nil
	}

	return 0, nil
}

/*
$protocol://$ip:$host_port:$container_port

expose:
  - 8080:80
  - UDP://127.0.0.1:8080:80
  - 10.10.0.53:8081:9090
  - port_name: 192.168.1.2:8090:80
*/

/*
var protocolPortRegexp = regexp.MustCompile(`^(udp|tcp)://([0-9.:]*)$`)

func (p *Port) InitFromString(str string) error {
	matches := protocolPortRegexp.FindStringSubmatch(str)
	if len(matches) > 0 {
		p.Protocol = Protocol(matches[1])
		str = matches[2]
	} else {
		p.Protocol = ProtocolTCP
	}

	segments := strings.Split(str, ":")
	parseIndex := 0

	ip := net.ParseIP(segments[parseIndex])
	if ip != nil {
		p.IP = segments[parseIndex]
		parseIndex++
	}

	remaining := len(segments) - parseIndex
	if remaining == 2 {
		p.HostPort = segments[parseIndex]
		p.ContainerPort = segments[parseIndex+1]
		return nil
	}
	if remaining == 1 {
		p.ContainerPort = segments[parseIndex]
		return nil
	}

	return serrors.InvalidInstanceErrorf(p, "couldn't parse (%s)", str)
}

func appendColonSegment(str, seg string) string {
	if len(str) == 0 {
		return seg
	}

	return fmt.Sprintf("%s:%s", str, seg)
}

func (p *Port) ToString() (string, error) {
	str := ""
	if len(p.IP) > 0 {
		str = appendColonSegment(str, p.IP)
	}

	if len(p.HostPort) > 0 {
		str = appendColonSegment(str, p.HostPort)
	}

	if len(p.ContainerPort) > 0 {
		str = appendColonSegment(str, p.ContainerPort)
	}

	if len(p.Protocol) == 0 || p.Protocol == ProtocolTCP {
		// No need to specify protocol
		return str, nil
	}

	return fmt.Sprintf("%s://%s", p.Protocol, str), nil
}
*/
