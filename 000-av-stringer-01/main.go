package main

import "fmt"

type Phrase1 struct{ saying string }

func (p1 Phrase1) String() string {
	return fmt.Sprint("phrase1: ", p1.saying)
}

// type Phrase2 string
// func (p2 Phrase2) String() string {
// 	return fmt.Sprint("phrase2: ", p2)
// }

func main() {
	var s1 Phrase1 = Phrase1{"happiness"}
	fmt.Println(s1)

	// var s2 Phrase2 = "happiness"
	// fmt.Println(s2)
}

// stringer interface
// https://pkg.go.dev/fmt#Stringer

// go vet .
// go test -v
// https://staticcheck.io/
