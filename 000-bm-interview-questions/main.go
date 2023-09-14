package main

import "fmt"

func main() {
	// Create a slice to store the questions
	var questions []string

	// ------------------------------------------------
	// BEGINNER QUESTIONS
	// ------------------------------------------------
	questions = append(questions, `What is the blank identifier?`)
	questions = append(questions, "What's the difference between declare, assign, and initialize?")
	questions = append(questions, "Explain allocation and initialization.")
	questions = append(questions, "What is the difference between make([]int, 10), make([]int, 0, 10), make([]int, 10, 10)")
	questions = append(questions, "What types can a map use as a key in the Go programming language?")
	questions = append(questions, "What does it mean to write idiomatic Go code?")
	questions = append(questions, "Which do you choose: performance or readability?")
	questions = append(questions, "Why was Go created, and who created it?")
	questions = append(questions, "Can you explain what is meant by 'Go is strongly typed?'")
	questions = append(questions, "What is the var keyword used for and when do you use it?")
	questions = append(questions, `When dealing with computer architecture, what does the 'word size' mean?`)
	questions = append(questions, "How does a computer work?")
	questions = append(questions, "What is a compiler?")
	questions = append(questions, "What is a garbage collector?")
	questions = append(questions, "Can you give an example of a problem you've solved using Go? What challenges did you encounter, and how did you overcome them?")
	questions = append(questions, "Explain how package management works in Go.")
	questions = append(questions, `What is the difference between make and new?`)
	questions = append(questions, `Tell us about bytes, code points, and characters in relation to strings and UTF-8.`)
	questions = append(questions, `What is embedding a struct and inner-type promotion?`)

	questions = append(questions, `Fix this code so that type HUMAN is embedded in type SECRETAGENT

	type human struct {
		name  string
		email string
	}
	
	type secretAgent struct {
		person human
		id  string
	}

	`)
	// fixed: https://go.dev/play/p/zO0ZBicZM7u

	questions = append(questions, `Fix this code working with a map:
	package main

	import "fmt"
	
	func main() {
		var m map[string]int
		m["b"] = 42
		fmt.Println(m)
	}
	
	`)
	// not fixed: https://go.dev/play/p/d001JBoJcAV
	// fixed: https://go.dev/play/p/RCVcb_4z2oK

	questions = append(questions, `What's problematic with this code working with append:

	func slice() {
		xi := make([]int, 10)
		for i := 0; i < 10; i++ {
			xi1 = append(xi1, i)
		}
		fmt.Println(xi1)
	}		
	`)
	// not fixed: https://go.dev/play/p/04NLYE1I7W4
	// fixed: https://go.dev/play/p/guYNTv-lfFi

	questions = append(questions, `What will this code print:
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		xs := sports[1:3]
		xs[0] = "CHANGED"

		inspectSlice(sports)
		inspectSlice(xs)
	}
	
	func inspectSlice(xs []string) {
		fmt.Printf("len: %v \ncap: %v \n", len(xs), cap(xs))
		for i := range xs {
			fmt.Printf("%p \t %v \n", &xs[i], xs[i])
		}
	}
	`)
	// https://go.dev/play/p/Rw5cmIrXlwT

	questions = append(questions, `Fix this code so that when xs[0] is changed this doesn't change 'sports':
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		xs := sports[1:3]
		xs[0] = "CHANGED"
		inspectSlice(sports)
		inspectSlice(xs)
	}
	
	func inspectSlice(xs []string) {
		fmt.Printf("len: %v \ncap: %v \n", len(xs), cap(xs))
		for i := range xs {
			fmt.Printf("%p \t %v \n", &xs[i], xs[i])
		}
	}
	`)
	// not fixed: https://go.dev/play/p/Rw5cmIrXlwT
	// fixed: https://go.dev/play/p/N-32AbjJJZw

	// ------------------------------------------------
	// INTERMEDIATE QUESTIONS
	// ------------------------------------------------
	questions = append(questions, "What is the comma ok idiom?")
	questions = append(questions, "Is Go an Object Oriented language?")
	questions = append(questions, `Somebody says that 'a string is a two word data structure' - what does this mean?`)
	questions = append(questions, `Somebody says that 'a slice is a three word data structure' - what does this mean?`)
	questions = append(questions, `Tell me the difference between a nil slice and an empty slice:
	

	package main

	import (
		"fmt"
	)

	func main() {
		var a []string
		b := []string{}

		fmt.Println(a, len(a), cap(a), a == nil)  // Output: [] 0 0 true
		fmt.Println(b, len(b), cap(b), b == nil)  // Output: [] 0 0 false

		a = append(a, "item")
		b = append(b, "item")

		fmt.Println(a, len(a), cap(a), a == nil)  // Output: [item] 1 1 false
		fmt.Println(b, len(b), cap(b), b == nil)  // Output: [item] 1 1 false
	}

	
	`)
	questions = append(questions, "What are typed and untyped constants, which do you prefer and why, and how do they relate to numeric literals?")
	questions = append(questions, "How does Go handle error handling compared to other languages?")
	questions = append(questions, "Conversion or casting, and why?")
	questions = append(questions, "What is nil?")
	questions = append(questions, "How do we do benchmarking in Go?")
	questions = append(questions, "What is the difference between an int and a uint, and how does this relate to mechanical sympathy?")
	questions = append(questions, "Put these in the correct order: centisecond, second, nanosecond, millisecond, decisecond, microsecond")
	questions = append(questions, "What are pointers? Show us pointers at work in code.")
	questions = append(questions, "Explain value semantics and pointer semantics. What are rules-of-thumb for using one versus the other?")
	questions = append(questions, "What is the stack, and what is the heap?")
	questions = append(questions, "What is escape analysis?")
	questions = append(questions, "What is an interface?")
	questions = append(questions, "What are method sets, and how do you use them?")
	questions = append(questions, "What is a type set?")
	questions = append(questions, "What is concrete data?")
	questions = append(questions, "Explain generics.")
	questions = append(questions, "What is the difference between concurrency and parallelism?")
	questions = append(questions, "What is a Goroutine?")
	questions = append(questions, "Describe a time when you used goroutines and channels?")
	questions = append(questions, `Should you use buffered channels? Why or why not?`)
	questions = append(questions, "What is the difference between switch and select?")
	questions = append(questions, "Explain internal, external, and data latencies in Go.")
	questions = append(questions, `Explain the differences between:
	var wg sync.WaitGroup
	atomic.AddInt64(&counter, 1)
	var mu sync.Mutex	
	`)

	questions = append(questions, `Will this print the same for both print statements:
	
	type person struct {
		age int
	}

	people := make([]person, 2)

	p1 := people[1]
	p1.age++
	
	fmt.Println(p1.age)			// ???
	fmt.Println(people[1].age)	// ???
	
	`)
	// initial code: https://go.dev/play/p/CgbIkKrWS0s
	// insights: https://go.dev/play/p/9RbDIRs1Xda

	questions = append(questions, `Fix this code with waitgroups and a deadlock:
	
	func main() {
		var wg sync.WaitGroup
		wg.Add(3)
	
		go func() {
			fmt.Println("process one")
			wg.Done()
		}()
	
		go func() {
			fmt.Println("process two")
			wg.Done()
		}()
	
		wg.Wait()
	}

	`)
	// not fixed: https://go.dev/play/p/JOQJ91Dmh-d
	// fixed: https://go.dev/play/p/fT1R6oX2NXK

	questions = append(questions, `Fix this codewith waitgroups and a race condition:
	
	func main() {
		var counter int
		var wg sync.WaitGroup
		wg.Add(3)
	
		go func() {
			counter++
			wg.Done()
		}()
	
		go func() {
			counter++
			wg.Done()
		}()
	
		wg.Wait()
		fmt.Println("Final Counter:", counter)
	}

	`)
	// not fixed: https://go.dev/play/p/cJbeHJvKqGt
	// fixed: https://go.dev/play/p/3a5SGAvsZ_6

	questions = append(questions, `What does this line of code do, and when might you use it:
	
	atomic.AddInt64(&counter, 1)

	`)

	questions = append(questions, `Explain what this code will print:
	
	// var sports [5]string
	sports := make([]string, 5)
	sports[0] = "ski"
	sports[1] = "surf"
	sports[2] = "swim"
	sports[3] = "sail"
	sports[4] = "sumo wrestling"

	for i, v := range sports {
		sports[i] = "biking"
		fmt.Println(v)			// ???
	}

	fmt.Println(sports)			// ???
	
	`)
	// https://go.dev/play/p/ifhWmIY5I-H

	// ------------------------------------------------
	// ADVANCED QUESTIONS
	// ------------------------------------------------
	questions = append(questions, `Teach us something most people don't know about Go.`)
	questions = append(questions, `Explain alignment, padding bytes, and mechanical sympathy.`)
	questions = append(questions, `What don't you like about the Go programming language?`)
	questions = append(questions, `What's the difference between synchronization & orchestration?`)
	questions = append(questions, "How might you use a nil channel?")
	questions = append(questions, "Explain how a linked list works.")
	questions = append(questions, "What is iota? Can you show us how to use iota with bitshifting?")
	questions = append(questions, "Have you worked on a large-scale project using Go? If so, how did Go's features contribute to the project's success?")
	questions = append(questions, "What tools do you typically use for testing Go code?")
	questions = append(questions, "What is a multiplexer?")
	questions = append(questions, "What is the difference between TCP and UDP?")

	questions = append(questions, `What will the last print statement print, and why?
	
	type person struct {
		age int
	}
	
	func main() {
		people := make([]person, 2)
	
		p1 := &people[1]
		p1.age++
	
		// Add a new person
		people = append(people, person{})
	
		// Add another year for p1
		p1.age++
		
		fmt.Println(people[1].age) 	// ???
		fmt.Println(p1.age)			// ???

	`)
	// https://go.dev/play/p/PZlMtBwXQhK

	// ------------------------------------------------
	// TOOLING  QUESTIONS
	// ------------------------------------------------
	questions = append(questions, "What is Docker and what made it innovative?")
	questions = append(questions, "Tell us about containers and container orchestration.")
	questions = append(questions, "How would you set up a CI/CD pipeline for a Go project?")
	questions = append(questions, `Tell us what these do:
	cobra
	viper
	opentelemetry
	jaeger
	docker
	chi
	zshell
	`)

	// ------------------------------------------------
	// SOFT SKILL QUESTIONS
	// ------------------------------------------------
	questions = append(questions, "Tell us about your soft-skills.")
	questions = append(questions, "What makes you a valuable member of a team?")
	questions = append(questions, "Can you argue in favor of your perspective, and accept when your perspective isn't chosen?")
	questions = append(questions, "Tell us about a time you failed, and what you learned from it.")
	questions = append(questions, "Tell us about a time you succeeded, and why you succeeded.")
	questions = append(questions, "What do you do to stay current with programming?")

	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)
	questions = append(questions, ``)

	// Print out each question
	for i, question := range questions {
		fmt.Printf("Question %d: %s\n", i+1, question)
	}

	sliceDiff()
}

