package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "mom"
	s2 := "topspot"
	s3 := "abbbba"
	s4 := "notpali"

	fmt.Printf("%v a palindrome? \t %t\n", s1, isPali(s1))
	fmt.Printf("%v a palindrome? \t %t\n", s2, isPali(s2))
	fmt.Printf("%v a palindrome? \t %t\n", s3, isPali(s3))
	fmt.Printf("%v a palindrome? \t %t\n", s4, isPali(s4))
}

func isPali(s string) bool {
	r := utf8.RuneCountInString(s)
	mid := r / 2
	for i := 0; i < mid; i++ {
		fmt.Println(i)
		if s[i] != s[r-i-1] {
			return false
		}
	}
	return true
}
