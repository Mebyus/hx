package reverse

import (
	"bytes"
	"os"
)

const bytesInOneLine = 16

func ReverseFile(filename string) (text []byte, err error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	buf := bytes.Buffer{}
	for i, b := range src {
		hex := byteAsHex[b]
		buf.WriteByte(hex[0])
		buf.WriteByte(hex[1])

		var whitespace byte
		if (i+1)%bytesInOneLine == 0 {
			whitespace = '\n'
		} else {
			whitespace = ' '
		}
		buf.WriteByte(whitespace)
	}

	return buf.Bytes(), nil
}