func sliceDiff() {
	xi1 := make([]int, 10)
	xi2 := make([]int, 0, 10)
	// xi3 := make([]int, 10, 10)

	for i := 0; i < 10; i++ {
		xi1 = append(xi1, i)
		xi2 = append(xi2, i)
		// xi3 = append(xi3, i)
	}
	fmt.Println(xi1)
	fmt.Println(xi2)
	// fmt.Println(xi3)

}

/*
The term "latency" in the context of programming generally refers to the time delay between an action and the response to that action. It's a measure of how long it takes for a system to respond to a given input or request. Latency can be influenced by many different factors, including network speed, server load, the complexity of a given task, and more. This concept applies to all programming languages, including Go.

In the context of the Go programming language, there are several different types of latencies that you might need to be concerned with:

1. Network Latency: This is the delay caused by the network infrastructure between where the Go application is running and the system it is communicating with. This type of latency is usually outside the control of the application itself but can have a significant impact on application performance.

2. Garbage Collection (GC) Latency: Go is a garbage-collected language, which means it automatically handles memory deallocation. When the garbage collector runs, it can introduce a delay or latency in the application, particularly if it leads to a "stop the world" pause where all other processing is temporarily halted. The Go team has put a lot of effort into minimizing GC latency, and it's usually quite small, but it can still be a concern in performance-critical applications.

3. Concurrency Latency: Go is well-known for its powerful concurrency primitives, including goroutines and channels. However, if not used properly, these can introduce latency. For example, if a goroutine is blocked waiting for data on a channel, this can delay the execution of that goroutine. Similarly, the overhead of creating and managing many small goroutines can also introduce latency.

4. IO Latency: This refers to the latency involved in reading and writing data, whether to a disk, a network connection, or another system. Go's standard library provides many tools for managing IO latency, such as buffered IO and asynchronous IO.

5. Scheduling Latency: This refers to the time taken by the Go scheduler to assign a goroutine to an available thread for execution. If the system is overloaded with many goroutines, the scheduling latency can increase.

6. Serialization/Deserialization Latency: When you are sending/receiving data to/from another system, you often need to convert the data to/from a format that can be sent over the network (a process often called serialization and deserialization). This process can introduce latency, particularly if the data structures involved are complex.

These are some of the main types of latency that might be encountered in a Go program, though the specifics can depend on the details of the particular program or system.
*/

