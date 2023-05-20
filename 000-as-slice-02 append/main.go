package main

import "fmt"

func main() {
	fmt.Println(concatOne([]string{"A", "B"}, []string{"C", "D", "E"}))
	fmt.Println(concatTwo([]string{"A", "B"}, []string{"C", "D", "E"}))
}

func concatOne(s1, s2 []string) []string {
	return append(s1, s2...)
}

func concatTwo(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}
