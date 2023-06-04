package main

import "fmt"

func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)

}

/*

“defer” multiple functions in main
show that a deferred func runs after the func containing it exits.
determine the order in which the multiple defer funcs run

*/

// deferred functions run in LIFO order
// last in first out LIFO