/*
In the Go programming language, `nil` is a predeclared identifier representing the zero value for a number of reference types, including:

- Pointer types
- Function types
- Interface types
- Map types
- Slice types
- Channel types

For these types, `nil` effectively represents a "zero value".

For example:

- When `nil` is used with pointer types, it means the pointer is not pointing to any memory location (uninitialized pointer).
- When `nil` is used with function types, it signifies that the function variable is not assigned to any function.
- When `nil` is used with interface types, it means the interface has no concrete type and value associated with it.
- When `nil` is used with map types, it represents an uninitialized map that can't be added to until it is initialized with `make`.
- When `nil` is used with slice types, it represents a slice with length and capacity 0.
- When `nil` is used with channel types, it represents an uninitialized channel.

It is important to note that `nil` is not the same as zero for numeric types (`0` for integers, `0.0` for floats), nor the same as false for boolean types, nor the same as the empty string (`""`) for string types.

// regarding `nil` pointer

An "uninitialized pointer" in a programming language like Go refers to a pointer variable that has been declared but hasn't been given a specific memory address to point to yet. By default, in Go, this pointer will have the zero value `nil`. It effectively means it's not pointing to any memory location.

Here's an example:

```go
var p *int // p is an uninitialized pointer, its value is nil
```

An uninitialized pointer like this is usually not very useful until it's set to point to an actual variable or memory location. It's important to note that if you attempt to dereference (access the value at the pointer's address) an uninitialized pointer (a nil pointer), Go will throw a runtime error, specifically a panic, indicating a `nil pointer dereference`.

To properly initialize a pointer, you can do something like:

```go
var x int = 5
p := &x // p now points to x
```

In the above code, `p` is a pointer to `x` and it's initialized with the memory address of `x`. You can now safely dereference `p` with `*p` to get the value of `x`.

It's good practice to ensure that your pointers are always properly initialized before use, to avoid runtime errors.

// regarding 'nil` function

In Go, just like other types, you can have variables of function types. These are variables that you can assign a function to.

Let's illustrate this with an example:

```go
package main

import "fmt"

// function type declaration
type MyFuncType func(a int, b int) int

// function that matches the type
func add(a int, b int) int {
    return a + b
}

func main() {
    var myFunc MyFuncType // declare a function variable
    fmt.Println(myFunc)   // <nil>

    myFunc = add          // assign the add function to the variable
    fmt.Println(myFunc(2, 3))  // 5
}
```

In this example, `MyFuncType` is a type that represents functions that take two integers as parameters and return an integer. `myFunc` is a variable of type `MyFuncType`. At the start, `myFunc` isn't assigned to any function, so its value is `nil`.

When the phrase "function variable is not assigned to any function" is used, it refers to the state of the function variable before it has been assigned a function. In the example, this is the state of `myFunc` immediately after it has been declared but before it has been assigned the `add` function.

It's important to note that if you try to call a function variable that hasn't been assigned to any function (a nil function), Go will throw a runtime error, specifically a panic, indicating a `nil pointer dereference`. It's good practice to ensure that your function variables are always properly assigned before use.


*/

