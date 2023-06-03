package token

type Kind int

const (
	// Zero value. Do not use explicitly
	empty Kind = iota

	EOF

	Colon      // :
	Minus      // -
	LeftBrace  // {
	RightBrace // }

	noStaticLiteral

	HexByte     // A7 (2 digits exactly)
	BinaryByte  // 11010001 (8 digits exactly)
	String      // "string literal"
	Label       // <label>
	Placement   // $.some_identifier
	Reference   // @.some_label
	Directive   // # opt some_option
	LineComment // // it's a line comment
	Illegal
)

func (k Kind) String() string {
	return Literal[k]
}

func (k Kind) HasStaticLiteral() bool {
	return k < noStaticLiteral
}
