package scanner

func isAlphanum(b int) bool {
	return ('a' <= b && b <= 'z') || b == '_' || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func isHexadecimalDigit(b int) bool {
	return ('0' <= b && b <= '9') || ('a' <= b && b <= 'f') || ('A' <= b && b <= 'F')
}

func isWhitespace(b int) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func stringFromByte(b byte) string {
	return string([]byte{byte(b)})
}
