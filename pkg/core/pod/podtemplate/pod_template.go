package podtemplate

import (
	"mantle/pkg/core/pod/affinity"
	"mantle/pkg/core/pod/container"
	"mantle/pkg/core/pod/hostalias"
	"mantle/pkg/core/pod/toleration"
	"mantle/pkg/core/pod/volume"
)

// PodTemplate defines attributes for a pod
type PodTemplate struct {
	Volumes                map[string]volume.Volume `json:"volumes,omitempty"`
	InitContainers         []container.Container    `json:"init_containers,omitempty"`
	Containers             []container.Container    `json:"containers,omitempty"`
	RestartPolicy          RestartPolicy            `json:"restart_policy,omitempty"`
	TerminationGracePeriod *int64                   `json:"termination_grace_period,omitempty"`
	ActiveDeadline         *int64                   `json:"active_deadline,omitempty"`
	DNSPolicy              DNSPolicy                `json:"dns_policy,omitempty"`
	NodeSelector           map[string]string        `json:"nodeSelector,omitempty"`
	Account                string                   `json:"account,omitempty"`
	AutomountAccountToken  *bool                    `json:"automountAccountToken,omitempty"`
	Node                   string                   `json:"node,omitempty"`
	HostMode               []HostMode               `json:"host_mode,omitempty"`
	ShareNamespace         *bool                    `json:"shareNamespace,omitempty"`
	FSGID                  *int64                   `json:"fs_gid,omitempty"`
	GIDs                   []int64                  `json:"gids,omitempty"`
	Registries             []string                 `json:"registry_secrets,omitempty"`
	Hostname               string                   `json:"hostname,omitempty"`
	Affinity               *affinity.Affinity       `json:"affinity,omitempty"`
	SchedulerName          string                   `json:"scheduler_name,omitempty"`
	Tolerations            []toleration.Toleration  `json:"tolerations,omitempty"`
	HostAliases            []hostalias.HostAlias    `json:"host_aliases,omitempty"`
	PriorityClass          string                   `json:"priorityClass,omitempty"`
	Priority               *int32                   `json:"priority",omitempty"`
	Nameservers            []string                 `json:"nameservers,omitempty"`
	SearchDomains          []string                 `json:"searchDomains,omitempty"`
	ResolverOptions        []ResolverOptions        `json:"resolverOptions,omitempty"`
	Gates                  []PodConditionType       `json:"gates,omitempty"`
	RuntimeClass           *string                  `json:"runtimeClass,omitempty"`
	ServiceLinks           *bool                    `json:"serviceLinks,omitempty"`
}

type PodConditionType int

const (
	PodConditionScheduled PodConditionType = iota
	PodConditionReady
	PodConditionInitialized
	PodConditionReasonUnschedulable
	PodConditionContainersReady
	PodConditionNone
)

type ResolverOptions struct {
	Name  string  `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// RestartPolicy defines the pod restart policy
type RestartPolicy int

const (
	RestartPolicyDefault RestartPolicy = iota
	RestartPolicyAlways
	RestartPolicyOnFailure
	RestartPolicyNever
)

// DNSPolicy defines the pod dns policy
type DNSPolicy int

const (
	DNSClusterFirstWithHostNet DNSPolicy = iota
	DNSClusterFirst
	DNSDefault
	DNSNone
	DNSUnset
)

// HostMode defines the pod host mode
type HostMode int

const (
	HostModeNet HostMode = iota
	HostModePID
	HostModeIPC
)