/*
In Go, the term "concrete type" is often used to distinguish from "interface types". A concrete type is a type that provides a specific implementation of a method or set of methods, as opposed to an interface type, which only declares method signatures but doesn't provide implementations.

Concrete types in Go include:

1. Basic types (like `int`, `float64`, `bool`, `string`)

2. Composite types, which are combinations of other types. They include:

   - Array types (like `[3]int`, `[4]string`)
   - Struct types (like `struct {Name string; Age int}`)
   - Function types (like `func(int, int) int`)
   - Pointer types (like `*int`, `*MyStruct`)
   - Slice types (like `[]int`, `[]string`)
   - Map types (like `map[string]int`, `map[int]bool`)
   - Channel types (like `chan int`, `chan string`)

A concrete type may satisfy one or more interfaces by implementing all the methods required by those interfaces. For example, if we have an interface:

```go
type Shape interface {
    Area() float64
}
```

And we have a concrete type `Circle`:

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

Here, `Circle` is a concrete type because it provides a specific implementation of the `Area` method, which makes it satisfy the `Shape` interface. We could assign an instance of `Circle` to a variable of type `Shape`:

```go
var s Shape
s = Circle{Radius: 5}
```

In this case, the concrete type of `s` is `Circle`, and it holds a `Circle` value.
*/

/*
In Go, a string is a two-word data structure that has two components:

1. Data: This is a pointer that points to an array of bytes in memory. These bytes represent the actual content of the string in a sequence of bytes. This is where the actual characters that make up the string are stored.

2. Length: This is an integer that represents the length of the string. This helps to know the size of the string and to prevent overrunning the end of the memory block when the string is being processed.

This two-word data structure is beneficial because it makes operations on strings, such as slicing, very efficient. Since the string data is immutable (it can't be changed once it's created), slicing doesn't involve copying the string data, but creating a new structure with a pointer that points to a different part of the data and a different length. This makes string operations in Go very memory efficient and fast.

This design also allows strings to safely share the underlying data arrays, even amongst different goroutines, as the data is immutable and cannot be changed once created. This makes string handling in Go concurrent-safe.

Remember that strings in Go are also UTF-8 encoded by default, which allows them to handle a wide range of characters from various languages, but this also means that the number of characters in a string might not be the same as the length of the string, because some characters might require more than one byte.
*/

/*
The `make` and `new` functions in Go language serve different purposes and are used with different types.

`new` function:
- `new` is a built-in function that allocates memory, but does not initialize the memory, it only zeros it.
- `new` is used with any type and returns a pointer to a newly allocated, zero value of the type passed as argument.
- Syntax: `new(T)` where T is any type. The result is `*T` pointing to a zero value of type `T`.

Example usage of `new`:

```go
ptr := new(int)
fmt.Println(*ptr)  // prints "0", because ptr points to a zero value for type int
```

`make` function:
- `make` is a built-in function that allocates and initializes memory for slice, map, and channel types. It returns an initialized (not zeroed) value of type T (not *T).
- Its primary use is for initializing slices, maps, and channels, which are data structures that are implemented using references to runtime types.
- Syntax: `make(T, args)` where T is a slice, map, or channel type.

Example usage of `make`:

```go
slice := make([]int, 5)  // creates a slice of integers of length 5
```

In summary, use `make` when you need to initialize the data structure, like slices, maps, and channels, and `new` when you just want to allocate but not initialize memory for a type.
*/

/*
RANDOM NOTES


 A good general guideline is that when we're working with built-in types, when the data were working with is either a numeric, a string, or a bool what we're going to want to use is our value semantics to move that data around the program. That's going to be our guideline. That means everybody gets their own copy of the int; everybody gets their own copy of the string; everybody gets their own copy of the bool. We're not going to share this data across program boundaries. In fact, this data has been designed to be copied, and anytime we can use those value semantics, we reduce allocations on the heap. That's going to help us lower latency costs." - Bill Kennedy

"Value semantics means that every piece of code is operating on its own copy of the data. We get benefits of that data isolation around mutation; we get data locality, which is going to help us later on with some performance around multi-threaded software. It's just cleaner when we can use those value semantics, but it comes with a cost, nothing is free, and that cost is inefficiency. And then we have our pointer semantics which means that we're going to only keep one copy of the data and it's going to be shared. And there's a lot of efficiency with that, but at the same time we have to worry about the side effects; we've got to worry about mutation a lot more." Bill Kennedy

"Just like our built-in types, we're going to be using value semantics for our reference types to move this data around our program. Out of the three classes of types we have in Go -- built-in, reference, and user-defined structs -- the reference types and built-in types use value semantics to move their data around. So once again, I don't want to see pointer semantics with built-in types or reference types. There are some exceptions." Bill Kennedy

"You have to be careful with reference types because even though we're going to use value semantics to move our data around, when you're reading and writing with reference types, you're also using pointer semantics. There's a duality to reference types. As reference types cross program boundaries, they don't pollute the heap." Bill Kennedy

"What makes the reference types interesting, is that when any of the reference types are set to their zero value, they're considered to be nil." Bill Kennedy

"You have to remember that when working with our reference types, like a slice, we might be using VALUE semantics to move it around our program so we're not polluting the heap with these little slice values, but any time we're reading and writing, we're using those pointer semantics." Bill Kennedy

Always watch for append. Any time a backing array can be replaced, double-check there are no errors or side-effects around this.
*/

