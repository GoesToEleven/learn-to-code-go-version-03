package main

import "fmt"

func main() {
	foo()

	bar("Todd")

	s := aloha("Todd")
	fmt.Println(s)

	s1, n := dogYears("Todd", 40)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

/*

func syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params, 2 returns
*/
