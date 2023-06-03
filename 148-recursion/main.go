package main

import "fmt"

func main() {
	// 4!
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))

}

func factorial(n int) int {
	fmt.Println("This is n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/*
factorial(4)		--> 4 *
factorial(4-1)		--> 3 *
factorial(3-1)		--> 2 *
factorial(2-1)		--> 1 *
factorial(1-1)		--> 0 // base case
*/

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
