package shell

import (
	"errors"
	"strings"
)

type Aliases struct {
	table map[string][]string
}

var (
	ErrNeedValue = errors.New("An alias needs a value")
	ErrNeedLabel = errors.New("An alias needs a label")
)

func NewAliases() *Aliases {
	return &Aliases{
		table: make(map[string][]string),
	}
}

func (a *Aliases) parse(s string) error {
	s = strings.TrimSpace(s)

	switch i := strings.Index(s, "="); i {
	case -1, len(s) - 1:
		return ErrNeedValue
	case 0:
		return ErrNeedLabel
	default:
		label, value := strings.TrimSpace(s[:i]), strings.TrimSpace(s[i+1:])
		a.table[label] = strings.Fields(value)
	}

	return nil
}

func (a *Aliases) has(s string) (exists bool) {
	_, exists = a.table[s]
	return
}

func (a *Aliases) get(s string) []string {
	return a.table[s]
}
