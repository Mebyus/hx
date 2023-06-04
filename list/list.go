package list

import (
	"io"

	"codeberg.org/mebyus/hx/lexer"
	"codeberg.org/mebyus/hx/token"
)

func List(w io.Writer, s lexer.Stream) error {
	for {
		tok := s.Lex()
		err := put(w, tok)
		if err != nil {
			return err
		}
		err = newline(w)
		if err != nil {
			return err
		}
		if tok.IsEOF() {
			return nil
		}
	}
}

func put(w io.Writer, tok token.Token) error {
	_, err := w.Write([]byte(tok.String()))
	return err
}

func newline(w io.Writer) error {
	_, err := w.Write([]byte("\n"))
	return err
}
