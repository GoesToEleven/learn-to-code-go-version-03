package main

import (
	"fmt"
	"mymodule/000-bc-regexp-01-example-doc/tokenizer"
)

func main() {
	ts := tokenizer.Tokenize("Let's go to the beach")
	fmt.Println(ts)
}

/*
to run

go test ./...

or 

go test -v ./...

*/
