# HX - hex to binary translator

- [Examples](#examples)
- [Usage](#usage)
- [Build](#build)

**hx** is a simple command-line utility which translates files
with plain text to binary files according to small set of rules:

1. File is read line by line
2. Blank and whitespace-only lines are ignored
3. Lines which have "//" as two first non-whitespace characters are ignored (aka comments) 
4. On all other lines words (separated via whitespace) are treated as hex numbers representing a byte. If it's not a number or number is negative or bigger than 1 byte an error is thrown, at which point execution stops with no output produced
5. Each byte is placed into output file in order it appears in input file (left to right, top to bottom)

## Examples

```
// this line and blank lines below will be ignored



// next line produces one byte 0x10 to output
10

// two next lines yield 0x01 0xA3 0x34 0x00
01 A3 34
00

    // whitespace at the beginning of the line and between words is ignored,
    // so next line yields 0x20 0x01 0x02
    20     01    02

// uncomment line below to see errors
// 130 -1 asd

// "130" is bigger than 1 byte
// "-1" is negative
// "asd" is not a hex number
```

Text above will be translated to binary with the following sequence of bytes
```
10 01 A3 34 00 20 01 02
```

## Usage

Invoke with simple command
```
hx my_text_file
```

This command will produce file `my_text_file.bin` in binary format. You can open it with hex editor

## Build

**hx** is written in Go, so in order to build it from source you will
need Go installed

Build with command (assuming you are in repository's root directory)
```
go build .
```

Install it to via go's install (places executable in $GOPATH/bin)
```
go install .
```

There are also convenient make targets for build
```
make
```

And install
```
make install
```
