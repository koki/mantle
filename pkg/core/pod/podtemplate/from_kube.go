package podtemplate

import (
	"fmt"
	"reflect"

	"mantle/pkg/core/pod/affinity"
	"mantle/pkg/core/pod/container"
	"mantle/pkg/core/pod/hostalias"
	"mantle/pkg/core/pod/toleration"
	"mantle/pkg/core/pod/volume"
	"mantle/pkg/core/selinux"

	"k8s.io/api/core/v1"

	serrors "github.com/koki/structurederrors"
)

// NewPodTemplateFromKubePodSpec will create a new
// PodTemplate object with the data from a provided kubernetes
// PodSpec object
func NewPodTemplateFromKubePodSpec(obj interface{}) (*PodTemplate, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.PodSpec{}):
		return fromKubePodSpecV1(obj.(v1.PodSpec))
	case reflect.TypeOf(&v1.PodSpec{}):
		o := obj.(*v1.PodSpec)
		return fromKubePodSpecV1(*o)
	default:
		return nil, fmt.Errorf("unknown PodSpec version: %s", reflect.TypeOf(obj))
	}
}

func fromKubePodSpecV1(kubeSpec v1.PodSpec) (*PodTemplate, error) {
	var err error

	mantlePod := &PodTemplate{}
	mantlePod.Volumes, err = fromKubeVolumesV1(kubeSpec.Volumes)
	if err != nil {
		return nil, serrors.ContextualizeErrorf(err, "pod volumes")
	}

	a, err := affinity.NewAffinityFromKubeAffinity(kubeSpec.Affinity)
	if err != nil {
		return nil, err
	}
	if a != nil {
		mantlePod.Affinity = *a
	}

	var initContainers []container.Container
	for _, kubeContainer := range kubeSpec.InitContainers {
		c, err := container.NewContainerFromKubeContainer(&kubeContainer)
		if err != nil {
			return nil, err
		}
		initContainers = append(initContainers, *c)
	}
	mantlePod.InitContainers = initContainers

	var containers []container.Container
	for _, kubeContainer := range kubeSpec.Containers {
		c, err := container.NewContainerFromKubeContainer(&kubeContainer)
		if err != nil {
			return nil, err
		}
		containers = append(containers, *c)
	}
	mantlePod.Containers = containers

	dnsPolicy, err := fromKubeDNSPolicyV1(kubeSpec.DNSPolicy)
	if err != nil {
		return nil, err
	}
	mantlePod.DNSPolicy = dnsPolicy

	aliases, err := fromKubeHostAliasesV1(kubeSpec.HostAliases)
	if err != nil {
		return nil, err
	}
	mantlePod.HostAliases = aliases

	mantlePod.HostMode = fromKubeHostModeV1(kubeSpec)
	mantlePod.Hostname = fromKubeHostnameV1(kubeSpec)
	mantlePod.Registries = fromKubeRegistriesV1(kubeSpec.ImagePullSecrets)

	restartPolicy, err := fromKubeRestartPolicyV1(kubeSpec.RestartPolicy)
	if err != nil {
		return nil, err
	}
	mantlePod.RestartPolicy = restartPolicy

	mantlePod.SchedulerName = kubeSpec.SchedulerName
	mantlePod.Account = kubeSpec.ServiceAccountName
	mantlePod.AutomountAccountToken = kubeSpec.AutomountServiceAccountToken

	tolerations, err := fromKubeTolerationsV1(kubeSpec.Tolerations)
	if err != nil {
		return nil, err
	}
	mantlePod.Tolerations = tolerations

	mantlePod.TerminationGracePeriod = kubeSpec.TerminationGracePeriodSeconds
	mantlePod.ActiveDeadline = kubeSpec.ActiveDeadlineSeconds
	mantlePod.Node = kubeSpec.NodeName
	mantlePod.PriorityClass = kubeSpec.PriorityClassName
	mantlePod.Priority = kubeSpec.Priority

	if kubeSpec.SecurityContext != nil {
		securityContext := kubeSpec.SecurityContext
		mantlePod.GIDs = securityContext.SupplementalGroups
		mantlePod.FSGID = securityContext.FSGroup
		for i := range mantlePod.Containers {
			container := &mantlePod.Containers[i]
			if container.SELinux == nil {
				sel, err := selinux.NewSELinuxFromKubeSELinuxOptions(securityContext.SELinuxOptions)
				if err != nil {
					return nil, err
				}
				container.SELinux = sel
			}
			if container.UID == nil {
				container.UID = securityContext.RunAsUser
			}
			if container.GID == nil {
				container.GID = securityContext.RunAsGroup
			}
			if container.ForceNonRoot == nil {
				container.ForceNonRoot = securityContext.RunAsNonRoot
			}
		}
	}

	return mantlePod, nil
}

