package main

import (
	"fmt"
	"os"
	"strings"

	"codeberg.org/mebyus/hx/translator"
)

const (
	ext    = ".hx"
	outExt = ".bin"
)

func fatal(v interface{}) {
	fmt.Println("fatal:", v)
	os.Exit(1)
}

func fatalf(format string, v ...interface{}) {
	fmt.Printf("fatal:"+format+"\n", v...)
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

	code, err := translator.TranslateFile(filename)
	if err != nil {
		fatal(err)
	}

	outputFilename := transformFilename(filename)
	err = os.WriteFile(outputFilename, code, 0664)
	if err != nil {
		fatal(err)
	}
}
