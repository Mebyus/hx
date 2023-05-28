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
	Label:       "<LABEL>",
	Placement:   "<PLACENT>",
	Reference:   "<REFRNCE>",
	Directive:   "<DIRCTVE>",
	LineComment: "<LINECOM>",
	Illegal:     "<ILLEGAL>",
}
