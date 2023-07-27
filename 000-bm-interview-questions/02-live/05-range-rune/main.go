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
)

func main() {

	s := "𒆸世界 means world 𒆸"

	for i, r := range s {
		fmt.Printf("%d \t %d \t %q \n", i, r, r)
	}
}
