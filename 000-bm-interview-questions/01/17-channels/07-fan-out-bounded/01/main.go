package main

import "fmt"

func main() {
	// sets index position 20 for last item
	work := []string{"paper", "paper", "paper", "paper", 20: "paper"}
	for i, w := range work {
		fmt.Printf("%#v - %#v \n", i, w)
	}
}
