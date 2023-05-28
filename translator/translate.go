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
		case token.Directive:
			dir, err := ParseDirective(t.tok)
			if err != nil {
				return nil, fmt.Errorf("parse directive [ %s ]: %v", t.tok.Compact(), err)
			}
			err = dir.Apply(t)
			if err != nil {
				return nil, fmt.Errorf("apply directive [ %s ]: %v", t.tok.Compact(), err)
			}
		case token.Identifier:
			name := t.tok.Lit[2:]
			val, ok := t.cvs[name]
			if !ok {
				return nil, fmt.Errorf("identifier '%s' not defined", name)
			}
			switch v := val.(type) {
			case string:
				t.code = append(t.code, []byte(v)...)
			case []byte:
				t.code = append(t.code, v...)
			default:
				panic("unexpected const value type: " + fmt.Sprintf("%v (%t)", v, v))
			}
		default:
			panic("unknown token: " + t.tok.Compact())
		}

		if err != nil {
			return nil, fmt.Errorf("translate token [ %s ]: %v", t.tok.Compact(), err)
		}
	}
}

func (t *Translator) translateByte() {
	t.code = append(t.code, byte(t.tok.Val))
}

func (t *Translator) translateString() {
	lit := t.tok.Lit
	if len(lit) < 2 {
		panic("malformed string token: " + t.tok.Compact())
	}
	str := lit[1 : len(lit)-1]
	t.code = append(t.code, []byte(str)...)
}