/*
BENCHMARK

package main

import "testing"

func BenchmarkValSemantics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qty := 10
		vs := make([]string, 0, qty)
		valSemantics(vs, qty)
	}
}

func BenchmarkPtrSemantics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qty := 10
		ps := make([]string, qty)
		ptrSemantics(ps, qty)
	}

}

*/

/*

	declare, assign, initialize, and allocate refer to different aspects of working with variables and memory.

	- Declare: Create a variable without assigning it a value.
	- Assign: Set a value to a previously declared variable.
	- Initialize: Declare a variable and assign it a value at the same time.
	- Allocate: Reserve memory for more complex types or for dynamic memory needs.

	1. Declare: Declaring a variable in Go means stating its type and name, but not necessarily giving it a value. This merely reserves a memory spot for the variable. A variable can be declared using the `var` keyword. For example, `var x int` declares an integer variable named `x`.

	2. Assign: Assigning a variable means giving a declared variable a value. This is typically done using the `=` operator. For instance, if you have declared `var x int` earlier, you could assign it a value with `x = 10`.

	3. Initialize: Initialization refers to the process of declaring a variable and assigning it a value at the same time. For instance, `var x int = 10` or `x := 10` (using short declaration) both declare and assign a value to `x`, thus initializing it. This is the typical way variables are created in Go.

	4. Allocate: Allocation refers to reserving a specific amount of memory during the execution of a program, usually for more complex types such as slices or structs, or when you need to create a variable dynamically at runtime. In Go, you typically use built-in functions like `make()` or `new()` to allocate memory. For example, `x := make([]int, 10)` allocates a slice of integers with a length of 10. `new()` is used to allocate memory and returns a pointer to it, but unlike `make()`, it doesn't initialize the memory.
*/

/*
In Go, semantics refers to how types behave during assignments and function calls. In particular, the semantics of a type is about whether its values are directly manipulated or whether they are indirectly manipulated. The distinction between value semantics and pointer semantics is crucial in Go, and understanding it can help you write more efficient and correct programs.

1. Value Semantics: When you pass a value with value semantics, you're creating a copy of that value. Any changes you make to the copy won't affect the original value. All the basic types in Go (like int, float64, bool, etc.) use value semantics. This is also true for composite types like structs and arrays.

    Here's an example of value semantics:

    ```go
    type S struct {
        a int
    }

    func change(s S) {
        s.a = 5
    }

    func main() {
        var s S
        change(s)
        fmt.Println(s.a)  // Will print 0, not 5
    }
    ```
    In the code above, `change` function receives a copy of `s`. When `s.a` is modified inside `change`, it only changes the copy, not the original `s` in the main function.

2. Pointer Semantics: On the other hand, when you pass a pointer, you're passing a reference to the value, not the value itself. So any changes you make to the value the pointer references will affect the original value.

    Here's an example of pointer semantics:

    ```go
    type S struct {
        a int
    }

    func change(s *S) {
        s.a = 5
    }

    func main() {
        var s S
        change(&s)
        fmt.Println(s.a)  // Will print 5
    }
    ```
    In this code, the `change` function receives a pointer to `s`. When `s.a` is modified inside `change`, it changes the original `s` in the main function, because `s` is a pointer to the original struct.

To summarize, if you want to modify the original value or if the data being passed is large (like a big struct or array), it might be better to use pointer semantics. However, if you want to ensure the function doesn't modify the original value or if you're working with simple small types (like int or bool), you might want to use value semantics.

*/

