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

	isIllegal := false
	for s.c != eof && s.c != '"' {
		if s.c != '\\' {
			s.store()
			continue
		}

		// handle escape sequence
		switch s.next {
		case 'n':
			s.addbuf('\n')
		case 'r':
			s.addbuf('\r')
		case 't':
			s.addbuf('\t')
		case '"':
			s.addbuf('"')
		case '\\':
			s.addbuf('\\')
		case eof:
			isIllegal = true
			s.store()
			continue
		default:
			isIllegal = true
			s.store()
			s.store()
			continue
		}
		s.advance()
		s.advance()
	}

	if s.c == eof || isIllegal {
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

	tok.Kind = token.HexByte
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
