package pod

import (
	"fmt"
	"reflect"

	. "mantle/pkg/core/pod/container"
	. "mantle/pkg/core/pod/podtemplate"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"
)

// NewPodFromKubePod will create a new Pod object with
// the data from a provided kubernetes pod object
func NewPodFromKubePod(pod interface{}) (*Pod, error) {
	switch reflect.TypeOf(pod) {
	case reflect.TypeOf(v1.Pod{}):
		obj := pod.(v1.Pod)
		return fromKubePodV1(&obj)
	case reflect.TypeOf(&v1.Pod{}):
		return fromKubePodV1(pod.(*v1.Pod))
	default:
		return nil, fmt.Errorf("unknown pod version: %s", reflect.TypeOf(pod))
	}
}

func fromKubePodV1(pod *v1.Pod) (*Pod, error) {
	var err error
	mantlePod := &Pod{}

	mantlePod.Version = pod.APIVersion

	templateMeta, err := NewPodTemplateMetaFromKubeObjectMeta(pod.ObjectMeta)
	if err != nil {
		return nil, err
	}
	mantlePod.PodTemplateMeta = *templateMeta

	template, err := NewPodTemplateFromKubePodSpec(pod.Spec)
	if err != nil {
		return nil, err
	}
	mantlePod.PodTemplate = *template

	mantlePod.Msg = pod.Status.Message
	mantlePod.Reason = pod.Status.Reason
	phase, err := fromKubePodPhaseV1(pod.Status.Phase)
	if err != nil {
		return nil, err
	}
	mantlePod.Phase = phase
	mantlePod.IP = pod.Status.PodIP
	mantlePod.NodeIP = pod.Status.HostIP
	mantlePod.StartTime = pod.Status.StartTime

	qosClass, err := fromKubePodQOSClassV1(pod.Status.QOSClass)
	if err != nil {
		return nil, err
	}
	mantlePod.QOS = qosClass

	conditions, err := fromKubePodConditionV1(pod.Status.Conditions)
	if err != nil {
		return nil, err
	}
	mantlePod.Conditions = conditions

	fromKubeContainerStatusV1(pod.Status.InitContainerStatuses, pod.Status.ContainerStatuses, mantlePod.Containers)

	return mantlePod, nil
}

func fromKubePodPhaseV1(phase v1.PodPhase) (PodPhase, error) {
	switch phase {
	case "":
		return PodPhaseNone, nil

	case v1.PodPending:
		return PodPhasePending, nil

	case v1.PodRunning:
		return PodPhaseRunning, nil

	case v1.PodSucceeded:
		return PodPhaseSucceeded, nil

	case v1.PodFailed:
		return PodPhaseFailed, nil

	case v1.PodUnknown:
		return PodPhaseUnknown, nil

	default:
		return PodPhaseNone, serrors.InvalidInstanceError(phase)
	}

}

func fromKubePodQOSClassV1(class v1.PodQOSClass) (PodQOSClass, error) {
	switch class {
	case "":
		return PodQOSClassNone, nil

	case v1.PodQOSGuaranteed:
		return PodQOSClassGuaranteed, nil

	case v1.PodQOSBurstable:
		return PodQOSClassBurstable, nil

	case v1.PodQOSBestEffort:
		return PodQOSClassBestEffort, nil

	default:
		return PodQOSClassNone, serrors.InvalidInstanceError(class)
	}
}

func fromKubePodConditionV1(kubeConditions []v1.PodCondition) ([]PodCondition, error) {
	var conditions []PodCondition

	for _, kubeCondition := range kubeConditions {
		condition := PodCondition{
			Msg:                kubeCondition.Message,
			Reason:             kubeCondition.Reason,
			LastProbeTime:      kubeCondition.LastProbeTime,
			LastTransitionTime: kubeCondition.LastTransitionTime,
		}

		typ, err := fromKubePodConditionTypeV1(kubeCondition.Type)
		if err != nil {
			return nil, err
		}
		condition.Type = typ

		status, err := fromKubeConditionStatusV1(kubeCondition.Status)
		if err != nil {
			return nil, err
		}
		condition.Status = status

		conditions = append(conditions, condition)
	}

	return conditions, nil
}

func fromKubePodConditionTypeV1(condition v1.PodConditionType) (PodConditionType, error) {
	switch condition {
	case "":
		return PodConditionNone, nil

	case v1.PodScheduled:
		return PodConditionScheduled, nil

	case v1.PodReady:
		return PodConditionReady, nil

	case v1.PodInitialized:
		return PodConditionInitialized, nil

	case v1.PodReasonUnschedulable:
		return PodConditionReasonUnschedulable, nil

	default:
		return PodConditionNone, serrors.InvalidInstanceError(condition)
	}
}

func fromKubeConditionStatusV1(status v1.ConditionStatus) (ConditionStatus, error) {
	switch status {
	case "":
		return ConditionStatusNone, nil

	case v1.ConditionTrue:
		return ConditionStatusTrue, nil

	case v1.ConditionFalse:
		return ConditionStatusFalse, nil

	case v1.ConditionUnknown:
		return ConditionStatusUnknown, nil

	default:
		return ConditionStatusNone, serrors.InvalidInstanceError(status)
	}
}

func fromKubeContainerStatusV1(initContainerStatuses, containerStatuses []v1.ContainerStatus, containers []Container) {
	allContainerStatuses := append(initContainerStatuses, containerStatuses...)

	for _, status := range allContainerStatuses {
		for _, container := range containers {
			if container.Name == status.Name {
				container.Restarts = status.RestartCount
				container.Ready = status.Ready
				container.ImageID = status.ImageID
				container.ContainerID = status.ContainerID
				container.CurrentState = fromKubeContainerStateV1(status.State)
				container.LastState = fromKubeContainerStateV1(status.LastTerminationState)
			}
		}
	}
}

func fromKubeContainerStateV1(state v1.ContainerState) *ContainerState {
	s := &ContainerState{}

	if state.Waiting != nil {
		s.Waiting = &ContainerStateWaiting{
			Reason: state.Waiting.Reason,
			Msg:    state.Waiting.Message,
		}
	}
	if state.Running != nil {
		s.Running = &ContainerStateRunning{
			StartTime: state.Running.StartedAt,
		}
	}
	if state.Terminated != nil {
		s.Terminated = &ContainerStateTerminated{
			StartTime:  state.Terminated.StartedAt,
			FinishTime: state.Terminated.FinishedAt,
			Reason:     state.Terminated.Reason,
			Msg:        state.Terminated.Message,
			Signal:     state.Terminated.Signal,
			ExitCode:   state.Terminated.ExitCode,
		}
	}
	return s
}