/*
BUILTIN TYPES
Built-in types are those that are built into the language itself, and they don't require any additional libraries to be imported. They are ready to use straight out of the box. The built-in types in Go are:

Boolean types:
- `bool`: A boolean type that can have the values `true` or `false`.

Numeric types:
- `uint8`, `uint16`, `uint32`, `uint64`: Unsigned (non-negative) integers that can be 8, 16, 32, or 64 bits long, respectively.
- `int8`, `int16`, `int32`, `int64`: Signed integers that can be 8, 16, 32, or 64 bits long.
- `uint`, `int`, `uintptr`: An unsigned integer, an integer, and a pointer-sized unsigned integer whose size depends on the type of architecture you're compiling for (32 or 64 bits).
- `float32`, `float64`: A floating point number that can be 32 or 64 bits long.
- `complex64`, `complex128`: A complex number with float32 or float64 real and imaginary parts.
- `byte`: An alias for `uint8`.
- `rune`: An alias for `int32`. It represents a Unicode code point.

String types:
- `string`: A sequence of bytes. It's often used to hold text.

Error type:
- `error`: A built-in interface type used for representing and manipulating errors.

REFERENCE TYPES

- Slices: A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment). If you pass a slice to a function, it will be able to modify the underlying array.

- Maps: A map is an unordered collection of key-value pairs. Maps in Go are implemented as hash tables and they are references to the underlying data structure. If you pass a map to a function, the function can change the map's elements.

- Channels: Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine. Like slices and maps, channels are reference types.

- Pointers: A pointer holds the memory address of a value. Pointers in Go support the "reference semantics" behavior.

- Interfaces: An interface type is defined by a set of methods. A value of interface type can hold any value that implements those methods. Interface values are a combination of a value and a concrete type. They have reference semantics.

- Functions: Functions in Go are first-class citizens and can be assigned to variables or passed around just like any other types. Functions have reference semantics as well.

While these types have reference-like behavior, it's worth noting that Go does not have "classes" or "objects" in the traditional sense, and therefore does not have reference types in the way that languages with a class-based object-oriented model do. Instead, Go has user-defined types and supports composition and interfaces to provide flexible ways of structuring your code.

STRUCTS
Structs: Structs are a way to group together variables of different types.

------------------------
other ways we can categorize types
------------------------

COMPOSITE TYPES / AGGREGATE TYPES:
array: A numbered sequence of elements of a specific length.

slice: A dynamically-sized, flexible view into the elements of an array.

struct: A composite data type that groups together zero or more values of different types.

pointer: Holds the memory address of a value.

function: A sub-routine which is a group of statements that together perform a task.

interface: Defines a contract for certain behavior (like methods) that the types should have.

map: An unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.

channel: A conduit through which you can send and receive values with the channel operator, <-.


USER DEFINED TYPES:
User-defined types are types that are defined by the programmer, not the language. They allow you to create complex data structures by COMPOSING together existing types, be they built-in or other user-defined types.

Here are some examples of user-defined types in Go:

- STRUCTS: Structs are a way to group together variables of different types. You could define a struct to represent a person, for example, with a name and an age:

type Person struct {
    Name string
    Age  int
}

- INTERFACES: Interfaces define a contract of methods that a type should implement. They are central to Go's type system and enable polymorphic behavior.

type Shape interface {
    Area() float64
}

- ALIASES: You can also define a new type as an alias of an existing one, which can be useful for improving code readability or for encapsulating functionality. For example:

type Age int

In this case, `Age` is a new type, but it has the same underlying type as `int`. You can define methods on these user-defined types.

- FUNCTIONS: In Go, functions are first-class citizens, meaning they can be defined as types and then used as arguments or return values in other functions. For example:

type BinaryOperation func(a int, b int) int

In this case, `BinaryOperation` is a function type that takes two `int` parameters and returns an `int`.

Note that while Go does allow you to define new types in these ways, it does not support classes or inheritance in the way that some other languages do. Instead, it encourages composition and the use of interfaces to define behavior.

*/

/*

Below are common metric units of time listed from largest to smallest, along with their definitions:

1. Kilosecond (ks): 1 kilosecond = 1,000 seconds = ~16.67 minutes.

2. Second (s): The second is the base unit of time in the International System of Units (SI).

3. Decisecond (ds): 1 decisecond = 0.1 seconds.

4. Centisecond (cs): 1 centisecond = 0.01 seconds.

5. Millisecond (ms): 1 millisecond = 0.001 seconds. This unit of time is often used in computing to measure response time.

6. Microsecond (μs): 1 microsecond = 0.000001 seconds = 1×10^-6 seconds. This unit of time is often used in scientific research and in telecommunications.

7. Nanosecond (ns): 1 nanosecond = 0.000000001 seconds = 1×10^-9 seconds. This unit of time is frequently used in computing and in the timing of electronic circuits.

8. Picosecond (ps): 1 picosecond = 0.000000000001 seconds = 1×10^-12 seconds. This unit of time is used in ultrafast science, like in laser physics and in the measurement of molecular dynamics.

9. Femtosecond (fs): 1 femtosecond = 0.000000000000001 seconds = 1×10^-15 seconds. This is a unit often used in research, like in ultrafast laser studies, where events can occur in these small intervals of time.

10. Attosecond (as): 1 attosecond = 0.000000000000000001 seconds = 1×10^-18 seconds. This is among the smallest units of time currently measurable, and has applications in research involving ultrafast processes like electron dynamics inside atoms.

11. Zeptosecond (zs): 1 zeptosecond = 1×10^-21 seconds.

12. Yoctosecond (ys): 1 yoctosecond = 1×10^-24 seconds. This is one of the smallest theoretically measurable units of time.

These measurements are part of the International System of Units (SI), and they all relate to the base unit, which is the second.

*/

