# HX - hex to binary translator

- [Language](#language-rules-in-examples)
- [Examples](#examples)
- [Usage](#usage)
- [Build](#build)
- [Reverse](#reverse)

**hx** is a command-line utility which translates files with plain text to binary files according to small language.
In simple words it's just a sequence of hex numbers which represent bytes with some macros for convenience.
One can say that an **hx** "program" is a binary file written in text.

On the abstraction scale this tool is somewhere between binary and assembly.

> To clarify: **hx** does not understand any particular assembly or machine architecture

## Language rules (in examples)

### Minimal program

A minimal program would be an empty file and upon translation **hx** will emit an empty file

### Comments

Translation ignores comments. Thus next listing will be translated to an empty file

```
// Line comments are ignored (start with "//")
// Multiline are not supported (yet)

// This program results in an empty file
```

### Bytes

Hex number represents a single byte (number must be 2 digits exactly)

```
// Bytes must be separated with whitespace

01 A3 34
00

// Bytes in hx program can be written in lowercase
bc 0a f3

// This program results in a 4 bytes long file containing sequence:
// 01 A3 34 00 BC 0A F3
```

Binary number with exactly 8 digits can be used instead of hex

```
// Usage of binary and hex numbers does not lead to ambiguity
// because each form uses different number of digits

01111000
AF

11010001

// This program results in a 3 bytes long file containing sequence:
// 78 AF D1
```

### Strings

String inside double quotes emits a sequence of bytes which encode that string in UTF-8

```
// String is denoted by double quotes

"Hello, world!"

// Bytes and strings can be used within one program

5F 05

// Line can start with whitespace if needed

    42

// Non-ASCII characters can be used in strings

    "Шмель"

// This program results in a 26 bytes long file containing sequence:
// 48 65 6C 6C 6F 2C 20 77 6F 72 6C 64 21 5F 05 42
// D0 A8 D0 BC D0 B5 D0 BB D1 8C
```

Strings can contain newlines

```
"First line,
 second line
 and third all in one string!"
```

Strings support a small set of C-style escape sequence

```
// \n = newline

"this string ends with newline\n"

// \t = tab
// \" = " (double quote character)
// \\ = \
// \r = carriage return

"double quote character \" can be placed inside string"
```

> Strings can contain byte sequences which does not represent a valid UTF-8 text.
> In fact **hx** takes hands-off approach to strings: any sequence of bytes within opening and closing double quote
> (except escape sequences) is placed in output as it is

### Number literals

There are different integer representations in memory. By default number literals are placed in
output stream in Little Endian (LE) and produce 8 bytes regardless of literal value. Negative
literals converted to 2's complement form before translation

```
// Hex literal
0xA3FF1

// Decimal literal
0t45107

// Literal can also be placed via explicit literal syntax
{45107}

// By default explicit literal argument is a decimal number

// Explicit literal can take additional translation options
{45107, LE}
{45107, 8}
{45107, BE, 4}
```

### Labels

Construct `some_label:` marks address in output stream with label. Label itself does not produce any output.
At any point in program construct `@.some_label` can be used (referenced) to emit an address (offset in output stream) of that label.

Two special labels are available (their explicit usage will result in error):

* `start` - marks output stream start
* `end` - marks output stream end

Construct `[label_b : label_a]` can be used to emit difference in `label_b` and `label_a` offsets. In other words it emits
value equal to number of bytes in output stream between `label_a` and `label_b` in that particular order, assuming `label_b` comes after
`label_a`. Incorrect order will result in error. Such construct is called *length*.

Three special forms of length construct can be used:

* `[label:]` - difference between `label` and output stream start
* `[:label]` - difference betweem output stream end and `label`
* `[:]` - total number of bytes in output stream

> `[label:]` is not the same as `@.label` because directives could change `@.start` value

Label with a particular name can only be placed once inside a program, but can be referenced multiple times (or not used at all)

### Directives

Translator directive starts with "#" and continues until line end. Directives can do 4 different things:

Declare constants:

```
# b: byte = 3F

# s: seq = A4 03

# n: lit = 45107
# n: lit = -45107
# n: lit{8} = 45107
# n: lit{LE} = 45107
# n: lit{LE, 8} = 45107

# s: str = "Hello, world!"
```

```
// Constants must be declared before usage
# some_string: str = "Hello, world!"

// Now an identifier $.some_string will be substituted for "Hello, world!" in output stream
$.some_string

// This program results in a 13 bytes long file containing sequence:
// 48 65 6C 6C 6F 2C 20 77 6F 72 6C 64 21
```

Change translator options:

```
# opt syntax = 0.6.3
# opt start = 0x005F

# opt endianness = LE
# opt size = 8

# opt address size = 4
# opt value size = 6

# opt address endianness = BE
# opt value endianness = LE
```

Include other text file as though it was copy-pasted (happens before translation):

```
# include "path/to_other/file.hx"
```

Include other binary file in output stream:

```
# before "some_binary_file"
# insert "binary_file"
# after "other_binary_file"
```

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

```sh
hx my_text_file
```

This command will produce file `my_text_file.bin` in binary format. You can open it with hex editor

## Build

**hx** is written in Go, so in order to build it from source you will
need Go installed

Build with command (assuming you are in repository's root directory)

```sh
go build .
```

Install it via go's install command (places executable in $GOPATH/bin)

```sh
go install .
```

There are also convenient make targets for build

```sh
make
```

And install

```sh
make install
```

## Reverse

Reverse command can translate binary file to text file. The resulting text file can be used
as **hx** input to produce original binary file

```sh
hx -r binary_file
```
