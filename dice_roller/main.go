package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	adv := flag.Bool("adv", false, "Roll twice and take the higher")
	dis := flag.Bool("dis", false, "Roll twice and take the lower")

	// Parse command line arguments
	flag.Parse()

	// Check if a dice expression was provided
	if len(os.Args) <= 1 {
		fmt.Println("No dice expression provided. Please provide a dice expression in the format XdY[+-]N.")
		return
	}

	// Get the dice expression from the command line arguments
	input := os.Args[1]

	// Rest of your code...
}