/*
	GHz stands for gigahertz. It is a unit of frequency equal to one billion (1,000,000,000) cycles per second. In computing, it's most commonly used to measure the speed (clock rate) of a computer's central processing unit (CPU). A CPU with a clock speed of 2.5 GHz, for instance, performs 2.5 billion cycles per second. However, it's important to note that a higher clock rate doesn't always mean better performance, as it also depends on other factors such as the processor's architecture and the efficiency of the software being used.

	In the world of computer processors, a cycle is a single tick of a computer's clock, which helps synchronize the operations of the different parts of the computer. Theoretically, one might think that in one cycle, one instruction is completed. But, the reality is more complex.

	The relationship between clock cycles and instructions is defined by two metrics: IPC (Instructions Per Cycle) and CPI (Cycles Per Instruction).

	1. IPC (Instructions Per Cycle): This is the average number of instructions that can be executed for each clock cycle. For some modern CPUs with advanced features like superscalar architecture, pipelining, or simultaneous multithreading, IPC can be greater than one. That means they can complete more than one instruction per clock cycle.

	2. CPI (Cycles Per Instruction): This is the inverse of IPC. It is the average number of clock cycles needed to execute one instruction. For simple processors, this might be just one. For more complex CPUs, this could be a fraction (when IPC > 1) or greater than one (when multiple cycles are needed to execute one instruction due to complex instructions or CPU architectures).

	So, to answer your question: does one operation occur for every cycle? It depends on the CPU's architecture and the type of instructions it's executing. With advanced modern processors, more than one operation can occur per cycle, but there can also be cases where an operation requires more than one cycle to complete.
*/

/*
In Go, arrays and slices are both sequence types, but they have important differences that relate to their size, flexibility, and use in Go's idiomatic coding patterns.

Here's a basic comparison:

1. Arrays

    - In Go, an array is a sequence of elements of a specific length. The length of an array is part of its type, so `[5]int` and `[10]int` are different types.

    - Arrays in Go are values. When you assign one array to another, the entire array is copied. If you pass an array to a function, it will receive a copy of the array, not a pointer to it.

    - Arrays have a fixed size. Once an array is declared, its size cannot be changed.

2. Slices

    - A slice is a flexible, dynamically-sized sequence of array elements. It's a reference type, which means when you assign a slice to another, both refer to the same underlying array.

    - Unlike arrays, slices are dynamically sized. A slice can grow or shrink as needed, up to the length of the underlying array.

    - Slices are much more commonly used in Go than arrays. The built-in `make` function can be used to create a slice of a particular length and capacity.

    - When you pass a slice to a function, it will receive a reference to the slice. If the function changes the elements of the slice, the changes are visible outside the function.

A key element to understand is the internal representation of slices in Go: a slice header, which is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment). So a slice can be thought of as a struct type with three fields:

```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```
This internal representation allows a slice to dynamically resize within the capacity of the underlying array.

RULES OF THUMB

Here are some general guidelines about using arrays and slices in Go:

1. Prefer slices over arrays: Slices are more common and flexible than arrays. Slices can grow and shrink as needed, which makes them ideal for many scenarios. If you don't know the size of your data set beforehand or if it is likely to change, you should use a slice, not an array.

2. Use arrays when size is known and constant: If the number of elements is known and will not change, arrays could be a better choice. For example, you may use an array to represent a chessboard, which has a fixed size of 8x8.

3. Use slices for function arguments and returns: When passing a collection to a function or returning a collection from a function, use slices, not arrays. This gives you more flexibility and also can be more efficient, because only a slice header (a small amount of data) is being copied, not potentially large amounts of data.

4. Preallocate slice capacity when possible: If you know a slice is going to grow, it's often a good idea to preallocate the capacity using the `make` function. This can be more efficient because it saves the Go runtime from having to repeatedly resize the underlying array.

5. Be aware of the 'gotcha' with slice references: Because slices are references to an underlying array, changes made to one slice can appear in another slice if they share the same array. This can be a source of bugs. One way to prevent this is to use the `copy` function to make sure that your slices don't overlap.

6. Use `range` to iterate over slices and arrays: The `range` keyword in Go allows you to easily iterate over arrays and slices. It's idiomatic and avoids off-by-one errors that can occur with traditional for loops.

7. Consider the implications for garbage collection: When you're holding onto a slice, the entire underlying array cannot be garbage collected, even if elements outside the slice are no longer needed. If you're done with a large slice, consider setting it to nil so that the garbage collector can reclaim the memory.

8. Use arrays when the size is fixed and known in advance. Since arrays have a fixed length that is part of their type, they can be useful when you know that a sequence of elements will always have a certain size. For example, you might use an array to hold the days of the week, or the months of the year.

9. Use slices for dynamically-sized collections. If you don't know how many elements there will be, or if the number of elements can change, then you should use a slice. Most of the time, you'll be using slices.

10. When performance matters, consider using arrays. Since arrays in Go are values, using them can avoid the heap allocation that typically comes with a slice. If you're writing performance-critical code, it might be more efficient to use an array. However, remember that passing large arrays as function parameters can be costly because the entire array will be copied.

11. When using slices, be aware of capacity and the underlying array. Since a slice is just a view into an array, you can sometimes end up with unexpected behavior if you're not careful. For example, if you create a new slice by slicing an existing one, you can end up modifying the original slice.

Remember, these are general guidelines, and there can be exceptions depending on the specific scenario or problem you're trying to solve. Always choose the data structure that most closely matches your specific use case and makes your code as clear and understandable as possible.

*/

