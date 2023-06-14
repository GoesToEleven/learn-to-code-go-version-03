package main

import (
	"fmt"
	"mymodule/000-bc-test-doc-01/tokenizer"
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
