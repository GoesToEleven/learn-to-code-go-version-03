package main

import "fmt"

func main() {
	s := "Moon 🚀"

	for i, r := range s {
		xb := []byte(string(r))
		for _, b := range xb {
			fmt.Printf("%d \t %d \t %x \t %s \n", i, b, b, string(r))
		}
	}
}
