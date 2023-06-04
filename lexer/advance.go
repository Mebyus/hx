package lexer

const (
	eof      = -1
	prefetch = 2
	nonASCII = 1 << 7
)

func (s *Lexer) advance() {
	if s.c != eof {
		if s.c == '\n' {
			s.pos.NextLine()
		} else if s.c < nonASCII {
			s.pos.NextCol()
		}
	}
	s.c = s.next
	if s.i < len(s.src) {
		s.next = int(s.src[s.i])
		s.i++
		s.pos.Offset = s.i
	} else {
		s.next = eof
	}
}

func (s *Lexer) collect() string {
	str := string(s.buf)

	// reset slice length, but keep capacity
	// to avoid new allocs
	s.buf = s.buf[:0]

	return str
}

func (s *Lexer) store() {
	s.addbuf(byte(s.c))
	s.advance()
}

func (s *Lexer) addbuf(b byte) {
	s.buf = append(s.buf, b)
}

func (s *Lexer) skipWhitespace() {
	for isWhitespace(s.c) {
		s.advance()
	}
}

func (s *Lexer) storeWord() {
	for isAlphanum(s.c) {
		s.store()
	}
}
