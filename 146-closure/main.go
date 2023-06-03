package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
	fmt.Println(f()) // 6

	g := incrementor()
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2
	fmt.Println(g()) // 3
	fmt.Println(g()) // 4
	fmt.Println(g()) // 5
	fmt.Println(g()) // 6
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
