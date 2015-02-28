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

func (r *Repl) parseAlias(s string) error {
	s = strings.TrimSpace(s)

	switch i := strings.Index(s, "="); i {
	case -1:
		return ErrNeedValue
	case 0:
		return ErrNeedLabel
	default:
		label, value := strings.TrimSpace(s[:i]), strings.TrimSpace(s[i+1:])
		r.aliases.table[label] = strings.Fields(value)
	}

	return nil
}

func (r *Repl) hasAlias(s string) (exists bool) {
	_, exists = r.aliases.table[s]
	return
}

func (r *Repl) getAlias(s string) []string {
	return r.aliases.table[s]
}
