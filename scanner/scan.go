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

	if s.c == '#' {
		return s.scanDirective()
	}

	if s.c == '$' {
		return s.scanIdentifier()
	}

	if s.c == '@' {
		return s.scanReference()
	}

	return s.scanOther()
}

func (s *Scanner) createToken(kind token.Kind) token.Token {
	return token.Token{
		Kind: kind,
		Pos:  s.pos,
	}
}

func (s *Scanner) scanWordWithPrefix(kind token.Kind) (tok token.Token) {
	tok.Pos = s.pos
	s.store()

	if !isAlphanum(s.next) {
		tok.Kind = token.Illegal
		tok.Lit = s.collect()
		return
	}

	s.storeWord()
	tok.Kind = kind
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanIdentifier() (tok token.Token) {
	return s.scanWordWithPrefix(token.Identifier)
}

func (s *Scanner) scanReference() (tok token.Token) {
	return s.scanWordWithPrefix(token.Reference)
}

func (s *Scanner) scanDirective() (tok token.Token) {
	tok.Kind = token.Directive
	tok.Pos = s.pos

	for s.c != eof && s.c != '\n' {
		s.store()
	}

	tok.Lit = s.collect()
	return
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

func (s *Scanner) scanHexByte() (tok token.Token) {
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

	tok.Lit = s.collect()
	if len(tok.Lit) != 2 {
		tok.Kind = token.Illegal
		return
	}

	tok.Kind = token.HexByte
	tok.Val = uint64(hexDigitsToByteValue(int(tok.Lit[0]), int(tok.Lit[1])))
	return
}

func (s *Scanner) scanBinaryByteAtPos(pos token.Pos) (tok token.Token) {
	tok.Pos = pos
	for s.c != eof && isBinaryDigit(s.c) {
		s.store()
	}

	if isAlphanum(s.c) {
		s.storeWord()
		tok.Kind = token.Illegal
		tok.Lit = s.collect()
		return
	}

	tok.Lit = s.collect()
	if len(tok.Lit) != 8 {
		tok.Kind = token.Illegal
		return
	}

	tok.Kind = token.BinaryByte
	tok.Val = uint64(binaryDigitsToByte(tok.Lit))
	return
}

func (s *Scanner) scanNumber() (tok token.Token) {
	if !isBinaryDigit(s.c) {
		tok = s.scanHexByte()
		return
	}
	if !isBinaryDigit(s.next) {
		tok = s.scanHexByte()
		return
	}

	pos := s.pos
	prev := s.c
	s.store()
	if isBinaryDigit(s.next) {
		tok = s.scanBinaryByteAtPos(pos)
		return
	}
	if isWhitespace(s.next) || s.next == eof {
		tok.Val = uint64(hexDigitsToByteValue(prev, s.c))
		s.store()
		tok.Kind = token.HexByte
		tok.Lit = s.collect()
		tok.Pos = pos
		return
	}
	return s.scanIllegalWordAtPos(pos)
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

func (s *Scanner) scanIllegalWord() (tok token.Token) {
	tok = s.createToken(token.Illegal)
	s.storeWord()
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanIllegalWordAtPos(pos token.Pos) (tok token.Token) {
	tok = s.createToken(token.Illegal)
	tok.Pos = pos
	s.storeWord()
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanOneByteToken(kind token.Kind) token.Token {
	tok := s.createToken(kind)
	s.advance()
	return tok
}

func (s *Scanner) scanOther() token.Token {
	switch s.c {
	case ':':
		return s.scanOneByteToken(token.Colon)
	case '-':
		return s.scanOneByteToken(token.Minus)
	case '{':
		return s.scanOneByteToken(token.LeftBrace)
	case '}':
		return s.scanOneByteToken(token.RightBrace)
	default:
		return s.scanIllegalByteToken()
	}
}
