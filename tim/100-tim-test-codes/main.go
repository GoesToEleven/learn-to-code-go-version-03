package main

import "fmt"

func main() {
	fmt.Println("abc")
	type Bunch[E any] []E
	b := Bunch[int]{1, 2, 3}
	fmt.Printf("b is of type %#v and with value %v", b, b)
}
