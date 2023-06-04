package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"codeberg.org/mebyus/hx/lexer"
	"codeberg.org/mebyus/hx/list"
	"codeberg.org/mebyus/hx/reverse"
	"codeberg.org/mebyus/hx/translator"
)

const (
	ext    = ".hx"
	outExt = ".bin"
)

func fatal(v interface{}) {
	fmt.Fprintln(os.Stderr, "fatal:", v)
	os.Exit(1)
}

func fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "fatal: "+format+"\n", v...)
	os.Exit(1)
}

func transformFilename(name string) string {
	return strings.TrimSuffix(name, ext) + outExt
}

func main() {
	var isReverseCommand bool
	var isListCommand bool
	flag.BoolVar(&isReverseCommand, "r", false, "transform binary file to text file in hx format")
	flag.BoolVar(&isListCommand, "l", false, "list tokens in text file")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fatal("filename was not specified")
	}
	filename := args[0]

	if isReverseCommand {
		text, err := reverse.ReverseFile(filename)
		if err != nil {
			fatal(err)
		}

		outputFilename := filename + ext
		err = os.WriteFile(outputFilename, text, 0o664)
		if err != nil {
			fatal(err)
		}
		return
	}

	if isListCommand {
		s, err := lexer.FromFile(filename)
		if err != nil {
			fatal(err)
		}
		err = list.List(os.Stdout, s)
		if err != nil {
			fatal(err)
		}
		return
	}

	code, err := translator.TranslateFile(filename)
	if err != nil {
		fatal(err)
	}

	outputFilename := transformFilename(filename)
	err = os.WriteFile(outputFilename, code, 0o664)
	if err != nil {
		fatal(err)
	}
}
