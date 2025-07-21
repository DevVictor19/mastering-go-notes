package main

import "fmt"

// %s	the uninterpreted bytes of the string or slice
// %c	the character represented by the corresponding Unicode code point
// %x	base 16, lower-case, two characters per byte
// %d	base 10
// %b	base 2

func main() {
	r := '€'                     // a rune is a int32 used to representing a single unicode point
	text := "My text €"          // a string is just an array of bytes
	bytes := []byte("My text €") // Unicode uses 3 bytes (226, 130, 172)

	fmt.Println("rune:", r) // unicode point
	b := text[len(text)-1]
	fmt.Println("last byte:", b) // 172

	fmt.Println("unicode representation of each char:")

	for _, v := range text {
		fmt.Println(v) // ASCII/Unicode representation
	}

	fmt.Println("bytes of string:")
	// A unicode character can use more than one byte
	for _, v := range bytes {
		fmt.Println(v) // Byte representation (decimal)
	}
}

/* (Unicode values)
77     // 'M'
121    // 'y'
32     // ' '
116    // 't'
101    // 'e'
120    // 'x'
116    // 't'
32     // ' '
8364   // '€' (U+20AC)
*/

/* (Decimal values of bytes)
77     // 'M'
121    // 'y'
32     // ' '
116    // 't'
101    // 'e'
120    // 'x'
116    // 't'
32     // ' '
226    // first byte of '€'
130    // second byte of '€'
172    // third byte of '€'
*/
