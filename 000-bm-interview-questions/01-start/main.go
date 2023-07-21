package main

import "fmt"

func main() {
	// Create a slice to store the questions
	var questions []string

	// Add questions to the slice
	questions = append(questions, "Put these in the correct order: centisecond, second, nanosecond, millisecond, decisecond, microsecond")
	questions = append(questions, "What is the difference between make([]int, 10), make([]int, 0, 10), make([]int, 10, 10)")
	questions = append(questions, "Explain these types in Go: builtin, reference, user defined")
	questions = append(questions, "Explain value semantics and pointer semantics. What are rules-of-thumb for using one versus the other?")
	questions = append(questions, "What types can a map use as a key in the Go programming language?")
	questions = append(questions, "Explain how a linked list works.")
	questions = append(questions, "What does it mean to write idiomatic Go code?")
	questions = append(questions, "Who is Rob Pike?")
	questions = append(questions, "Who is Ken Thompson?")
	questions = append(questions, "Who is Robert Griesemer")
	questions = append(questions, "What is the short-declaration operator and when do you use it?")
	questions = append(questions, "What is the var keyword used for and when do you use it?")
	questions = append(questions, "How are arrays and slices different in Go? What are rules-of-thumb around arrays and slices in Go?")
	questions = append(questions, "What are pointers?")
	questions = append(questions, "How does a computer work?")
	questions = append(questions, "What is iota?")
	questions = append(questions, "What is a Goroutine?")
	questions = append(questions, "What is the difference between concurrency and parallelism?")
	questions = append(questions, "What is a compiler?")
	questions = append(questions, "What is the stack?")
	questions = append(questions, "What is the heap?")
	questions = append(questions, "What is escape analysis?")
	questions = append(questions, "Which do you choose: performance or readability?")
	questions = append(questions, "Which do you choose: brevity or readability?")
	questions = append(questions, "What is a garbage collector?")
	questions = append(questions, "Is Go an Object Oriented language?")
	questions = append(questions, "Is Go a Data Oriented language?")
	questions = append(questions, "How does Go handle error handling compared to other languages?")
	questions = append(questions, "Why was Go created, and who created it?")
	questions = append(questions, "Can you explain what is meant by 'Go is strongly typed?'")
	questions = append(questions, "Can you give an example of a problem you've solved using Go? What challenges did you encounter, and how did you overcome them?")
	questions = append(questions, "Can you describe a time when you used channels and goroutines to solve a concurrency problem?")
	questions = append(questions, "Have you worked on a large-scale project using Go? If so, how did Go's features contribute to the project's success?")
	questions = append(questions, "What tools do you typically use for testing Go code?")
	questions = append(questions, "How would you set up a CI/CD pipeline for a Go project?")
	questions = append(questions, "Can you explain how package management works in Go?")
	questions = append(questions, `What's problematic with this code:

	func slice() {
		xi := make([]int, 10)
		for i := 0; i < 10; i++ {
			xi1 = append(xi1, i)
		}
		fmt.Println(xi1)
	}		
	`)
	questions = append(questions, "What is your IDE for Go, and why?")
	questions = append(questions, "Can you argue in favor of your perspective, and accept when your perspective isn't chosen?")
	questions = append(questions, "What do you do to stay current with programming?")
	questions = append(questions, `Do you recognize any of these, and if so, what are they and have you worked with them:
	cobra
	viper
	opentelemetry
	jaeger
	docker
	chi
	zshell
	`)
	questions = append(questions, "What is a multiplexer?")
	questions = append(questions, "What is the difference between TCP and UDP?")
	questions = append(questions, "Who is Caleb Doxsey?")
	questions = append(questions, "Tell me about containers and container orchestration.")
	questions = append(questions, "What made Docker innovative?")
	questions = append(questions, "Tell us about your soft-skills.")
	questions = append(questions, "Tell us about a time you failed, and what you learned from it.")
	questions = append(questions, "Tell us about a time you succeeded, and why you succeeded.")
	questions = append(questions, "What makes you a valuable member of a team?")
	questions = append(questions, "Explain internal, external, and data latencies in Go.")
	questions = append(questions, "What is the mascot of the Go programming language?")
	questions = append(questions, "How do we do benchmarking in Go?")
	questions = append(questions, "What are method sets, and how do you use them?")
	questions = append(questions, "What is an interface?")
	questions = append(questions, "What is a type set?")
	questions = append(questions, "What is a concrete type?")
	questions = append(questions, "Explain generics.")
	questions = append(questions, "What's the difference between declare, assign, and initialize?")
	questions = append(questions, "Explain allocation and initialization.")
	questions = append(questions, "What is nil?")
	questions = append(questions, "How might you use a nil channel?")
	questions = append(questions, "What is the difference between switch and select?")
	questions = append(questions, "What is the comma ok idiom?")
	questions = append(questions, `Fix this code:
	package main

	import "fmt"
	
	func main() {
		var m map[string]int
		m["b"] = 42
		fmt.Println(m)
	}
	
	`)
	questions = append(questions, `Should you use buffered channels? Why or why not?`)
	questions = append(questions, `What's a cache line, what's it's common size, and why would you think about this in relation to the Go programming language?`)
	questions = append(questions, `Somebody says that 'a string is a two word data structure' - what does this mean?`)
	questions = append(questions, `Somebody says that 'a slice is a three word data structure' - what does this mean?`)
	questions = append(questions, `When dealing with computer architecture, what does the 'word size' mean?`)
	questions = append(questions, `Why does this code not print 'biking'
	
	var sports [5]string
	sports[0] = "ski"
	sports[1] = "surf"
	sports[2] = "swim"
	sports[3] = "sail"
	sports[4] = "sumo wrestling"

	for i, v := range sports {
		sports[i] = "biking"
		fmt.Println(v)
	}
	
	`)
	questions = append(questions, `What is the difference between make and new?`)
	questions = append(questions, `What is the blank identifier?`)
	questions = append(questions, `Fix this code, then benchmark the two functions:

	func main() {
		qty := 10
		var vs []string
		ps := make([]string, 0, qty)

		vs = valSemantics(vs, qty)
		fmt.Println(len(vs), cap(vs))

		ptrSemantics(ps, qty)
		fmt.Println(len(ps), cap(ps))
	}

	func valSemantics(ss []string, n int) []string {
		for i := 0; i < n; i++ {
			s := fmt.Sprintf("item %d", i)
			ss = append(ss, s)
		}
		return ss
	}

	func ptrSemantics(ss []string, n int) {
		for i := 0; i < n; i++ {
			ss[i] = fmt.Sprintf("item %d", i)
		}
	}

	
	
	`)
	questions = append(questions, `What will this code print:
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		xs := sports[1:3:3]
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
	questions = append(questions, `Fix this code so that when xs[0] is changed this doesn't change 'sports':
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		xs := sports[1:3:3]
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

	questions = append(questions, `Will this print the same for both print statements:
	
	people := make([]person, 2)

	p1 := people[1]
	p1.age++
	fmt.Println(p1.age)
	fmt.Println(people[1].age)
	
	`)
	questions = append(questions, `What will the last print statement print, and why?
	
	people := make([]person, 2)

	p1 := &people[1]
	p1.age++
	fmt.Println(people[1].age)

	// Add a new person
	people = append(people, person{})

	// Add another year for p1
	p1.age++
	fmt.Println(people[1].age)

	`)
	questions = append(questions, `Fix this code:
	
	people := make([]person, 2)

	p1 := &people[1]
	p1.age++
	fmt.Println(people[1].age)

	// Add a new person
	people = append(people, person{})

	// Add another year for p1
	p1.age++
	fmt.Println(people[1].age)

	`)
	questions = append(questions, `Teach us something most people don't know about Go.`)
	questions = append(questions, `Tell us about bytes, code points, and characters in relation to strings and UTF-8.`)
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
	xi3 := make([]int, 10, 10)

	for i := 0; i < 10; i++ {
		xi1 = append(xi1, i)
		xi2 = append(xi2, i)
		xi3 = append(xi3, i)
	}
	fmt.Println(xi1)
	fmt.Println(xi2)
	fmt.Println(xi3)

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
