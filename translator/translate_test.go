package translator

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"codeberg.org/mebyus/hx/lexer"
	"codeberg.org/mebyus/hx/token"
)

const testDataDir = "testdata"

const testResultLiteral = "TEST_RESULT"

func TestTranslate(t *testing.T) {
	names, err := discoverTestFiles(testDataDir)
	if err != nil {
		t.Fatalf("load test data: %v", err)
	}
	for _, name := range names {
		path := filepath.Join(testDataDir, name)
		lx, err := lexer.FromFile(path)
		if err != nil {
			t.Errorf("[ %s ] read test file: %v", name, err)
			continue
		}
		tokens, reachedEOF := scanUntilTestResultLiteral(lx)
		if reachedEOF {
			t.Errorf("[ %s ] test file doesn't have result section", name)
			continue
		}
		s := lexer.FromTokens(tokens)
		tr := FromStream(s)
		gotCode, err := tr.Translate()
		if err != nil {
			t.Errorf("[ %s ] failed to translate test data: %v", name, err)
			continue
		}
		wantCode, err := translateHexByteStream(lx)
		if err != nil {
			t.Errorf("[ %s ] failed to translate test result data: %v", name, err)
			continue
		}
		err = compareTranslatedCode(gotCode, wantCode)
		if err != nil {
			t.Errorf("[ %s ] got  % X", name, gotCode)
			t.Errorf("[ %s ] want % X", name, wantCode)
			t.Errorf("[ %s ] test results are not equal: %v", name, err)
			continue
		}
	}
}

func compareTranslatedCode(got, want []byte) error {
	if len(got) != len(want) {
		return fmt.Errorf("lengths are different, got=%d, want=%d", len(got), len(want))
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return fmt.Errorf("bytes at index=%d are different, got=%02X, want=%02X", i, got[i], want[i])
		}
	}
	return nil
}

func scanUntilTestResultLiteral(lx lexer.Stream) (tokens []token.Token, reachedEOF bool) {
	for {
		tok := lx.Lex()
		switch tok.Kind {
		case token.LineComment:
			lit := strings.TrimSpace(strings.TrimPrefix(tok.Lit, "//"))
			if lit == testResultLiteral {
				return
			}
			tokens = append(tokens, tok)
		case token.EOF:
			reachedEOF = true
			return
		default:
			tokens = append(tokens, tok)
		}
	}
}

func discoverTestFiles(dir string) (names []string, err error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.Type().IsRegular() && strings.HasSuffix(name, ".hx") {
			names = append(names, name)
		}
	}
	return
}

func translateHexByteStream(s lexer.Stream) (code []byte, err error) {
	for {
		tok := s.Lex()
		switch tok.Kind {
		case token.EOF:
			return code, nil
		case token.LineComment:
			// skip comment
		case token.HexByte:
			if tok.Val > 255 {
				return nil, fmt.Errorf("translate token ( %s ): %v", tok.Compact(), ErrOutOfByteRange)
			}
			code = append(code, byte(tok.Val))
		default:
			return nil, fmt.Errorf("unexpected token ( %s )", tok.Compact())
		}
	}
}

func parseHexByte(lit string) (b byte, err error) {
	if len(lit) != 2 {
		return 0, ErrBadByteFormat
	}
	v, err := strconv.ParseUint(lit, 16, 64)
	if err != nil {
		return 0, ErrBadByteFormat
	}
	if v >= 1<<8 {
		return 0, ErrOutOfByteRange
	}
	return byte(v), nil
}
