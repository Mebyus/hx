package token

import "strconv"

// Pos represents a token position in source code
type Pos struct {
	Line int
	Col  int
}

func (p *Pos) NextLine() {
	p.Line++
	p.Col = 0
}

func (p *Pos) NextCol() {
	p.Col++
}

func (p Pos) String() string {
	return strconv.FormatInt(int64(p.Line+1), 10) + ":" +
		strconv.FormatInt(int64(p.Col+1), 10)
}
