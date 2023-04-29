package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init(){
	fmt.Println("Begin initialization")
}

func main() {
	//SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// CONDITIONAL
	// if statements
	// switch statements

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

	/*
		"If" statements specify the conditional execution of two branches
		according to the value of a boolean expression. If the expression evaluates
		to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
	*/
	// https://go.dev/ref/spec#If_statements

	/*
		Comparison operators
		Comparison operators compare two operands and yield an untyped boolean value.

		==    equal
		!=    not equal
		<     less
		<=    less or equal
		>     greater
		>=    greater or equal
	*/
	// https://go.dev/ref/spec#Comparison_operators

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 {
		fmt.Println("x is not 42")
	}

	/*
		Logical operators
		Logical operators apply to boolean values
		and yield a result of the same type as the operands.
		The right operand is evaluated conditionally.

		&&    conditional AND    p && q  is  "if p then q else false"
		||    conditional OR     p || q  is  "if p then true else q"
		!     NOT                !p      is  "not p"
	*/
	// https://go.dev/ref/spec#Logical_operators

	/*
		The expression [evaluated in an if statement ]may be preceded by a simple statement,
		which executes before the expression is evaluated.
	*/
	// https://go.dev/ref/spec#If_statements
	/*
		The comma ok idiom is also like this
	*/
	// https://go.dev/play/p/OXGzjxVkag0

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x")
	}

	// concurrency
	// switch on channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(r.Int63n(250))
	d2 := time.Duration(r.Int63n(250))
	// fmt.Printf("%v \t %T\n", d1, d1)
	// time.Sleep(d1 * time.Millisecond)
	// fmt.Println("Done")

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

	// LOOPS / INTERATIVE
	// for statements

	/*
		The Go for loop is similar to—but not the same as—C's.
		It unifies for and while and there is no do-while.
		There are three forms, only one of which has semicolons.

		// Like a C for
		for init; condition; post { }

		// Like a C while
		for condition { }

		// Like a C for(;;)
		for { }

	*/
	// https://go.dev/doc/effective_go#for

	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first  for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// break
	// takes you out of the loop
	for {
		fmt.Printf("y is %v \t\t third  for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}

	// nested loops
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}

	// for range loop
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structures - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}
}
