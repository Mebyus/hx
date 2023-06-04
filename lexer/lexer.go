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

// Parrot implements Stream by yielding tokens from supplied list
type Parrot struct {
	toks []token.Token
	i    int
}

func FromTokens(toks []token.Token) *Parrot {
	return &Parrot{
		toks: toks,
	}
}

func (p *Parrot) Lex() token.Token {
	if p.i >= len(p.toks) {
		tok := token.Token{Kind: token.EOF}
		if len(p.toks) == 0 {
			return tok
		}
		pos := p.toks[len(p.toks)-1].Pos
		pos.NextLine()
		tok.Pos = pos
		return tok
	}
	tok := p.toks[p.i]
	p.i++
	return tok
}
