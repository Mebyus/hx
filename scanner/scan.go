package scanner

import (
	"codeberg.org/mebyus/hx/token"
)

func (s *Scanner) Scan() token.Token {
	if s.c == eof {
		return s.createToken(token.EOF)
	}

	s.skipWhitespace()
	if s.c == eof {
		return s.createToken(token.EOF)
	}

	if isHexadecimalDigit(s.c) {
		return s.scanNumber()
	}

	if s.c == '"' {
		return s.scanString()
	}

	if s.c == '/' && s.next == '/' {
		return s.scanLineComment()
	}

	return s.scanOther()
}

func (s *Scanner) createToken(kind token.Kind) token.Token {
	return token.Token{
		Kind: kind,
		Pos:  s.pos,
	}
}

func (s *Scanner) scanString() (tok token.Token) {
	tok.Pos = s.pos
	s.store() // consume "

	for s.c != eof && s.c != '"' {
		s.store()
	}

	if s.c == eof {
		tok.Kind = token.Illegal
	} else {
		tok.Kind = token.String
		s.store() // consume "
	}
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanHexadecimalNumber() (tok token.Token) {
	tok.Pos = s.pos

	for s.c != eof && isHexadecimalDigit(s.c) {
		s.store()
	}

	if isAlphanum(s.c) {
		s.storeWord()
		tok.Kind = token.Illegal
		tok.Lit = s.collect()
		return
	}

	tok.Kind = token.HexInteger
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanNumber() (tok token.Token) {
	tok = s.scanHexadecimalNumber()
	return
}

func (s *Scanner) scanLineComment() (tok token.Token) {
	tok.Kind = token.LineComment
	tok.Pos = s.pos

	for s.c != eof && s.c != '\n' {
		s.store()
	}

	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanIllegalByteToken() token.Token {
	tok := s.createToken(token.Illegal)
	tok.Lit = stringFromByte(byte(s.c))
	s.advance()
	return tok
}

func (s *Scanner) scanOther() token.Token {
	return s.scanIllegalByteToken()
}
