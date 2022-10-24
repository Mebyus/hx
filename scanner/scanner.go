package scanner

import (
	"io"
	"os"

	"codeberg.org/mebyus/hx/token"
)

type Scanner struct {
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

	// Scanner position inside source text
	pos token.Pos
}

func FromBytes(b []byte) (s *Scanner) {
	s = &Scanner{src: b}

	// init Scanner's current and next runes
	for i := 0; i < prefetch; i++ {
		s.advance()
	}
	s.pos = token.Pos{}
	return s
}

func FromFile(filename string) (s *Scanner, err error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	return FromBytes(src), nil
}

func FromReader(r io.Reader) (s *Scanner, err error) {
	src, err := io.ReadAll(r)
	if err != nil {
		return
	}
	return FromBytes(src), nil
}
