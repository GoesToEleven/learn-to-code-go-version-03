package main

import "fmt"

func main() {
	s := "Hello Gophers!🚀"
	fmt.Println(s)

	// see emojis
	// Windows logo key + .

	/*
		ascii
		American Standard Code for Information Interchange
		2^8 (1 byte) = 256 unique values

		unicode
		2^32 (4 bytes) = 4,294,967,296 unique values
		more than enough to account for every character in every language

		utf-8
		(up to 4 bytes)
		stores unicode as binary
		If a character needs 1 byte that’s all it will use.
		If a character needs 4 bytes it will use 4 bytes.
		variable length encoding = efficient memory use
		common characters like “C” take 8 bits,
		rare characters like “💕” take 32 bits.
		Other characters take 16 or 24 bits.

		https://blog.hubspot.com/website/what-is-utf-8
		https://deliciousbrains.com/how-unicode-works/

		rune
		In Go, a "rune" is an alias for the int32 type 
		and it represents a Unicode code point.
		A Unicode code point is a unique number 
		that maps to a specific character 
		in the Unicode character set. 
		Unicode code points can represent 
		standard ASCII characters (like 'A', 'B', '1', etc.), 
		but they can also represent a variety of other characters 
		like 'é', 'α', '世', and many more.
	*/

	// when you range over a string, you get
	// the runes, aka, each individual character
	for _, rune := range s {
		// fmt.Printf("%v \t %T \n", rune, rune)
		// fmt.Printf("%v \t %T \n\n", string(rune), string(rune))

		runeBytes := []byte(string(rune))
		
		fmt.Println("--")
		for _, byte := range runeBytes {
			fmt.Printf("%v \t %X \t %d \n", string(rune), byte, byte)
		}
	}
}
