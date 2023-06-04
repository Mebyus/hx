package translator

import (
	"io"
	"os"

	"codeberg.org/mebyus/hx/lexer"
	"codeberg.org/mebyus/hx/token"
)

type Translator struct {
	code []byte

	// string const values mapped by name
	cvs map[string]any

	// label offsets mapped by name
	labels map[string]uint64

	// labels that are used, but not yet declared
	//
	// list of positions mapped by label name
	late map[string][]uint64

	// buffer for preparing reference value placement
	lbuf [8]byte

	tok token.Token
	s   lexer.Stream
}

func FromStream(s lexer.Stream) (t *Translator) {
	return &Translator{
		s:      s,
		cvs:    make(map[string]any),
		labels: make(map[string]uint64),
		late:   make(map[string][]uint64),
	}
}

func FromBytes(b []byte) (t *Translator) {
	t = FromStream(lexer.FromBytes(b))
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
