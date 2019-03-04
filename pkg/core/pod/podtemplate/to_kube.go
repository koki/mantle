package podtemplate

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	serrors "github.com/koki/structurederrors"
)

// ToKube converts the PodTemplate object to a
// kubernetes pod spec object of the version specified
func (pt *PodTemplate) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return pt.toKubeV1(version)
	case "":
		return pt.toKubeV1(version)
	default:
		return nil, fmt.Errorf("unsupported api version for PodTemplate: %s", version)
	}
}

func (pt *PodTemplate) toKubeV1(apiVersion string) (*v1.PodSpec, error) {
	var err error

	spec := v1.PodSpec{}

	spec.Volumes, err = pt.toKubeVolumesV1()
	if err != nil {
		return nil, err
	}

	fields := strings.SplitN(pt.Hostname, ".", 2)
	if len(fields) == 1 {
		spec.Hostname = pt.Hostname
	} else {
		spec.Hostname = fields[0]
		spec.Subdomain = fields[1]
	}

	var initContainers []v1.Container
	for _, container := range pt.InitContainers {
		kc, err := container.ToKube("v1")
		if err != nil {
			return nil, err
		}
		kubeContainer := kc.(v1.Container)
		initContainers = append(initContainers, kubeContainer)
	}
	spec.InitContainers = initContainers

	var kubeContainers []v1.Container
	for _, container := range pt.Containers {
		kc, err := container.ToKube("v1")
		if err != nil {
			return nil, err
		}
		kubeContainer := kc.(v1.Container)
		kubeContainers = append(kubeContainers, kubeContainer)
	}
	spec.Containers = kubeContainers

	hostAliases, err := pt.toKubeHostAliasesV1()
	if err != nil {
		return nil, err
	}
	spec.HostAliases = hostAliases

	restartPolicy, err := pt.toKubeRestartPolicyV1()
	if err != nil {
		return nil, err
	}
	spec.RestartPolicy = restartPolicy

	affinity, err := pt.Affinity.ToKube("v1")
	if err != nil {
		return nil, err
	}
	spec.Affinity = affinity.(*v1.Affinity)

	spec.TerminationGracePeriodSeconds = pt.TerminationGracePeriod
	spec.ActiveDeadlineSeconds = pt.ActiveDeadline

	dnsPolicy, err := pt.toKubeDNSPolicyV1()
	if err != nil {
		return nil, err
	}
	spec.DNSPolicy = dnsPolicy

	spec.NodeSelector = pt.NodeSelector
	spec.ServiceAccountName = pt.Account
	spec.AutomountServiceAccountToken = pt.AutomountAccountToken
	spec.NodeName = pt.Node

	net, pid, ipc, err := pt.toKubeHostModesV1()
	if err != nil {
		return nil, err
	}
	spec.HostNetwork = net
	spec.HostPID = pid
	spec.HostIPC = ipc
	spec.ShareProcessNamespace = pt.ShareNamespace
	spec.ImagePullSecrets = pt.toKubeRegistriesV1()
	spec.SchedulerName = pt.SchedulerName

	tolerations, err := pt.toKubeTolerationsV1()
	if err != nil {
		return nil, err
	}
	spec.Tolerations = tolerations

	if pt.FSGID != nil || pt.GIDs != nil {
		spec.SecurityContext = &v1.PodSecurityContext{}
		spec.SecurityContext.FSGroup = pt.FSGID
		spec.SecurityContext.SupplementalGroups = pt.GIDs
	}

	spec.Priority = pt.Priority
	spec.PriorityClassName = pt.PriorityClass
	spec.DNSConfig = pt.toKubeDNSConfigV1()

	spec.ReadinessGates = pt.toKubePodReadinessGatesV1()
	spec.RuntimeClassName = pt.RuntimeClass
	spec.EnableServiceLinks = pt.ServiceLinks

	return &spec, nil
}

func (pt *PodTemplate) toKubeVolumesV1() ([]v1.Volume, error) {
	kubeVolumes := []v1.Volume{}

	for name, vol := range pt.Volumes {
		v, err := vol.ToKube("v1")
		if err != nil {
			return nil, err
		}
		kubeVol := v.(*v1.Volume)
		kubeVol.Name = name
		kubeVolumes = append(kubeVolumes, *kubeVol)
	}

	return kubeVolumes, nil
}

