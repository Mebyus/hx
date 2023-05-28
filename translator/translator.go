package translator

import (
	"io"
	"os"

	"codeberg.org/mebyus/hx/scanner"
	"codeberg.org/mebyus/hx/token"
)

type Scanner interface {
	Scan() token.Token
}

type Translator struct {
	code []byte

	// string const values mapped by name
	cvs map[string]any

	tok token.Token
	sc  Scanner
}

func FromScanner(sc Scanner) (t *Translator) {
	return &Translator{
		sc:  sc,
		cvs: make(map[string]any),
	}
}

func FromBytes(b []byte) (t *Translator) {
	t = FromScanner(scanner.FromBytes(b))
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
