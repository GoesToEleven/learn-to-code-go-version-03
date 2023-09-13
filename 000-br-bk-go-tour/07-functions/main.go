package main

import (
	"fmt"
	"mymodule/000-br-bk-go-tour/07-functions/sofun"
)

func main() {

	// syntax
	// func (r receiver) identifier(parameter(s)) returns { code block }

	// #1
	// basic
	hello()

	// #2
	// return
	s2 := hello2()
	fmt.Println(s2)

	// #3
	// return
	s3, i3 := hello3()
	fmt.Println(s3, i3)

	// #4
	// parameter
	p4 := param4("Jenny")
	fmt.Println(p4)

	// #5
	// parameter
	p5 := param5(41, 1)
	fmt.Println(p5)

	// #6
	// variadic argument
	p6 := param6(41, 42, 43, 44, 45)
	fmt.Println(p6)

	// #7
	// unfurling a slice
	xi := []int{46, 47, 48, 49, 50}
	x := param6(xi...)
	fmt.Println(x)

	// #8
	// anonymous func
	add := func(x, y int) int {
		return x + y
	}
	af := add(3, 5)
	fmt.Println(af)

	// #9
	// closure
	counterA := cloEx()
	counterB := cloEx()

	fmt.Println(counterA()) // Output: 1
	fmt.Println(counterA()) // Output: 2
	fmt.Println(counterA()) // Output: 3

	fmt.Println(counterB()) // Output: 1
	fmt.Println(counterB()) // Output: 2

	// #10
	// method
	p1 := person{"Bond, James Bond"}
	p1.sayHello()

	// #11
	// multiple returns - error
	sofun.CountWords()
}

// #1
// basic
func hello() {
	fmt.Println("Hello gophers!")
}

// #2
// return
func hello2() string {
	return "Hello gophers!"
}

// #3
// returns
func hello3() (string, int) {
	return "Hello gophers!", 42
}

// #4
// parameters
func param4(fn string) string {
	s := "Hello" + fn
	return s
}

// #5
// parameters
func param5(a int, b int) int {
	return a + b
}

// #6
// variadic parameter
func param6(xv ...int) int {
	fmt.Println("---------------VARIADIAC PARAMETERS")
	fmt.Printf("%T \t %v \n", xv, xv)
	total := 0
	for _, v := range xv {
		total += v
	}
	return total
}

// #9
// closure
func cloEx() func() int {
	// count is closed over by the returned func
	count := 0

	return func() int {
		count++
		return count
	}
}

// #10
// method
type person struct {
	intro string
}

func (p person) sayHello() {
	fmt.Println("Hello, I am", p.intro)
}