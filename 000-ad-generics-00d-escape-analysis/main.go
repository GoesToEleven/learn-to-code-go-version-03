package main

import "fmt"

func main() {
	x := 42
	f := foo(&x)
	fmt.Println(f)
	fmt.Println("hello")
}

func foo(n *int) *int {
	y := 2 * *n
	return &y
}

// go build -gcflags="-m"
// go build -gcflags="-m -m"

/*
$ go build -gcflags="-m"
# mymodule/000-ad-generics-00d-escape-analysis
.\main.go:11:6: can inline foo
.\main.go:12:13: inlining call to fmt.Println
.\main.go:5:6: can inline main
.\main.go:7:5: inlining call to foo
.\main.go:7:5: inlining call to fmt.Println
.\main.go:7:5: ... argument does not escape
.\main.go:7:5: "Hello" escapes to heap
.\main.go:12:13: ... argument does not escape
.\main.go:12:14: "Hello" escapes to heap


*/
