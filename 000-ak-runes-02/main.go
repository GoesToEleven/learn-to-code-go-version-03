package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "G😎"
	fmt.Println(s)

	fmt.Println("len gives us the number of bytes\t", len(s))

	fmt.Println("RuneCountInString gives us the number of runes, eg characters\t", utf8.RuneCountInString(s))

	fmt.Printf("%c \t %T\n", s[0], s[0])
	fmt.Printf("%c \t %T\n", s[1], s[1])

	for i, v := range s {
		fmt.Println(i, v)
	}
}

/*
In the Go programming language, a string is a sequence of bytes that represent characters/runes.
It is a built-in data type that is used to store and manipulate text data.

Strings in Go are represented as a sequence of bytes and are immutable,
which means that once a string is created, its value cannot be changed.
However, it is possible to create a new string by concatenating two or more existing strings.

In Go, strings are enclosed in double quotes (" ") or backticks ( ), and can contain any Unicode character.
For example:

name := "Alice"

message := `Hello, world!
My message is now on two lines.
`

emoji := "😀"


*/
