package probe

import (
	"mantle/pkg/core/action"
)

// Probe defines the parameters for a probe
type Probe struct {
	action.Action   `json:",inline"`
	Delay           int32 `json:"delay,omitempty"`
	Interval        int32 `json:"interval,omitempty"`
	MinCountSuccess int32 `json:"min_count_success,omitempty"`
	MinCountFailure int32 `json:"min_count_fail,omitempty"`
	Timeout         int32 `json:"timeout,omitempty"`
}
