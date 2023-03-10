package scanner

func isAlphanum(b int) bool {
	return ('a' <= b && b <= 'z') || b == '_' || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func isHexadecimalDigit(b int) bool {
	return ('0' <= b && b <= '9') || ('a' <= b && b <= 'f') || ('A' <= b && b <= 'F')
}

func isBinaryDigit(b int) bool {
	return b == '0' || b == '1'
}

func isWhitespace(b int) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
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
