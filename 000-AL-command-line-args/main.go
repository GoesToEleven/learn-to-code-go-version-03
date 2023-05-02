package main

import (
	"fmt"
	"os"
)

func main() {

	// all arguments given to command line
	fmt.Println("all arguments\t", os.Args)

	// arguments beyond running main.go
	fmt.Println("beyond main.go\t", os.Args[1:])

	// all arguments beyond running main.go
	for i, v := range os.Args[1:] {
		fmt.Printf("argument %v is %v\n", i, v)
	}
}

/*
 go run main.go a b c d e f g
*/
