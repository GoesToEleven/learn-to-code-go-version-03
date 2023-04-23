package main

import (
	"fmt"
)

func main() {
	var s1 interface{} = "James"
	var s2 any = "Bond"

	if v, ok := s1.(string); ok {
		fmt.Println(s1.(string))
		fmt.Println(v)
	}

	if _, ok := s2.(string); ok {
		fmt.Println(s2.(string))
	}
}