func fromKubeVolumesV1(kubeVolumes []v1.Volume) (map[string]volume.Volume, error) {
	volumes := map[string]volume.Volume{}

	for _, kubeVolume := range kubeVolumes {
		name := kubeVolume.Name
		volume, err := volume.NewVolumeFromKubeVolume(kubeVolume)
		if err != nil {
			return nil, serrors.ContextualizeErrorf(err, "volume (%s)", name)
		}
		volumes[name] = *volume
	}

	return volumes, nil
}

func fromKubeDNSPolicyV1(dnsPolicy v1.DNSPolicy) (DNSPolicy, error) {
	switch dnsPolicy {
	case "":
		return DNSUnset, nil

	case v1.DNSClusterFirstWithHostNet:
		return DNSClusterFirstWithHostNet, nil

	case v1.DNSClusterFirst:
		return DNSClusterFirst, nil

	case v1.DNSDefault:
		return DNSDefault, nil

	case v1.DNSNone:
		return DNSNone, nil
	}

	return DNSUnset, serrors.InvalidInstanceError(dnsPolicy)
}

func fromKubeHostAliasesV1(kubeAliases []v1.HostAlias) ([]hostalias.HostAlias, error) {
	var aliases []hostalias.HostAlias

	for _, alias := range kubeAliases {
		a, err := hostalias.NewHostAliasFromKubeHostAlias(alias)
		if err != nil {
			return nil, err
		}
		aliases = append(aliases, *a)
	}

	return aliases, nil
}

func fromKubeHostModeV1(spec v1.PodSpec) []HostMode {
	var hostModes []HostMode

	if spec.HostNetwork {
		hostModes = append(hostModes, HostModeNet)
	}
	if spec.HostPID {
		hostModes = append(hostModes, HostModePID)
	}
	if spec.HostIPC {
		hostModes = append(hostModes, HostModeIPC)
	}

	return hostModes
}

func fromKubeHostnameV1(spec v1.PodSpec) string {
	hostName := ""

	if spec.Hostname != "" {
		hostName = fmt.Sprintf("%s", spec.Hostname)
	}

	// TODO: verify that .subdomain is a valid input. i.e. without hostname
	if spec.Subdomain != "" {
		hostName = fmt.Sprintf("%s.%s", spec.Subdomain, hostName)
	}

	return hostName
}

func fromKubeRegistriesV1(ref []v1.LocalObjectReference) []string {
	var registries []string

	for _, r := range ref {
		registries = append(registries, r.Name)
	}

	return registries
}

func fromKubeRestartPolicyV1(policy v1.RestartPolicy) (RestartPolicy, error) {
	switch policy {
	case "":
		return RestartPolicyDefault, nil

	case v1.RestartPolicyAlways:
		return RestartPolicyAlways, nil

	case v1.RestartPolicyOnFailure:
		return RestartPolicyOnFailure, nil

	case v1.RestartPolicyNever:
		return RestartPolicyNever, nil
	}

	return RestartPolicyDefault, serrors.InvalidInstanceError(policy)
}

func fromKubeTolerationsV1(tolerations []v1.Toleration) ([]toleration.Toleration, error) {
	var tols []toleration.Toleration

	for _, t := range tolerations {
		tol, err := toleration.NewTolerationFromKubeToleration(t)
		if err != nil {
			return nil, err
		}
		tols = append(tols, *tol)
	}

	return tols, nil
}
