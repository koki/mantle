package affinity

type Affinity struct {
	NodeAffinity    map[AffinityType][]NodeTerm `json:"node,omitempty"`
	PodAffinity     map[AffinityType][]PodTerm  `json:"pod,omitempty"`
	PodAntiAffinity map[AffinityType][]PodTerm  `json:"anti_pod,omitempty"`
}

type PodTerm struct {
	Weight     int32    `json:"weight,omitempty"`
	Selector   Selector `json:"selector,omitempty"`
	Topology   string   `json:"topology,omitempty"`
	Namespaces []string `json:"namespaces,omitempty"`
}

type NodeTerm struct {
	Weight      int32             `json:"weight,omitempty"`
	Expressions []NodeExpressions `json:"expressions,omitempty"`
	Fields      []NodeExpressions `json:"fields,omitempty"`
}

type AffinityType int

const (
	AffinityHard AffinityType = iota
	AffinitySoft
)

type Selector struct {
	Labels      map[string]string     `json:"labels,omitempty"`
	Expressions []SelectorExpressions `json:"expressions,omitempty"`
}

type SelectorExpressions struct {
	Key    string           `json:"key,omitempty"`
	Op     SelectorOperator `json:"op,omitempty"`
	Values []string         `json:"values,omitempty"`
}

type SelectorOperator int

const (
	SelectorOperatorIn SelectorOperator = iota
	SelectorOperatorNotIn
	SelectorOperatorExists
	SelectorOperatorDoesNotExist
)

type NodeExpressions struct {
	Key    string       `json:"key,omitempty"`
	Op     NodeOperator `json:"op,omitempty"`
	Values []string     `json:"values,omitempty"`
}

type NodeOperator int

const (
	NodeOperatorIn NodeOperator = iota
	NodeOperatorNotIn
	NodeOperatorExists
	NodeOperatorDoesNotExist
	NodeOperatorGt
	NodeOperatorLt
)
