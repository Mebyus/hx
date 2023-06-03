package token

var Literal = [...]string{
	empty: "<EMPTY>",

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
	Placement:   "<PLC>",
	Reference:   "<REF>",
	Directive:   "<DIR>",
	LineComment: "<LINECOM>",
	Illegal:     "<ILLEGAL>",
}
