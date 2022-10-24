package translator

import (
	"io"
	"os"

	"codeberg.org/mebyus/hx/scanner"
	"codeberg.org/mebyus/hx/token"
)

type Translator struct {
	code []byte

	tok token.Token
	sc  *scanner.Scanner
}

func FromBytes(b []byte) (t *Translator) {
	t = &Translator{sc: scanner.FromBytes(b)}
	return
}

func FromFile(filename string) (t *Translator, err error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	return FromBytes(src), nil
}

func FromReader(r io.Reader) (t *Translator, err error) {
	src, err := io.ReadAll(r)
	if err != nil {
		return
	}
	return FromBytes(src), nil
}
