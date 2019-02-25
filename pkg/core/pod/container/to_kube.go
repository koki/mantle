package container

import (
	"fmt"
	"reflect"
	"strings"

	"mantle/pkg/util"

	serrors "github.com/koki/structurederrors"

	"k8s.io/api/core/v1"

	"github.com/imdario/mergo"
)

// ToKube will return a kubernetes container object of the api version provided
func (c *Container) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return c.toKubeV1()
	case "":
		return c.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Container: %s", version)
	}
}

func (c *Container) toKubeV1() (v1.Container, error) {
	kubeContainer := v1.Container{}

	kubeContainer.Name = c.Name
	kubeContainer.Args = c.toKubeContainerArgsV1()
	kubeContainer.Command = c.Command
	kubeContainer.Image = c.Image
	kubeContainer.WorkingDir = c.WorkingDir

	kubeContainerPorts, err := c.toKubeContainerPortV1()
	if err != nil {
		return v1.Container{}, err
	}
	kubeContainer.Ports = kubeContainerPorts

	envs, envFroms, err := c.toKubeEnvVarV1()
	if err != nil {
		return v1.Container{}, err
	}
	kubeContainer.Env = envs
	kubeContainer.EnvFrom = envFroms

	resources, err := c.toKubeResourcesV1()
	if err != nil {
		return v1.Container{}, err
	}
	if resources != nil {
		kubeContainer.Resources = *resources
	}

	if c.LivenessProbe != nil {
		livenessProbe, err := c.LivenessProbe.ToKube("v1")
		if err != nil {
			return v1.Container{}, err
		}
		kubeContainer.LivenessProbe = livenessProbe.(*v1.Probe)
	}

	if c.ReadinessProbe != nil {
		readinessProbe, err := c.ReadinessProbe.ToKube("v1")
		if err != nil {
			return v1.Container{}, err
		}
		kubeContainer.ReadinessProbe = readinessProbe.(*v1.Probe)
	}

	kubeContainer.TerminationMessagePath = c.TerminationMsgPath
	kubeContainer.TerminationMessagePolicy = c.toKubeTerminationMsgPolicyV1()
	kubeContainer.ImagePullPolicy = c.toKubePullPolicyV1()
	vm, err := c.toKubeVolumeMountV1()
	if err != nil {
		return v1.Container{}, err
	}
	kubeContainer.VolumeMounts = vm

	kubeContainer.Stdin = c.Stdin
	kubeContainer.StdinOnce = c.StdinOnce
	kubeContainer.TTY = c.TTY

	lc, err := c.toKubeLifecycleV1()
	if err != nil {
		return v1.Container{}, err
	}
	kubeContainer.Lifecycle = lc

	sc, err := c.toKubeSecurityContextV1()
	if err != nil {
		return v1.Container{}, err
	}
	kubeContainer.SecurityContext = sc

	return kubeContainer, nil
}

func (c *Container) toKubeContainerArgsV1() []string {
	if c.Args == nil {
		return nil
	}
	kubeArgs := make([]string, len(c.Args))
	for i, arg := range c.Args {
		kubeArgs[i] = arg.String()
	}

	return kubeArgs
}

func (c *Container) toKubeContainerPortV1() ([]v1.ContainerPort, error) {
	var kubeContainerPorts []v1.ContainerPort

	for _, port := range c.Expose {
		kubePort, err := port.ToKube("v1")
		if err != nil {
			return nil, err
		}
		port := kubePort.(*v1.ContainerPort)
		kubeContainerPorts = append(kubeContainerPorts, *port)
	}
	return kubeContainerPorts, nil
}

func (c *Container) toKubeEnvVarV1() ([]v1.EnvVar, []v1.EnvFromSource, error) {
	var envVars []v1.EnvVar
	var envsFromSource []v1.EnvFromSource

	for _, e := range c.Env {
		envVar, envFromSrc, err := e.ToKube("v1")
		if err != nil {
			return nil, nil, err
		}

		if !reflect.ValueOf(envVar).IsNil() {
			e := envVar.(*v1.EnvVar)
			envVars = append(envVars, *e)
		}

		if !reflect.ValueOf(envFromSrc).IsNil() {
			e := envVar.(*v1.EnvFromSource)
			envsFromSource = append(envsFromSource, *e)
		}
	}

	return envVars, envsFromSource, nil
}

func (c *Container) toKubeResourcesV1() (*v1.ResourceRequirements, error) {
	limits := v1.ResourceList{}
	requests := v1.ResourceList{}
	requirements := v1.ResourceRequirements{
		Limits:   limits,
		Requests: requests,
	}

	if c.CPU != nil {
		cpuResources, err := c.CPU.ToKube("v1")
		if err != nil {
			return nil, err
		}

		if !reflect.ValueOf(cpuResources).IsNil() {
			res := cpuResources.(*v1.ResourceRequirements)
			err := mergo.Merge(&limits, res.Limits)
			if err != nil {
				return nil, err
			}
			err = mergo.Merge(&requests, res.Requests)
			if err != nil {
				return nil, err
			}
		}
	}

	if c.Mem != nil {
		memResources, err := c.Mem.ToKube("v1")
		if err != nil {
			return nil, err
		}

		if !reflect.ValueOf(memResources).IsNil() {
			res := memResources.(*v1.ResourceRequirements)
			err := mergo.Merge(&limits, res.Limits)
			if err != nil {
				return nil, err
			}
			err = mergo.Merge(&requests, res.Requests)
			if err != nil {
				return nil, err
			}
		}
	}

	if len(limits) > 0 || len(requests) > 0 {
		return &requirements, nil
	}

	return nil, nil
}

