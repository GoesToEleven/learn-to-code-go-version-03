package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// string to be parsed into time
	// Auguest 11 2009

	t, err := time.Parse("January 02, 2006", "August 11, 2009")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(t)
}