func (pt *PodTemplate) toKubeHostAliasesV1() ([]v1.HostAlias, error) {
	var hostAliases []v1.HostAlias

	for _, alias := range pt.HostAliases {
		hostAlias, err := alias.ToKube("v1")
		if err != nil {
			return nil, err
		}

		ha := hostAlias.(*v1.HostAlias)
		hostAliases = append(hostAliases, *ha)
	}
	return hostAliases, nil
}

func (pt *PodTemplate) toKubeRestartPolicyV1() (v1.RestartPolicy, error) {
	switch pt.RestartPolicy {
	case RestartPolicyDefault:
		return "", nil

	case RestartPolicyAlways:
		return v1.RestartPolicyAlways, nil

	case RestartPolicyOnFailure:
		return v1.RestartPolicyOnFailure, nil

	case RestartPolicyNever:
		return v1.RestartPolicyNever, nil
	}

	return "", serrors.InvalidInstanceError(pt.RestartPolicy)
}

func (pt *PodTemplate) toKubeDNSPolicyV1() (v1.DNSPolicy, error) {
	switch pt.DNSPolicy {
	case DNSClusterFirstWithHostNet:
		return v1.DNSClusterFirstWithHostNet, nil

	case DNSClusterFirst:
		return v1.DNSClusterFirst, nil

	case DNSDefault:
		return v1.DNSDefault, nil

	case DNSNone:
		return v1.DNSNone, nil

	case DNSUnset:
		return "", nil
	}

	return "", serrors.InvalidInstanceError(pt.DNSPolicy)
}

func (pt *PodTemplate) toKubeHostModesV1() (net bool, pid bool, ipc bool, err error) {
	for _, mode := range pt.HostMode {
		switch mode {
		case HostModeNet:
			net = true
		case HostModePID:
			pid = true
		case HostModeIPC:
			ipc = true
		default:
			return false, false, false, serrors.InvalidInstanceError(mode)
		}
	}

	return net, pid, ipc, nil
}

func (pt *PodTemplate) toKubeRegistriesV1() []v1.LocalObjectReference {
	var kubeRegistries []v1.LocalObjectReference

	for _, reg := range pt.Registries {
		ref := v1.LocalObjectReference{
			Name: reg,
		}
		kubeRegistries = append(kubeRegistries, ref)
	}

	return kubeRegistries
}

func (pt *PodTemplate) toKubeTolerationsV1() ([]v1.Toleration, error) {
	var tolerations []v1.Toleration

	for _, t := range pt.Tolerations {
		tol, err := t.ToKube("v1")
		if err != nil {
			return nil, err
		}

		tolV1 := tol.(*v1.Toleration)
		tolerations = append(tolerations, *tolV1)
	}

	return tolerations, nil
}

func (pt *PodTemplate) toKubeDNSConfigV1() *v1.PodDNSConfig {
	if len(pt.Nameservers) == 0 && len(pt.SearchDomains) == 0 && len(pt.ResolverOptions) == 0 {
		return nil
	}

	options := []v1.PodDNSConfigOption{}
	for _, opt := range pt.ResolverOptions {
		o := v1.PodDNSConfigOption{
			Name:  opt.Name,
			Value: opt.Value,
		}
		options = append(options, o)
	}

	return &v1.PodDNSConfig{
		Nameservers: pt.Nameservers,
		Searches:    pt.SearchDomains,
		Options:     options,
	}
}

func (pt *PodTemplate) toKubePodReadinessGatesV1() []v1.PodReadinessGate {
	var readinessGates []v1.PodReadinessGate

	if len(pt.Gates) > 0 {
		readinessGates = []v1.PodReadinessGate{}
	}

	for _, gate := range pt.Gates {
		podGate := v1.PodReadinessGate{}

		condition := ToKubePodConditionTypeV1(gate)
		if len(condition) > 0 {
			podGate.ConditionType = condition
			readinessGates = append(readinessGates, podGate)
		}
	}

	return readinessGates
}

func ToKubePodConditionTypeV1(cond PodConditionType) v1.PodConditionType {
	switch cond {
	case PodConditionScheduled:
		return v1.PodScheduled

	case PodConditionReady:
		return v1.PodReady

	case PodConditionInitialized:
		return v1.PodInitialized

	case PodConditionReasonUnschedulable:
		return v1.PodReasonUnschedulable

	case PodConditionContainersReady:
		return v1.ContainersReady

	default:
		return ""
	}
}
