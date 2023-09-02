package main

import "fmt"

func main() {
	s := "Hello Gophers!🚀"
	fmt.Println(s)

	for _, rune := range s {
		// fmt.Println("--")
		// fmt.Printf("%v \t %T \n\n", string(rune), string(rune))

		runeBytes := []byte(string(rune))

		fmt.Println("--")
		for _, byte := range runeBytes {
			fmt.Printf("%v \t %X \t %d \n", string(rune), byte, byte)
		}
	}

	// when you range over a string, you get
	// the runes, aka, each individual character
	// for _, rune := range s {
	// 	// fmt.Printf("%v \t %T \n", rune, rune)
	// 	// fmt.Printf("%v \t %T \n\n", string(rune), string(rune))

	// 	runeBytes := []byte(string(rune))

	// 	fmt.Println("--")
	// 	for _, byte := range runeBytes {
	// 		fmt.Printf("%v \t %X \t %d \n", string(rune), byte, byte)
	// 	}
	// }
}

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

/*
UTF-8 stands for "Unicode Transformation Format - 8-bit."
It is a variable-width character encoding that can represent every character in the Unicode standard.
UTF-8 has become the dominant character encoding for the World Wide Web, accounting for more than 90% of all web pages.

UTF-8 uses one to four bytes for each symbol.
The first 128 characters (US-ASCII) need one byte.
The next 1,920 characters need two bytes to encode, which covers the remainder of almost all Latin-script alphabets,
and also Greek, Cyrillic, Coptic, Armenian, Hebrew, Arabic, Syriac, Thaana and N'Ko alphabets, as well as
Combining Diacritical Marks. Three bytes are needed for the rest of the Basic Multilingual Plane (
which contains virtually all characters in common use). Four bytes are needed for characters in the other planes
of Unicode, which include various historic scripts, mathematical symbols, and emoji (pictographic symbols).

Here are some advantages of UTF-8:

1. Backward Compatibility:
The first 128 characters of Unicode, which correspond one-to-one with ASCII,
are encoded using a single byte with the same binary value as ASCII, making any ASCII text valid UTF-8 encoded Unicode as well.

2. Self-Synchronization:
If bytes get lost due to issues like network errors,
it's easy to determine the start of the next UTF-8-encoded code point and resynchronize.

3. Efficiency for Latin Text:
For languages that primarily use Latin characters (like English),
UTF-8 is efficient in terms of storage.

4. Wide Adoption:
Given its compatibility and efficiency, UTF-8 is widely adopted in a variety of applications,
from web pages to file storage, making it a de facto standard for text files.

5. Universal:
It can represent any character in the Unicode standard, yet it's backward-compatible with ASCII.
*/
