package token

import "fmt"

type Token struct {
	Pos Pos

	// Not empty only for tokens which can have
	// arbitrary literal
	//
	// Examples: labels or strings
	Lit string

	// Meaning depends on Token.Kind
	//
	// Kind: HexByte or BinaryByte. Decoded byte value
	Val uint64

	Kind Kind
}

func (tok *Token) String() string {
	if tok.Kind.HasStaticLiteral() {
		return fmt.Sprintf("%-12s%s", tok.Pos.String(), Literal[tok.Kind])
	}
	if tok.Kind == HexByte {
		return fmt.Sprintf("%-12s%-12s%02X", tok.Pos.String(), Literal[tok.Kind], tok.Val)
	}
	if tok.Kind == BinaryByte {
		return fmt.Sprintf("%-12s%-12s%08b", tok.Pos.String(), Literal[tok.Kind], tok.Val)
	}
	return fmt.Sprintf("%-12s%-12s%s", tok.Pos.String(), Literal[tok.Kind], tok.Lit)
}

func (tok *Token) Compact() string {
	if tok.Kind.HasStaticLiteral() {
		return fmt.Sprintf("%s  %s", tok.Pos.String(), Literal[tok.Kind])
	}
	if tok.Kind == HexByte {
		return fmt.Sprintf("%s  %s  %02X", tok.Pos.String(), Literal[tok.Kind], tok.Val)
	}
	if tok.Kind == BinaryByte {
		return fmt.Sprintf("%s  %s  %08b", tok.Pos.String(), Literal[tok.Kind], tok.Val)
	}
	return fmt.Sprintf("%s  %s  %s", tok.Pos.String(), Literal[tok.Kind], tok.Lit)
}

func (tok *Token) IsEOF() bool {
	return tok.Kind == EOF
}

func (tok *Token) ErasePos() {
	tok.Pos = Pos{}
}
