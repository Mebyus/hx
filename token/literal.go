package token

var Literal = [...]string{
	// Non static or empty literals
	EOF: "<EOF>",

	HexInteger:  "<HEXINT>",
	String:      "<STRING>",
	LineComment: "<LINECOM>",
	Illegal:     "<ILLEGAL>",
}
