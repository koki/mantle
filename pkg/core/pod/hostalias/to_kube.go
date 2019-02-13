package hostalias

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return a kubernetes host alias object of the api version provided
func (h *HostAlias) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return h.toKubeV1()
	case "":
		return h.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for HostAlias: %s", version)
	}
}

func (h *HostAlias) toKubeV1() (*v1.HostAlias, error) {
	hostAlias := v1.HostAlias{}

	hostAlias.IP = h.IP
	hostAlias.Hostnames = h.Hostnames

	return &hostAlias, nil
}
