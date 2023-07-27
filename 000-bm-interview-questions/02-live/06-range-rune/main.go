/*
	https://blog.golang.org/strings

	Go source code is always UTF-8.

	a string is a
	[]byte
	is
	code points, called runes
	is
	characters

*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "𒆸世界 means world 𒆸"

	// UTFMax is 4 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	for i, r := range s {

		rl := utf8.RuneLen(r)
		si := i + rl

		// every ARRAY in Go is just a SLICE waiting to happen
		copy(buf[:], s[i:si])

		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
