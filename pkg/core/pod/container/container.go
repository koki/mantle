package container

import (
	"mantle/pkg/core/action"
	"mantle/pkg/core/pod/container/env"
	"mantle/pkg/core/pod/container/port"
	"mantle/pkg/core/pod/container/probe"
	"mantle/pkg/core/pod/container/resources"
	"mantle/pkg/core/pod/container/volumemount"
	"mantle/pkg/core/selinux"
	"mantle/pkg/util/floatstr"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Container struct {
	Command              []string                  `json:"command,omitempty"`
	Args                 []floatstr.FloatOrString  `json:"args,omitempty"`
	Env                  []env.Env                 `json:"env,omitempty"`
	Image                string                    `json:"image"`
	Pull                 PullPolicy                `json:"pull,omitempty"`
	OnStart              *action.Action            `json:"on_start,omitempty"`
	PreStop              *action.Action            `json:"pre_stop,omitempty"`
	CPU                  *resources.CPU            `json:"cpu,omitempty"`
	Mem                  *resources.Mem            `json:"mem,omitempty"`
	Name                 string                    `json:"name,omitempty"`
	AddCapabilities      []string                  `json:"cap_add,omitempty"`
	DelCapabilities      []string                  `json:"cap_drop,omitempty"`
	Privileged           *bool                     `json:"privileged,omitempty"`
	AllowEscalation      *bool                     `json:"allow_escalation,omitempty"`
	RW                   *bool                     `json:"rw,omitempty"`
	RO                   *bool                     `json:"ro,omitempty"`
	ForceNonRoot         *bool                     `json:"force_non_root,omitempty"`
	UID                  *int64                    `json:"uid,omitempty"`
	GID                  *int64                    `json:"gid,omitempty"`
	SELinux              *selinux.SELinux          `json:"selinux,omitempty"`
	LivenessProbe        *probe.Probe              `json:"liveness_probe,omitempty"`
	ReadinessProbe       *probe.Probe              `json:"readiness_probe,omitempty"`
	Expose               []port.Port               `json:"expose,omitempty"`
	Stdin                bool                      `json:"stdin,omitempty"`
	StdinOnce            bool                      `json:"stdin_once,omitempty"`
	TTY                  bool                      `json:"tty,omitempty"`
	WorkingDir           string                    `json:"wd,omitempty"`
	TerminationMsgPath   string                    `json:"termination_msg_path,omitempty"`
	TerminationMsgPolicy TerminationMessagePolicy  `json:"termination_msg_policy,omitempty"`
	ContainerID          string                    `json:"container_id,omitempty"`
	ImageID              string                    `json:"image_id,omitempty"`
	Ready                bool                      `json:"ready,omitempty"`
	LastState            *ContainerState           `json:"last_state,omitempty"`
	CurrentState         *ContainerState           `json:"current_state,omitempty"`
	VolumeMounts         []volumemount.VolumeMount `json:"volume,omitempty"`
	Restarts             int32                     `json:"restarts,omitempty"`
}

type ContainerState struct {
	Waiting    *ContainerStateWaiting    `json:"waiting,omitempty"`
	Terminated *ContainerStateTerminated `json:"terminated,omitempty"`
	Running    *ContainerStateRunning    `json:"running,omitempty"`
}

type ContainerStateWaiting struct {
	Reason string `json:"reason,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

type ContainerStateRunning struct {
	StartTime metav1.Time `json:"start_time,omitempty"`
}

type ContainerStateTerminated struct {
	StartTime  metav1.Time `json:"start_time,omitempty"`
	FinishTime metav1.Time `json:"finish_time,omitempty"`
	Reason     string      `json:"reason,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	ExitCode   int32       `json:"exit_code,omitempty"`
	Signal     int32       `json:"signal,omitempty"`
}

type TerminationMessagePolicy int

const (
	TerminationMessageReadFile TerminationMessagePolicy = iota
	TerminationMessageFallbackToLogsOnError
	TerminationMessageDefault
)

type PullPolicy int

const (
	PullAlways PullPolicy = iota
	PullNever
	PullIfNotPresent
	PullDefault
)
