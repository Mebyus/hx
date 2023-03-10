package token

var Literal = [...]string{
	// Non static or empty literals
	EOF: "<EOF>",

	Colon:      ":",
	Minus:      "-",
	LeftBrace:  "{",
	RightBrace: "}",

	HexByte:     "<HEXBYTE>",
	BinaryByte:  "<BINBYTE>",
	String:      "<STRING>",
	Identifier:  "<IDENT>",
	Reference:   "<REFNCE>",
	Directive:   "<DIRTVE>",
	LineComment: "<LINECOM>",
	Illegal:     "<ILLEGAL>",
}
