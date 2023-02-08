package token

import "fmt"

type Token struct {
	Pos Pos

	// Not empty only for tokens which can have
	// arbitrary literal
	//
	// Examples: numbers or strings
	Lit string

	Kind Kind
}

func (tok *Token) String() string {
	if tok.Kind.HasStaticLiteral() {
		return fmt.Sprintf("%-12s%s", tok.Pos.String(), Literal[tok.Kind])
	}
	return fmt.Sprintf("%-12s%-12s%s", tok.Pos.String(), Literal[tok.Kind], tok.Lit)
}