/*
type ByteSize int

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ZB
	// YB
)

func main() {
	fmt.Printf("KB %d \t\t %b\n", KB, KB)
	fmt.Printf("MB %d \t\t %b\n", MB, MB)
	fmt.Printf("GB %d \t\t %b\n", GB, GB)
	fmt.Printf("TB %d \t %b\n", TB, TB)
	fmt.Printf("PB %d \t %b\n", PB, PB)
	fmt.Printf("EB %d \t %b\n", EB, EB)
}
*/

/*
In Go (often referred to as Golang), the keyword `const` is used to declare a constant. A constant is a simple, immutable value that remains the same throughout the life of a program. This contrasts with variables, which can have their values changed.

Here's an example of a constant declaration:

```go
const Pi = 3.14159
```

Now, in terms of "constants of a kind" and "constants of a type", it helps to understand the Go specification.

The specification refers to the concepts of "types" and "kinds". In Go, every type has a kind. For instance, the kind of the type `int` is `int`, the kind of the type `*int` is `pointer`, the kind of the type `[]int` is `slice`, and so on. There are several predeclared types such as `bool`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, `float32`, `float64`, `complex64`, `complex128`, `string`, etc., each of these types has its own kind.

In Go, constants can be untyped or typed.

Untyped Constants

Untyped constants in Go are just that: constants that don't have a type yet. These constants are known as "constants of a kind". They have a default type that is used when a type is needed, such as when assigning to a variable, but they can be used wherever any type of constant is allowed as long as they are representable in that type.

Here is an example:

```go
const MyConstant = 123  // MyConstant is a constant of kind int
```

In this example, `MyConstant` doesn't have a specified type, but it has a kind: `int`.

Typed Constants

When you declare a constant with a specific type, you create a "constant of a type". This means that the constant not only has a value, but also has a specific type associated with it.

Here is an example:

```go
const MyTypedConstant int = 123  // MyTypedConstant is a constant of type int
```

In this example, `MyTypedConstant` is a constant of type `int`. Unlike the untyped constant, this constant cannot be used wherever any type of constant is allowed. It can only be used where an `int` is allowed.

In summary, a "constant of a kind" (or an untyped constant) in Go is a constant that hasn't been given a specific type but can be used with any compatible type, while a "constant of a type" (or a typed constant) in Go is a constant that has been given a specific type and can only be used as that type.


The type of an untyped constant, also known as a "constant of a kind", is determined at compile time in Go.

Untyped constants in Go are a bit special. They have a default type associated with them, but they can "become" other types if it's clear from the context what type they should be. This is why they are referred to as untyped constants, they're more flexible than typed constants.

Here's an example:

```go
const MyConstant = 123 // untyped constant

var i int = MyConstant // MyConstant becomes an int here
var f float64 = MyConstant // MyConstant becomes a float64 here
```

In this example, `MyConstant` is an untyped constant with a value of `123`. When we assign `MyConstant` to `i`, which is of type `int`, `MyConstant` "becomes" an `int`. Similarly, when we assign `MyConstant` to `f`, which is of type `float64`, `MyConstant` "becomes" a `float64`.

All of this happens at compile time, not at runtime. The Go compiler determines the type of the untyped constant from the context in which it's used.

However, this flexibility has its limits. An untyped constant can only "become" a type if it's a valid representation for that type. For example, the untyped constant `123` can "become" an `int` or a `float64`, but it can't "become" a `string` because `123` isn't a valid `string`. If you try to do something like this:

```go
const MyConstant = 123 // untyped constant

var s string = MyConstant // this will not compile
```

The Go compiler will throw an error because it can't use `123` as a `string`. This determination of type compatibility is also done at compile time.
*/

/*
In the context of Go programming language and concurrent computing, synchronization and orchestration are two distinct concepts, although they might be intertwined in practice.

1. Synchronization is a concept that deals with the order and timing of operations. In Go, it's about making sure that different goroutines (lightweight threads) interact in a safe and predictable way. There are multiple synchronization mechanisms provided by the Go language, such as:

   - Mutexes (`sync.Mutex`, `sync.RWMutex`): This mechanism is used to ensure that only one goroutine accesses a critical section of code at a time to prevent race conditions.

   - Channels: These are used to synchronize goroutines through communication. By default, channels block send/receive operations until both the sender and receiver are ready, providing a natural way to synchronize execution.

   - WaitGroups (`sync.WaitGroup`): These are used to block until a group of goroutines have all completed their execution.

   - Condition variables (`sync.Cond`): These are used to signal changes to goroutines that are waiting for a specific condition to become true.

2. Orchestration on the other hand, involves coordinating the flow of control across different components (goroutines, in the case of Go). The concept applies to managing the lifecycle, interaction, and scheduling of goroutines. While synchronization may be used as a part of orchestration, orchestration itself can involve other aspects like determining which goroutines to start, which to stop, how to handle errors, etc.

   - The primary orchestration tool in Go is the goroutine itself. By structuring your code in a way that different components run in separate goroutines, you are effectively orchestrating your concurrent program.

   - Channels are also used in orchestration. For example, a common pattern in Go is to have a `done` channel that is closed when it's time for a goroutine to stop executing. This is a way of orchestrating the goroutines to stop at a particular time.

In summary, while synchronization is about ensuring goroutines can work together safely and predictably, orchestration is about managing how these goroutines work together to achieve the overall task.
*/
