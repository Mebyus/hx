package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ext    = ".hx"
	outExt = ".bin"
)

func fatal(v interface{}) {
	fmt.Println(v)
	os.Exit(1)
}

func fatalf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
	os.Exit(1)
}

func transformFilename(name string) string {
	return strings.TrimSuffix(name, ext) + outExt
}

func main() {
	if len(os.Args) < 2 {
		fatal("specify file with hex codes to translate")
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fatal(err)
	}

	var code []byte
	line := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line++
		str := strings.TrimSpace(scanner.Text())
		if str == "" || strings.HasPrefix(str, "//") {
			continue
		}
		split := strings.Fields(str)
		if len(split) == 0 {
			continue
		}
		for _, s := range split {
			number, err := strconv.ParseUint(s, 16, 64)
			if err != nil {
				fatalf("bad hex number %s at line %d", s, line)
			}
			if number > 255 {
				fatalf("number %s at line %d is bigger than 1 byte", s, line)
			}
			code = append(code, byte(number))
		}
	}
	err = scanner.Err()
	if err != nil {
		fatal(err)
	}

	outputFilename := transformFilename(filename)
	err = os.WriteFile(outputFilename, code, 0664)
	if err != nil {
		fatal(err)
	}
}
