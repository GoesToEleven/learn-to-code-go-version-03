package main

import (
	"bytes"
	"log"
	"os"
)

func main() {

	xb1, err := os.ReadFile("one.txt")
	if err != nil {
		log.Fatalf("%s", err)
	}

	xb2 := bytes.Replace(xb1, []byte("wanna"), []byte("want to"), -1)

	if err = os.WriteFile("two.txt", xb2, 0666); err != nil {
		log.Fatalf("%s", err)
	}
}
