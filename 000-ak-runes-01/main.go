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
