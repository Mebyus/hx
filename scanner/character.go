package scanner

func isAlphanum(c int) bool {
	return isLetterOrUnderscore(c) || isDecimalDigit(c)
}

func isHexadecimalDigit(c int) bool {
	return isDecimalDigit(c) || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

func isBinaryDigit(c int) bool {
	return c == '0' || c == '1'
}

func isWhitespace(c int) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func isLetterOrUnderscore(c int) bool {
	return isLetter(c) || c == '_'
}

func isLetter(c int) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDecimalDigit(c int) bool {
	return '0' <= c && c <= '9'
}

func binaryDigitsToByte(s string) byte {
	var b byte
	for i := 0; i < 8; i++ {
		b <<= 1
		b += s[i] - '0'
	}
	return b
}

func hexDigitToVal(d int) uint8 {
	if d <= '9' {
		return uint8(d) - '0'
	}
	if d <= 'F' {
		return uint8(d) - 'A' + 0x0A
	}
	return uint8(d) - 'a' + 0x0A
}

func hexDigitsToByteValue(c, d int) byte {
	v1 := hexDigitToVal(c)
	v2 := hexDigitToVal(d)
	return v1<<4 + v2
}

func stringFromByte(b byte) string {
	return string([]byte{byte(b)})
}
