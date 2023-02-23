package translator

import (
	"fmt"
	"strconv"

	"codeberg.org/mebyus/hx/token"
)

var (
	ErrBadByteFormat  = fmt.Errorf("bad byte format")
	ErrOutOfByteRange = fmt.Errorf("out of byte range")
)

func TranslateFile(filename string) (code []byte, err error) {
	t, err := FromFile(filename)
	if err != nil {
		return
	}
	return t.Translate()
}

func (t *Translator) Translate() (code []byte, err error) {
	for {
		t.tok = t.sc.Scan()
		switch t.tok.Kind {
		case token.EOF:
			return t.code, nil
		case token.Illegal:
			return nil, fmt.Errorf("illegal token " + t.tok.Compact())
		case token.LineComment:
			// skip comment
		case token.HexByte:
			err = t.translateHexInteger()
		case token.String:
			err = t.translateString()
		default:
			panic("unknown token: " + t.tok.String())
		}

		if err != nil {
			return nil, fmt.Errorf("translate token [ %s ]: %v", t.tok.Compact(), err)
		}
	}
}

func (t *Translator) translateHexInteger() (err error) {
	lit := t.tok.Lit
	if len(lit) != 2 {
		return ErrBadByteFormat
	}
	v, err := strconv.ParseUint(lit, 16, 64)
	if err != nil {
		return ErrBadByteFormat
	}
	if v >= 1<<8 {
		return ErrOutOfByteRange
	}
	t.code = append(t.code, byte(v))
	return
}

func (t *Translator) translateString() (err error) {
	lit := t.tok.Lit
	if len(lit) < 2 {
		panic("malformed string token: " + t.tok.String())
	}
	str := lit[1 : len(lit)-1]
	t.code = append(t.code, []byte(str)...)
	return
}
