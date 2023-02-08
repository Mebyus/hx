package scanner

const (
	eof      = -1
	prefetch = 2
	nonASCII = 1 << 7
)

func (s *Scanner) advance() {
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
	} else {
		s.next = eof
	}
}

func (s *Scanner) collect() string {
	str := string(s.buf)

	// reset slice length, but keep capacity
	// to avoid new allocs
	s.buf = s.buf[:0]

	return str
}

func (s *Scanner) store() {
	s.buf = append(s.buf, byte(s.c))
	s.advance()
}

func (s *Scanner) skipWhitespace() {
	for isWhitespace(s.c) {
		s.advance()
	}
}

func (s *Scanner) storeWord() {
	for isAlphanum(s.c) {
		s.store()
	}
}
