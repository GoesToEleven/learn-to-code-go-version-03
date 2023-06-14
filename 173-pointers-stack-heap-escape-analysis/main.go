package main

import "fmt"

func main() {
	a := 1
	fmt.Println(&a)
}

/*
escapes to the heap
variable shared between main() and Println()

moved to heap
variable moved to the heap
*/
