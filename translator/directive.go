package translator

import (
	"fmt"
	"strconv"
	"strings"

	"codeberg.org/mebyus/hx/token"
)

type Directive interface {
	Apply(*Translator) error
}

func ParseDirective(tok token.Token) (Directive, error) {
	i := strings.Index(tok.Lit, ":")
	j := strings.Index(tok.Lit, "=")
	if i < 3 || j < 6 {
		return nil, fmt.Errorf("bad directive format")
	}
	name := strings.TrimSpace(tok.Lit[2:i])
	constType := strings.TrimSpace(tok.Lit[i+1 : j])
	switch constType {
	case "str":
		val, err := parseStringConstValue(tok.Lit[j+1:])
		if err != nil {
			return nil, err
		}
		return &StringConstDefinition{
			Name: name,
			Val:  val,
		}, nil
	case "seq":
		val, err := parseSequenceConstValue(tok.Lit[j+1:])
		if err != nil {
			return nil, err
		}
		return &SequenceConstDefinition{
			Name: name,
			Val:  val,
		}, nil
	default:
		return nil, fmt.Errorf("unknown type: %s", constType)
	}
}

func parseStringConstValue(str string) (string, error) {
	i := strings.Index(str, "\"")
	if i < 0 {
		return "", fmt.Errorf("bad directive format")
	}
	return str[i+1 : len(str)-1], nil
}

func parseSequenceConstValue(str string) ([]byte, error) {
	fields := strings.Fields(str)
	if len(fields) == 0 {
		return nil, fmt.Errorf("empty sequence")
	}
	var bb []byte
	for _, f := range fields {
		if len(f) != 2 {
			return nil, fmt.Errorf("bad directive format")
		}
		b, err := strconv.ParseUint(f, 16, 8)
		if err != nil {
			return nil, fmt.Errorf("bad directive format")
		}
		bb = append(bb, byte(b))
	}
	return bb, nil
}

type StringConstDefinition struct {
	Name string
	Val  string
}

func (c *StringConstDefinition) Apply(t *Translator) error {
	t.cvs[c.Name] = c.Val
	return nil
}

type SequenceConstDefinition struct {
	Name string
	Val  []byte
}

func (c *SequenceConstDefinition) Apply(t *Translator) error {
	t.cvs[c.Name] = c.Val
	return nil
}
