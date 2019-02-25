package resources

type CPU struct {
	Min string `json:"min,omitempty"`
	Max string `json:"max,omitempty"`
}

func (c *CPU) IsEmpty() bool {
	return len(c.Min) == 0 && len(c.Max) == 0
}

type Mem struct {
	Min string `json:"min,omitempty"`
	Max string `json:"max,omitempty"`
}

func (m *Mem) IsEmpty() bool {
	return len(m.Min) == 0 && len(m.Max) == 0
}
