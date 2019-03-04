package action

type ActionType int

const (
	ActionTypeCommand ActionType = iota
	ActionTypeHTTP
	ActionTypeHTTPS
	ActionTypeTCP
)

type Action struct {
	ActionType ActionType `json:"actionType,omitempty"`
	Command    []string   `json:"command,omitempty"`
	Headers    []string   `json:"headers,omitempty"`
	Host       string     `json:"host,omitempty"`
	Port       string     `json:"host,omitempty"`
	Path       string     `json:"path,omitempty"`
}
