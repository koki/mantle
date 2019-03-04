package pod

import (
	. "mantle/pkg/core/pod/podtemplate"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodCondition struct {
	LastProbeTime      metav1.Time      `json:"last_probe_time,omitempty"`
	LastTransitionTime metav1.Time      `json:"last_change,omitempty"`
	Msg                string           `json:"msg,omitempty"`
	Reason             string           `json:"reason,omitempty"`
	Status             ConditionStatus  `json:"status,omitempty"`
	Type               PodConditionType `json:"type,omitempty"`
}

type ConditionStatus int

const (
	ConditionStatusTrue ConditionStatus = iota
	ConditionStatusFalse
	ConditionStatusUnknown
	ConditionStatusNonTolerationse
	ConditionStatusNone
)

type PodPhase int

const (
	PodPhasePending PodPhase = iota
	PodPhaseRunning
	PodPhaseSucceeded
	PodPhaseFailed
	PodPhaseUnknown
	PodPhaseNone
)

type PodQOSClass int

const (
	PodQOSClassGuaranteed PodQOSClass = iota
	PodQOSClassBurstable
	PodQOSClassBestEffort
	PodQOSClassNone
)

// Pod defines a pod object
type Pod struct {
	Version string `json:"version,omitempty"`

	Conditions []PodCondition `json:"condition,omitempty"`
	NodeIP     string         `json:"node_ip,omitempty"`
	StartTime  *metav1.Time   `json:"start_time,omitempty"`
	Msg        string         `json:"msg,omitempty"`
	Phase      PodPhase       `json:"phase,omitempty"`
	IP         string         `json:"ip,omitempty"`
	QOS        PodQOSClass    `json:"qos,omitempty"`
	Reason     string         `json:"reason,omitempty"`

	PodTemplateMeta `json:",inline"`
	PodTemplate     `json:",inline"`
}
