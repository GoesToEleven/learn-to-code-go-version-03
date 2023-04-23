package main

import (
	"fmt"
)

func main() {
	var s1 interface{} = "James"
	var s2 any = "Bond"

	fmt.Println(s1.(string))
	fmt.Println(s2.(string))
}