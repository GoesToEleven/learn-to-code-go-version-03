# "Why did you create a new language?"
"Go was born out of frustration with existing languages and environments for the work we were doing at Google. Programming had become too difficult and the choice of languages was partly to blame. One had to choose either ***efficient compilation, efficient execution, or ease of programming;*** all three were not available in the same mainstream language. Programmers who could were choosing ease over safety and efficiency by moving to dynamically typed languages such as Python and JavaScript rather than C++ or, to a lesser extent, Java.

We were not alone in our concerns. After many years with a pretty quiet landscape for programming languages, Go was among the first of several new languages—Rust, Elixir, Swift, and more—that have made programming language development an active, almost mainstream field again.

Go addressed these issues by attempting to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aimed to be modern, with support for networked and multicore computing. Finally, working with Go is intended to be fast: it should take at most a few seconds to build a large executable on a single computer. To meet these goals required addressing a number of linguistic issues: an expressive but lightweight type system; concurrency and garbage collection; rigid dependency specification; and so on. These cannot be addressed well by libraries or tools; a new language was called for."
[Golang FAQ](https://go.dev/doc/faq#creating_a_new_language)

# Video
[SOON TO BE PUBLISHED Here is the video on youtube]()

# Takeaways
1. Functions and modularity
1. Parameters and arguments
1. Syntax: func, receiver, identifier, [parameter(s)], [return(s)], code block
1. No named returns
1. Variadic parameters, variadic arguments, and unfurling a slice
1. Multiple returns and the blank identifier
1. Anonymous functions
1. Functions are a type: you can use functions as parameters, arguments, & returns
1. Defer
1. Panic
1. Closures
1. Type error and error checking
1. Syntactic sugar

# Table of Contents

1. [Function overview](#function-overview)
1. [Function Declaration](#function-declaration)
1. [Parameters](#parameters)
1. [Return Values](#return-values)
1. [Multiple Return Values](#multiple-return-values)
1. [Named Return Values](#named-return-values)
1. [Function Calls](#function-calls)
1. [Function Variadic Parameters](#function-variadic-parameters)
1. [Anonymous Functions](#anonymous-functions)
1. [Function as a Type](#function-as-a-type)
1. [Defer and Panic](#defer-and-panic)
1. [Closures](#closures)
1. [Type error error checking and returning errors](#type-error-error-checking-and-returning-errors)
1. [Syntactic sugar](#syntactic-sugar)
1. [Code review check](#code-review-check)
1. [Hachikō waited ❤️ ♥️ ⽝](#hachikō-waited)

# Function overview

Functions are blocks of code that perform a specific task or a set of related tasks. Functions are a fundamental building block of programs and are used to encapsulate and organize code for better modularity and reusability. Here's a detailed explanation of functions in Go:

# Function Declaration
   To declare a function in Go, you use the `func` keyword followed by the function name, a list of parameters enclosed in parentheses, and an optional return type. The general syntax is:

   ```go
   func functionName(parameter1 type, parameter2 type, ...) return_type {
       // Function body
   }
   ```

   For example:

   ```go
   func add(x int, y int) int {
       return x + y
   }
   ```

# Parameters
   Functions can take zero or more parameters, which are the values passed to the function when it is called. Each parameter consists of a name and a type. In the example above, `add` takes two integer parameters, `x` and `y`.

# Return Values
   Functions can return zero or one value (or none, in which case you can omit the return type). The `return` statement is used to specify the value to be returned. In the example above, `add` returns an integer.

# Multiple Return Values
   Go allows functions to return multiple values. This is useful in cases where you want to return more than one result from a function. For example:

   ```go
   func divide(x, y float64) (float64, error) {
       if y == 0 {
           return 0, errors.New("division by zero")
       }
       return x / y, nil
   }
   ```

   In this example, the `divide` function returns both a floating-point result and an error (if the division by zero occurs).

# Named Return Values
   Go allows you to specify names for the return values in a function's signature. Named return values are automatically initialized and can be used as regular variables within the function. They are particularly useful when dealing with complex functions or error handling. For example:

   ```go
   func divide(x, y float64) (result float64, err error) {
       if y == 0 {
           err = errors.New("division by zero")
           return
       }
       result = x / y
       return
   }
   ```

# Function Calls
   To call a function in Go, you use the function name followed by a list of arguments enclosed in parentheses. If the function has multiple return values, you can capture them in variables. For example:

   ```go
   sum := add(3, 5)           // Call the add function and assign its result to sum
   result, err := divide(8, 2) // Call the divide function and capture both result and error
   ```

# Function Variadic Parameters
   Go supports variadic functions, which allow you to pass a variable number of arguments to a function. To define a variadic parameter, you use an ellipsis (`...`) followed by the parameter type. For example:

   ```go
   func calculateSum(xv ...int) int {
       total := 0
       for _, v := range xv {
           total += v
       }
       return total
   }
   ```

   You can call this function with any number of integers.

# Anonymous Functions
   Go supports anonymous functions. These are functions without a name, and they can be assigned to variables and used as arguments to other functions.

   ```go
   add := func(x, y int) int {
       return x + y
   }

   result := add(3, 5)
   ```

# Function as a Type
   In Go, functions can also be used as types. This allows you to define functions that take other functions as arguments or return them as results. It's a powerful feature for implementing higher-order functions and callbacks.

   ```go
   type MathFunc func(int, int) int

   func operate(x, y int, op MathFunc) int {
       return op(x, y)
   }
   ```

   You can then pass functions as arguments to `operate`.

# Defer and Panic
    Go provides two special built-in functions called `defer` and `panic` for handling exceptional situations and resource management. `defer` is used to schedule a function call to be executed just before the function returns, while `panic` is used to trigger a run-time error and unwind the stack.

	"The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller. To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions. This continues until all functions in the executing goroutine have stopped, in reverse order. At that point, the program is terminated with a non-zero exit code. This termination sequence is called panicking and can be controlled by the built-in function recover."
	[func panic](https://pkg.go.dev/builtin#panic)

# Closures
A closure is a function value that references variables from outside its own scope. These variables are then "closed over" by the closure, allowing them to maintain state even after the surrounding function has finished execution.

Here's an example:

```go
package main

import "fmt"

// Counter returns an anonymous function 
// The anonymous func is of this type: func() int
func counter() func() int {
    // count is closed over by the returned func
    count := 0
    
    return func() int {
        count++
        return count
    }
}

func main() {
    counterA := counter()
    counterB := counter()

    fmt.Println(counterA())  // Output: 1
    fmt.Println(counterA())  // Output: 2
    fmt.Println(counterA())  // Output: 3

    fmt.Println(counterB())  // Output: 1
    fmt.Println(counterB())  // Output: 2

    // Note that counterA and counterB maintain separate 'count' states.
}
```

In this example, the function `counter` returns another function of type `func() int`. The returned function closes over the `count` variable, meaning it retains access to that variable's state even after `counter()` has returned. I think about it like this, "The returned func CLOSES OVER the count variable." I also think about it like this, "The returned func takes the count variable with it."

When you call `counterA()` and `counterB()`, you get two separate instances of the closure, each with its own "closed-over" `count` variable. So even though both closures were generated by the same function (`counter`), they maintain separate state.

This can be particularly useful when you need to maintain state across function calls without using global variables, or when you need to dynamically generate functions that behave slightly differently based on parameters or runtime conditions.

Closure is also necessary when working with goroutines. Often you will need to send a variable along with a goroutine, and we use closure to do this.

Functions in Go play a central role in structuring code and promoting clean, maintainable, and efficient software. Understanding how to declare, define, and use functions is crucial for writing effective Go programs.

# Type error error checking and returning errors

In Go, error handling is fundamental to building robust applications. 

Below are some of the key concepts and techniques to understand about error handling in Go.

### Error Checking

In Go, errors are values that implement the `error` interface:

```go
type error interface {
    Error() string
}
```

A typical Go function that can return an error usually has a signature like this:

```go
func DoSomething() (ResultType, error) {
    // function body
}
```

To handle errors, you typically use an `if` statement to check if an error has occurred:

```go
result, err := DoSomething()
if err != nil {
    // Handle error
}
```

### Type Error

Go is a statically-typed language, so type errors sometimes occur at compile-time. However, you can also encounter them at runtime, for instance, when you're using interfaces and type assertions:

```go
var i interface{} = "a string"

val, ok := i.(int)
if !ok {
    // Handle the error: i is not of type int
}
```

### Returning Errors from Functions

When defining a function that might fail, you usually return an error as the last return value. The idiomatic way to indicate an error is to return `nil` if everything went fine, or an `error` if something went wrong.

Here's an example that uses the standard library's `errors` package:

```go
import "errors"

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

### `log.Fatalf`

The `log.Fatalf` function from Go's standard library is a convenient way to log an error message and immediately exit the program:

```go
import "log"

func main() {
    result, err := Divide(10, 0)
    if err != nil {
        log.Fatalf("An error occurred: %v", err)
    }
    log.Println("Result:", result)
}
```

When `log.Fatalf` is called, it prints the formatted error message and stops the program with an exit code that indicates an error occurred. Note that `defer` statements will not be run when calling `log.Fatalf`.

In summary, Go encourages an explicit, straightforward approach to error-handling that makes it easy to understand what a piece of code is doing. Always check for errors and handle them in a way that's appropriate for your application.


```go
func CountWords() {
	// Open the file
	file, err := os.Open("gatsby.txt")
	if err != nil {
		log.Fatalf("error opening file: %s \n", err)
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Count the words
	count := 0
	for scanner.Scan() {
		count++
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s \n", err)
	}

	// Print the word count
	fmt.Printf("Total words: %d\n", count)
}
```
# Syntactic sugar

In computer science, "syntactic sugar" refers to syntax in a programming language that is designed to make things easier to read or express. It doesn't introduce a new feature in the language but provides a more convenient way of utilizing one that already exists.

In the context of Go, one of the syntactic sugar features is automatically dereferencing pointers to structs when accessing fields or calling methods.

Let's look at a simple example:

```go
type user struct {
	first string
}

func ffUser(s string) (*user, error) {
	u := user{s}
	return &u, nil
}

func RunSugar() {

	u, err := ffUser("Jenny")
	if err != nil {
		log.Fatalf("error creating user %s \n", err)
	}

	// address of a struct
	fmt.Println(u)
	fmt.Printf("%T \n", u)
	fmt.Println((*u).first)

	// syntatic sugar
	fmt.Println(u.first)
}
```

Here, we have a `User` struct with a `First` field. We create a pointer `p` to a value `u` of type `User`. Normally, to access fields in a struct through a pointer, you would have to dereference the pointer using the `*` operator, like so: `(*p).First`.

However, Go allows you to use syntactic sugar to skip this step, directly using `p.First` instead. The Go runtime automatically understands that you want to dereference the pointer and access the `First` field of the underlying struct. This makes the code easier to read and write.

In summary, in Go, you can access fields from a pointer to a struct as if you were directly working with the struct, thanks to this syntactic sugar.

# Code review check

### No named returns
- named returns decrease readability
    - an empty `return` is not idiomatic

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# Hachikō waited

![Hachikō](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/07-functions/image/hachiko.png)

The Hachikō statue in Shibuya, Tokyo, is one of the city's most famous landmarks. It commemorates Hachikō, an Akita dog who is remembered for his remarkable loyalty to his owner. The story goes that Hachikō would wait for his owner, Hidesaburo Ueno, every day at Shibuya Station to return from work. Unfortunately, Ueno passed away suddenly in 1925, but Hachikō continued to wait at the station for his owner every day until his own death nearly 10 years later in 1935. This act of devotion touched many people, and Hachikō became a national symbol of loyalty and fidelity.

The bronze statue was erected in 1934, a year before Hachikō's death, near the Shibuya Station where he waited. 

![Hachikō](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/07-functions/image/hachiko2.jpeg)

photo source: https://en.wikipedia.org/wiki/File:Faithful_Dog_Hachiko_Photo.png
