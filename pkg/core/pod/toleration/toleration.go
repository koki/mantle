package toleration

type Toleration struct {
	Key               string             `json:"key,omitempty"`
	Op                TolerationOperator `json:"op,omitempty"`
	Value             string             `json:"value,omitempty"`
	Effect            TaintEffect        `json:"effect,omitempty"`
	ExpirationSeconds *int64             `json:"expirationSeconds,omitempty"`
}

type TolerationOperator int

const (
	TolerationOperatorExists TolerationOperator = iota
	TolerationOperatorEqual
)

type TaintEffect int

const (
	TaintEffectNoSchedule TaintEffect = iota
	TaintEffectPreferNoSchedule
	TaintEffectNoExecute
)
