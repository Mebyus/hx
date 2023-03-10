package translator

import (
	"fmt"

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
		case token.HexByte, token.BinaryByte:
			t.translateByte()
		case token.String:
			t.translateString()
		default:
			panic("unknown token: " + t.tok.String())
		}

		if err != nil {
			return nil, fmt.Errorf("translate token [ %s ]: %v", t.tok.Compact(), err)
		}
	}
}

func (t *Translator) translateByte() {
	t.code = append(t.code, byte(t.tok.Val))
	return
}

func (t *Translator) translateString() {
	lit := t.tok.Lit
	if len(lit) < 2 {
		panic("malformed string token: " + t.tok.String())
	}
	str := lit[1 : len(lit)-1]
	t.code = append(t.code, []byte(str)...)
	return
}
