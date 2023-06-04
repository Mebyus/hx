package lexer

import (
	"io"
	"os"

	"codeberg.org/mebyus/hx/token"
)

type Stream interface {
	Lex() token.Token
}

type Lexer struct {
	// charcode at current Scanner position
	c int

	// next charcode
	next int

	// src reading index
	i int

	// literal buffer
	buf []byte

	// source text which is scanned by the Scanner
	src []byte

	// Lexer position inside source text
	pos token.Pos
}

func FromBytes(b []byte) (s *Lexer) {
	s = &Lexer{src: b}

	// init Scanner's current and next runes
	for i := 0; i < prefetch; i++ {
		s.advance()
	}

	// reset scanner position
	s.pos = token.Pos{}
	return s
}

func FromFile(filename string) (s *Lexer, err error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	return FromBytes(src), nil
}

func FromReader(r io.Reader) (s *Lexer, err error) {
	src, err := io.ReadAll(r)
	if err != nil {
		return
	}
	return FromBytes(src), nil
}
