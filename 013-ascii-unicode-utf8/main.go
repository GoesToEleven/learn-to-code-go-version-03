package main

import "fmt"

func main() {

	fmt.Println("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")

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
	*/
}
