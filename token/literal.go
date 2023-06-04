package token

var Literal = [...]string{
	empty: "<EMPTY>",

	// Non static or empty literals
	EOF: "<EOF>",

	Colon:      ":",
	Minus:      "-",
	LeftBrace:  "{",
	RightBrace: "}",

	HexByte:     "<HEX>",
	BinaryByte:  "<BIN>",
	String:      "<STR>",
	Label:       "<LBL>",
	Placement:   "<PLC>",
	Reference:   "<REF>",
	Directive:   "<DIR>",
	LineComment: "<COM>",
	Illegal:     "<ILLEGAL>",
}
