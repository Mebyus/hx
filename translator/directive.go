package translator

import (
	"fmt"
	"strings"

	"codeberg.org/mebyus/hx/token"
)

type Directive interface {
	Apply(*Translator) error
}

func ParseDirective(tok token.Token) (Directive, error) {
	i := strings.Index(tok.Lit, ":")
	j := strings.Index(tok.Lit, "\"")
	if i < 3 || j < 5 {
		return nil, fmt.Errorf("bad directive format")
	}
	name := tok.Lit[2:i]
	val := tok.Lit[j+1 : len(tok.Lit)-1]

	return &StringConstDefinition{
		Name: name,
		Val:  val,
	}, nil
}

type StringConstDefinition struct {
	Name string
	Val  string
}

func (c *StringConstDefinition) Apply(t *Translator) error {
	t.cvs[c.Name] = c.Val
	return nil
}
