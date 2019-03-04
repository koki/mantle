package pod

import (
	"fmt"
	"strings"

	"mantle/pkg/core/pod/podtemplate"

	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ToKube will return a kubernetes pod object of the api version
// type defined in the pod
func (pod *Pod) ToKube() (runtime.Object, error) {
	switch strings.ToLower(pod.Version) {
	case "v1":
		return pod.toKubeV1()
	case "":
		return pod.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Pod: %s", pod.Version)
	}
}

func (pod *Pod) toKubeV1() (*v1.Pod, error) {
	var err error
	kubePod := &v1.Pod{}

	kubePod.APIVersion = pod.Version
	kubePod.Kind = "Pod"

	meta, err := pod.PodTemplateMeta.ToKube(pod.Version)
	if err != nil {
		return nil, err
	}
	metaV1 := meta.(*metav1.ObjectMeta)
	kubePod.ObjectMeta = *metaV1

	spec, err := pod.PodTemplate.ToKube(pod.Version)
	if err != nil {
		return nil, err
	}
	podSpec := spec.(*v1.PodSpec)
	kubePod.Spec = *podSpec

	kubePod.Status = v1.PodStatus{}
	kubePod.Status.Phase = pod.toKubePodPhaseV1()
	kubePod.Status.Message = pod.Msg
	kubePod.Status.Reason = pod.Reason
	kubePod.Status.HostIP = pod.NodeIP
	kubePod.Status.PodIP = pod.IP
	kubePod.Status.QOSClass = pod.toKubeQOSClassV1()
	kubePod.Status.StartTime = pod.StartTime
	kubePod.Status.Conditions = pod.toKubePodConditionV1()

	var initContainerStatuses []v1.ContainerStatus
	for _, container := range pod.InitContainers {
		s, err := container.ToKubeStatus("v1")
		if err != nil {
			return nil, err
		}
		status := s.(v1.ContainerStatus)
		initContainerStatuses = append(initContainerStatuses, status)
	}
	kubePod.Status.InitContainerStatuses = initContainerStatuses

	var containerStatuses []v1.ContainerStatus
	for _, container := range pod.Containers {
		s, err := container.ToKubeStatus("v1")
		if err != nil {
			return nil, err
		}

		status := s.(v1.ContainerStatus)
		if status.ContainerID != "" {
			containerStatuses = append(containerStatuses, status)
		}
	}
	kubePod.Status.ContainerStatuses = containerStatuses

	return kubePod, nil
}

func (pod *Pod) toKubePodPhaseV1() v1.PodPhase {
	switch pod.Phase {
	case PodPhasePending:
		return v1.PodPending
	case PodPhaseRunning:
		return v1.PodRunning
	case PodPhaseSucceeded:
		return v1.PodSucceeded
	case PodPhaseFailed:
		return v1.PodFailed
	case PodPhaseUnknown:
		return v1.PodUnknown
	default:
		return ""
	}
}

func (pod *Pod) toKubeQOSClassV1() v1.PodQOSClass {
	switch pod.QOS {
	case PodQOSClassGuaranteed:
		return v1.PodQOSGuaranteed
	case PodQOSClassBurstable:
		return v1.PodQOSBurstable
	case PodQOSClassBestEffort:
		return v1.PodQOSBestEffort
	default:
		return ""
	}
}

func (pod *Pod) toKubePodConditionV1() []v1.PodCondition {
	var kubeConditions []v1.PodCondition

	for _, condition := range pod.Conditions {
		kubeCondition := v1.PodCondition{
			LastProbeTime:      condition.LastProbeTime,
			LastTransitionTime: condition.LastTransitionTime,
			Message:            condition.Msg,
			Reason:             condition.Reason,
		}

		kubeCondition.Type = podtemplate.ToKubePodConditionTypeV1(condition.Type)
		kubeCondition.Status = pod.toKubeConditionStatusV1(condition.Status)

		kubeConditions = append(kubeConditions, kubeCondition)
	}

	return kubeConditions
}

func (pod *Pod) toKubeConditionStatusV1(status ConditionStatus) v1.ConditionStatus {
	switch status {
	case ConditionStatusTrue:
		return v1.ConditionTrue
	case ConditionStatusFalse:
		return v1.ConditionFalse
	case ConditionStatusUnknown:
		return v1.ConditionUnknown
	default:
		return ""
	}
}
