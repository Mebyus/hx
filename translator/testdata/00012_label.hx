// Label just marks place in output stream and by itself does not
// produce bytes
<used_label>

// By default label usage produces 8 bytes offset (at label declaration position)
// in little endian
A9 0B 1C @.used_label

<unused_label>

// Label can be used before its declaration
11 03 @.other_label

<other_label>

// TEST_RESULT

A9 0B 1C 00 00 00 00 00 00 00 00 11 03 15 00 00
00 00 00 00 00
