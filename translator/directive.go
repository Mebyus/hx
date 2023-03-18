package translator

import "codeberg.org/mebyus/hx/token"

type Directive interface {
	Apply(*Translator) error
}

func ParseDirective(tok token.Token) (Directive, error) {
	return &StringConstDefinition{}, nil
}

type StringConstDefinition struct {
	Val ConstValue
}

func (c *StringConstDefinition) Apply(t *Translator) error {
	return nil
}

type ConstValue struct {
}
