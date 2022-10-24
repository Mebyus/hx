package token

type Kind int

const (
	EOF Kind = iota

	noStaticLiteral

	HexInteger  // A7
	String      // "string literal"
	LineComment // // it's a line comment
	Illegal
)

func (kind Kind) String() string {
	return Literal[kind]
}

func (kind Kind) HasStaticLiteral() bool {
	return kind < noStaticLiteral
}
