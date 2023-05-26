package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		// print i when it is an odd number
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
