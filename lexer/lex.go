package lexer

import (
	"codeberg.org/mebyus/hx/token"
)

func (s *Lexer) Lex() token.Token {
	if s.c == eof {
		return s.createToken(token.EOF)
	}

	s.skipWhitespace()
	if s.c == eof {
		return s.createToken(token.EOF)
	}

	if isHexadecimalDigit(s.c) {
		return s.lexNumber()
	}

	if s.c == '"' {
		return s.lexString()
	}

	if s.c == '<' {
		return s.lexLabel()
	}

	if s.c == '/' && s.next == '/' {
		return s.lexLineComment()
	}

	if s.c == '#' {
		return s.lexDirective()
	}

	if s.c == '$' && s.next == '.' {
		return s.lexPlacement()
	}

	if s.c == '@' && s.next == '.' {
		return s.lexReference()
	}

	return s.lexOther()
}

func (s *Lexer) createToken(kind token.Kind) token.Token {
	return token.Token{
		Kind: kind,
		Pos:  s.pos,
	}
}

func (s *Lexer) lexWordWithPrefix(kind token.Kind) (tok token.Token) {
	tok.Pos = s.pos
	s.store()
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

func (s *Lexer) lexPlacement() (tok token.Token) {
	return s.lexWordWithPrefix(token.Placement)
}

func (s *Lexer) lexReference() (tok token.Token) {
	return s.lexWordWithPrefix(token.Reference)
}

func (s *Lexer) lexDirective() (tok token.Token) {
	tok.Kind = token.Directive
	tok.Pos = s.pos

	for s.c != eof && s.c != '\n' {
		s.store()
	}

	tok.Lit = s.collect()
	return
}

func (s *Lexer) lexString() (tok token.Token) {
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

func (s *Lexer) lexHexByte() (tok token.Token) {
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

	lit := s.collect()
	if len(lit) != 2 {
		tok.Lit = lit
		tok.Kind = token.Illegal
		return
	}

	tok.Kind = token.HexByte
	tok.Val = uint64(hexByteToValue(lit))
	return
}

func (s *Lexer) lexBinaryByteAtPos(pos token.Pos) (tok token.Token) {
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

func (s *Lexer) lexLabel() (tok token.Token) {
	tok.Pos = s.pos
	s.store() // consume <
	if !isLetterOrUnderscore(s.c) {
		s.storeWord()
		if s.c == '>' {
			s.store()
		}
		tok.Lit = s.collect()
		tok.Kind = token.Illegal
		return
	}

	s.storeWord()
	if s.c == '>' {
		s.store()
		tok.Lit = s.collect()
		if len(tok.Lit) < 3 {
			tok.Kind = token.Illegal
			return
		}
		tok.Kind = token.Label
		return
	}

	tok.Kind = token.Label
	tok.Lit = s.collect()
	return
}

func (s *Lexer) lexNumber() (tok token.Token) {
	if !isBinaryDigit(s.c) {
		tok = s.lexHexByte()
		return
	}
	if !isBinaryDigit(s.next) {
		tok = s.lexHexByte()
		return
	}

	pos := s.pos
	prev := s.c
	s.store()
	if isBinaryDigit(s.next) {
		tok = s.lexBinaryByteAtPos(pos)
		return
	}
	if isWhitespace(s.next) || s.next == eof {
		tok.Val = uint64(hexDigitsToByteValue(byte(prev), byte(s.c)))
		s.store()
		tok.Kind = token.HexByte
		_ = s.collect()
		tok.Pos = pos
		return
	}
	return s.lexIllegalWordAtPos(pos)
}

func (s *Lexer) lexLineComment() (tok token.Token) {
	tok.Kind = token.LineComment
	tok.Pos = s.pos

	for s.c != eof && s.c != '\n' {
		s.store()
	}

	tok.Lit = s.collect()
	return
}

func (s *Lexer) lexIllegalByteToken() token.Token {
	tok := s.createToken(token.Illegal)
	tok.Lit = stringFromByte(byte(s.c))
	s.advance()
	return tok
}

func (s *Lexer) lexIllegalWordAtPos(pos token.Pos) (tok token.Token) {
	tok = s.createToken(token.Illegal)
	tok.Pos = pos
	s.storeWord()
	tok.Lit = s.collect()
	return
}

func (s *Lexer) lexOneByteToken(kind token.Kind) token.Token {
	tok := s.createToken(kind)
	s.advance()
	return tok
}

func (s *Lexer) lexOther() token.Token {
	switch s.c {
	case ':':
		return s.lexOneByteToken(token.Colon)
	case '-':
		return s.lexOneByteToken(token.Minus)
	case '{':
		return s.lexOneByteToken(token.LeftBrace)
	case '}':
		return s.lexOneByteToken(token.RightBrace)
	default:
		return s.lexIllegalByteToken()
	}
}
