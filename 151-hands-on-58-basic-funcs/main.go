package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "James Bond"
}

/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/