func (c *Container) toKubeTerminationMsgPolicyV1() v1.TerminationMessagePolicy {
	if c.TerminationMsgPolicy == TerminationMessageReadFile {
		return v1.TerminationMessageReadFile
	}

	if c.TerminationMsgPolicy == TerminationMessageFallbackToLogsOnError {
		return v1.TerminationMessageFallbackToLogsOnError
	}

	return ""
}

func (c *Container) toKubePullPolicyV1() v1.PullPolicy {
	switch c.Pull {
	case PullAlways:
		return v1.PullAlways

	case PullNever:
		return v1.PullNever

	case PullIfNotPresent:
		return v1.PullIfNotPresent

	default:
		return ""
	}
}

func (c *Container) toKubeVolumeMountV1() ([]v1.VolumeMount, error) {
	var kubeMounts []v1.VolumeMount

	for _, mount := range c.VolumeMounts {
		m, err := mount.ToKube("v1")
		if err != nil {
			return nil, err
		}
		mountV1 := m.(*v1.VolumeMount)
		kubeMounts = append(kubeMounts, *mountV1)
	}

	return kubeMounts, nil
}

func (c *Container) toKubeLifecycleV1() (*v1.Lifecycle, error) {
	var lc *v1.Lifecycle
	var kubeOnStart *v1.Handler
	var kubePreStop *v1.Handler

	if c.OnStart != nil {
		kos, err := c.OnStart.ToKube("v1")
		if err != nil {
			return nil, err
		}
		kubeOnStart = kos.(*v1.Handler)
	}

	if c.PreStop != nil {
		kps, err := c.PreStop.ToKube("v1")
		if err != nil {
			return nil, err
		}
		kubePreStop = kps.(*v1.Handler)
	}

	if kubeOnStart != nil || kubePreStop != nil {
		lc = &v1.Lifecycle{
			PostStart: kubeOnStart,
			PreStop:   kubePreStop,
		}
	}

	return lc, nil
}

func (c *Container) toKubeSecurityContextV1() (*v1.SecurityContext, error) {
	sc := &v1.SecurityContext{}

	var mark bool

	if c.Privileged != nil {
		sc.Privileged = c.Privileged
		mark = true
	}

	if c.AllowEscalation != nil {
		sc.AllowPrivilegeEscalation = c.AllowEscalation
		mark = true
	}

	if c.RO != nil || c.RW != nil {
		ro := util.FromBoolPtr(c.RO)
		rw := util.FromBoolPtr(c.RW)

		if !((!ro && rw) || (!rw && ro)) {
			return nil, serrors.InvalidInstanceErrorf(c, "conflicting value (Read Only) %v and (ReadWrite) %v", ro, rw)
		}

		sc.ReadOnlyRootFilesystem = &ro
		mark = true
	}

	if c.ForceNonRoot != nil {
		sc.RunAsNonRoot = c.ForceNonRoot
		mark = true
	}

	if c.UID != nil {
		sc.RunAsUser = c.UID
		mark = true
	}

	if c.GID != nil {
		sc.RunAsGroup = c.GID
		mark = true
	}

	if c.AddCapabilities != nil || c.DelCapabilities != nil {
		caps := &v1.Capabilities{}
		var capMark bool
		for _, capability := range c.AddCapabilities {
			caps.Add = append(caps.Add, v1.Capability(capability))
			capMark = true
		}

		for _, capability := range c.DelCapabilities {
			caps.Drop = append(caps.Drop, v1.Capability(capability))
			capMark = true
		}

		if capMark {
			sc.Capabilities = caps
			mark = true
		}
	}

	if c.SELinux != nil {
		sel, err := c.SELinux.ToKube("v1")
		if err != nil {
			return nil, err
		}
		sc.SELinuxOptions = sel.(*v1.SELinuxOptions)
		mark = true
	}

	if !mark {
		return nil, nil
	}

	return sc, nil
}

// ToKube will return a kubernetes container status object of the api version provided
func (c *Container) ToKubeStatus(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return c.toKubeStatusV1()
	case "":
		return c.toKubeStatusV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Container: %s", version)
	}
}

func (c *Container) toKubeStatusV1() (v1.ContainerStatus, error) {
	var status v1.ContainerStatus

	status.ContainerID = c.ContainerID
	status.ImageID = c.ImageID
	status.RestartCount = c.Restarts
	status.Ready = c.Ready
	status.State = c.toKubeStateV1(c.CurrentState)
	status.LastTerminationState = c.toKubeStateV1(c.LastState)

	return status, nil
}

func (c *Container) toKubeStateV1(state *ContainerState) v1.ContainerState {
	containerState := v1.ContainerState{}

	if state == nil {
		return containerState
	}

	if state.Waiting != nil {
		containerState.Waiting = &v1.ContainerStateWaiting{
			Reason:  state.Waiting.Reason,
			Message: state.Waiting.Msg,
		}
	}

	if state.Running != nil {
		containerState.Running = &v1.ContainerStateRunning{
			StartedAt: state.Running.StartTime,
		}
	}

	if state.Terminated != nil {
		containerState.Terminated = &v1.ContainerStateTerminated{
			StartedAt:  state.Terminated.StartTime,
			FinishedAt: state.Terminated.FinishTime,
			Reason:     state.Terminated.Reason,
			Message:    state.Terminated.Msg,
			Signal:     state.Terminated.Signal,
			ExitCode:   state.Terminated.ExitCode,
		}
	}

	return containerState
}
