package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The sum is", x)
}

// func (r receiver) identifier(p parameters) (return(s)) {code}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
