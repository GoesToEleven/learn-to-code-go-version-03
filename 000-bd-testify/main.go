package main

import (
	"fmt"
	"mymodule/000-bd-testify/tokenizer"
)

func main() {
	ts := tokenizer.Tokenize("Let's go to the beach")
	fmt.Println(ts)
}

//https://pkg.go.dev/github.com/stretchr/testify

/*
to run

go test ./...

or 

go test -v ./...

*/
