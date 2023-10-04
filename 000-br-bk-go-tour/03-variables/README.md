# "Go is easy to learn and hard to master." ~ Andrew Brudnak

# Video
[Here is the video #5 on youtube](https://www.youtube.com/watch?v=AmiomWHD6F8)
[Here is the video #6 on youtube](https://youtu.be/wpS3T83IRfc)

# Takeaways
1. Type tells us (1) size and (2) representation
2. Three classes of type: builtin, reference, struct (user defined composite)
3. initialize = declare & assign
4. Conversion over casting
5. Every line of code either reads, writes, or allocates memory
6. integrity: accurate, consistent, efficient
7. We have a language we use to talk about the language
8. Remember, nobody can stop you from learning.

# Table of Contents

1. [Notes on Variables](#notes-on-variables)
2. [Types Memory](#types-memory)
3. [Declare Assign Initialize and Zero Value](#declare-assign-initialize-and-zero-value)
4. [Builtin Reference and Struct types](#builtin-reference-and-struct-types)
5. [String is a builtin type](#string-is-a-builtin-type)
6. [Zero Value Details](#zero-value-details)
7. [Value Semantics and Pointer Semantics](#value-semantics-and-pointer-semantics)
8. [Conversion vs Casting](#conversion-vs-casting)
9. [Advantages of Statically Typed Languages](#advantages-of-statically-typed-languages)
10. [Go](#go)
11. [Code review](#code-review)
12. [Coupons for Go courses](#coupons-for-go-courses)

# Notes on Variables

### Type
- type tells us two things
    - the amount of memory we’re going to be reading and writing; amount of memory used
    - what that memory represents; what is stored there
- three categories of types in Go
    - builtin: numeric, string, bool, array
    - reference: slices, maps, channels, pointers
    - struct
- initialize = declaring and assigning
- var
    - use for zero value
- short declaration operator
    - inside code block
- conversion over casting
    - conversion has more safety than casting - it is A NEW ALLOCATION of memory - conversion allocates new memory and copies values in
    - "Instead of telling the compiler to map a set of bytes to a different representation, the bytes need to be copied to a new memory location for the new representation." Bill Kennedy
    
# Types Memory
`bool` uses 1 byte of memory
`int` uses 1-8 bytes of memory
`string` uses 2 bytes of memory
`slice` uses 3 bytes of memory

# Declare Assign Initialize and Zero Value
In the Go programming language, the terms "declare," "assign," "initialize," and "zero value" have specific meanings that are important for managing variables and their states throughout the lifecycle of a program. Let's break down each term:

### Declare

Declaration is the process of telling the Go compiler that you intend to use a variable by specifying its name and type. This does not assign any value to the variable; it merely allocates space for it. For example:

```go
var x int  // Declaration without initialization
```

### Assign

Assignment is the act of setting a value to a variable that has already been declared. The assignment operator `=` is used for this purpose. You can assign a value to a variable either at the time of declaration or later in the program.

```go
x = 42  // Assignment
```

### Initialize

Initialization is the act of declaring a variable and assigning it a value simultaneously. This can be done in a variety of ways in Go:

```go
var x int = 42  // Declaration with initialization
var y = 42  // Type inference
z := 42  // Short variable declaration and initialization
```

### Zero Value

In Go, variables that are declared but not initialized are automatically assigned a "zero value." The zero value is type-dependent:

- For numeric types, the zero value is `0`.
- For the boolean type, the zero value is `false`.
- For strings, the zero value is an empty string `""`.
- For pointers, the zero value is `nil`.
- For other composite types like arrays, slices, maps, and structs, the zero value is each element or field set to its respective zero value.

Example:

```go
var a int  // a is 0
var b bool  // b is false
var c string  // c is ""
var d *int  // d is nil
```

### Summary - Declare, Assign, Initialize, Zero Value

- **Declare**: Telling the compiler about the variable name and type.
- **Assign**: Setting a value to a declared variable.
- **Initialize**: Declaring and assigning a value to a variable at the same time.
- **Zero Value**: The default value that variables take on when they are declared but not initialized.

Understanding these concepts is crucial for effectively managing variable states and resource allocation in Go programs.

# Builtin Reference and Struct types

In Go, types are categorized in various ways, with some of the most common categories being built-in types, reference types, and struct types. Understanding these different types is essential for effective Go programming. Here's a breakdown of each:

### Built-in Types

Built-in types are types that are part of the Go language itself and provide the basic building blocks for constructing more complex types. Some commonly used built-in types include:

- Numeric types (`int`, `float64`, etc.)
- Boolean type (`bool`)
- String type (`string`)

Examples:

```go
var i int = 10
var f float64 = 3.14
var b bool = true
var s string = "hello"
```

### Reference Types

Reference types are types that hold references to the underlying data. They can be thought of as "containers" that point to the actual value or a collection of values. Reference types in Go include:

- Slices
- Maps
- Channels
- Pointers

Note: While arrays in Go are not reference types (they are value types), slices, which are built on top of arrays, are reference types. This means that when you pass a slice to a function, modifications to the slice within the function affect the original slice.

Examples:

```go
// Slice
nums := []int{1, 2, 3}

// Map
mapping := map[string]int{
    "one": 1,
    "two": 2,
}

// Channel
ch := make(chan int)

// Pointer
var ptr *int
```

### Struct Types

Struct types are composite data types that group variables under a single name. These variables can be of different types, and each field in the struct is given a name. Structs are used for grouping data and are particularly useful when you want to create complex types that aggregate multiple pieces of information.

Example:

```go
type Person struct {
    Name string
    Age  int
    Address string
}

var p Person
p.Name = "John"
p.Age = 30
p.Address = "123 Main St"
```

### Summary - Built-in, Reference, & Struct types

- **Built-in Types**: Basic types that are part of the language, like `int`, `float64`, `bool`, and `string`.
- **Reference Types**: Types that hold references to the underlying data, like slices, maps, channels, and pointers.
- **Struct Types**: Composite types that allow you to group different types of variables together.

Each of these types has its own set of characteristics and best-use cases, and they can be combined to create more complex data structures. For instance, you can have a slice of structs or a map whose values are channels, and so on. Understanding the nuances between these types is crucial for writing idiomatic and efficient Go code.

# String is a builtin type

In Go, the `string` type is often described as being "like" a data structure containing a pointer to its underlying bytes and a length, similar to how slices work. Despite this, strings in Go are considered value types, not reference types, for several important reasons:

### Immutability

Strings in Go are immutable, meaning that once a string is created, its contents cannot be changed. This is different from slices, maps, and channels, which are mutable. When you "change" a string in Go, you're actually creating a new string with new underlying bytes. The original string remains unchanged.

```go
s := "hello"
s += ", world"  // This creates a new string and assigns it to `s`, the original string "hello" remains unchanged.
```

Because of their immutability, strings can be copied by value without concern for unintended modifications. This makes them behave more like value types in practice.

### Value Semantics

When you pass a string to a function or assign it to a new variable, the string's data is not copied. Instead, the new variable or function parameter gets its own copy of the pointer and length information, pointing to the same underlying bytes. However, due to immutability, this "copying" doesn't have the same implications as it would for a mutable reference type like a slice. Any "modification" to the string results in a new string, leaving the original data intact.

```go
func modifyString(s string) string {
    return s + " modified"
}

original := "original"
modified := modifyString(original)

// original is still "original", not affected by the function.
```

### Practical Implications

From a practical perspective, treating strings as value types simplifies many things:

1. It eliminates the need for defensive copying when passing strings between functions.
2. It makes strings safer to use in concurrent programming contexts, as they are immutable and thus cannot be changed unexpectedly by another goroutine.
3. It allows the compiler to optimize string usage in ways that might not be safe for mutable reference types.

In summary, even though strings in Go might seem similar to slices—having a pointer to underlying data and a length—they are immutable and exhibit value semantics, making them fundamentally different from reference types like slices, maps, and channels. This is why they are considered value types in Go.

# Zero Value Details

In Go, the concept of a zero value is significant for several reasons, some of which can help ensure the integrity of code:

### Initialization
In Go, variables declared without explicit initialization are set to their zero values. This makes sure that all variables are in a valid, known state when they are used, thus reducing the risk of unexpected behavior. For instance, the zero value for a pointer is `nil`, for numeric types, it's `0`, and for strings, it's an empty string `""`.

### Consistency
The zero value provides a consistent starting point for variable types. This can make the program easier to reason about and reduce the chance of errors caused by uninitialized variables.

### Simplicity and Readability
The zero value concept simplifies code. You don't have to explicitly initialize every variable if the zero value is appropriate for your use case. This can make the code more concise and readable.

### Error Handling
The zero value can serve as a signal for uninitialized or "not set" data. For example, if a function returns a pointer, it can return `nil` to indicate an error or a non-result, simplifying error handling.

### Reduced Bugs
Languages where uninitialized variables can have indeterminate values can introduce subtle bugs. In Go, the zero value eliminates this class of bugs, as variables are automatically initialized to a well-defined value.

### Easier to Test
Testing code often requires setting up known states. The zero value provides a predictable starting point for test conditions.

### Optimizations
Having a consistent and predictable zero value allows for some optimizations. For example, arrays and slices of zero values can be efficiently represented in memory.

### Type Safety
The zero value is always of the same type as the variable, contributing to Go's strong type safety. This makes it harder to inadvertently mix incompatible types, another benefit for code integrity.

So, while the concept of a zero value may seem trivial, it serves several important roles in making Go code robust, readable, and maintainable.

# Value Semantics and Pointer Semantics 

In the Go programming language, the terms "value semantics" and "pointer semantics" refer to how data is accessed and manipulated through variables, particularly in the context of function calls and assignments. Understanding these semantics helps you to write correct and efficient Go programs. Here's a breakdown of each:

### Value Semantics

With value semantics, data is passed or assigned by value, meaning that a copy of the data is created. Simple types like integers, floats, and booleans, as well as composite types like structs and arrays, exhibit value semantics by default.

When you pass a variable with value semantics to a function or assign it to another variable, you are actually creating a separate copy of that data. Changes made to the copy have no effect on the original data.

Here's an example using a struct:

```go
type Point struct {
    X, Y int
}

func modifyPoint(p Point) {
    p.X = 10
    p.Y = 20
}

func main() {
    originalPoint := Point{X: 1, Y: 1}
    modifyPoint(originalPoint)
    // originalPoint remains unchanged; it's still {X: 1, Y: 1}
}
```

### Pointer Semantics

With pointer semantics, you're working with a reference to the original data, rather than a copy. Reference types like slices, maps, and channels exhibit pointer semantics, as do pointers to any type, whether built-in or user-defined.

When you pass a variable with pointer semantics to a function or assign it to another variable, you're actually sharing the underlying data. Any changes made will be visible through all variables that reference that data.

Here's an example using pointers to a struct:

```go
func modifyPoint(p *Point) {
    p.X = 10
    p.Y = 20
}

func main() {
    originalPoint := &Point{X: 1, Y: 1}
    modifyPoint(originalPoint)
    // originalPoint is modified; it's now {X: 10, Y: 20}
}
```

### Summary - Value Semantics & Pointer Semantics 

- **Value Semantics**: When you pass or assign data, a copy is made. Changes to the copy do not affect the original data.
- **Pointer Semantics**: When you pass or assign data, a reference is used. Changes through the reference affect the original data.

The choice between value and pointer semantics depends on various factors like:

1. **Performance**: Copying large structs could be expensive in terms of memory and time, whereas passing a pointer is usually cheap.
2. **Mutability**: If you want to modify the original data, you have to use pointer semantics.
3. **Safety**: Value semantics make it easier to reason about the program's behavior since data isn't shared between different parts of the program.

Go allows you to choose between these semantics, offering flexibility in how you design your programs.

# Conversion vs Casting

### Overview
- Conversion creates a new variable of a different type, allocating new memory.
- Casting changes the type associated with existing bits

In the Go programming language, conversion and casting are sometimes used interchangeably, but they have subtle differences. In general, both refer to the process of converting a value from one type to another. However, they are used in different contexts and for different kinds of transformations.

### Conversion
Conversion in Go is a way to create a new value from an existing one, but the new value has a different type. You usually do this with a built-in conversion operation, where the source type and destination type are explicitly mentioned. For simple types, this operation is straightforward.

For example, converting an `int` to `float64`:
```go
var myInt int = 42
var myFloat float64 = float64(myInt)
```

Or converting a `[]byte` to `string`:
```go
var myBytes []byte = []byte{72, 101, 108, 108, 111}
var myString string = string(myBytes)
```

The Go runtime will handle the details of the conversion, and as long as it's a valid conversion (e.g., there's no loss of data, or it's allowed by the language rules), it will happen smoothly.

### Casting

Casting is usually used to mean a more direct re-interpretation of the underlying bits that represent a value. In Go, you don't "cast" values in the way you might in languages like C or C++; Go is stricter about type conversions. The closest thing might be when you're working with `unsafe.Pointer` to perform low-level operations, but this is generally not recommended unless absolutely necessary.

For example, using `unsafe` to cast an `int` pointer to a `float64` pointer:
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 42
	var y *float64 = (*float64)(unsafe.Pointer(&x))

	fmt.Println(*y)  // Output will be undefined behavior
}
```

In this case, we're directly re-interpreting the underlying bits of an `int` as a `float64`, which is considered unsafe and should be avoided in idiomatic Go code.

### Summary - Conversion vs Casting

- **Conversion**: Involves a type change that's explicitly handled by Go's runtime, creating a new value of a different type based on an existing one.
- **Casting**: Involves re-interpreting underlying bits, typically using unsafe operations. Generally not recommended in idiomatic Go.

The term "casting" may be used casually to describe conversions in Go, but it's important to understand the subtleties to write correct and idiomatic Go code.

# Advantages of Statically Typed Languages

1. **Type Safety**: Because types are checked at compile-time, many type-related errors are caught early in the development cycle, which can make the code more robust and easier to maintain.
  
2. **Performance**: Knowing the types of variables at compile-time can enable the compiler to optimize the generated code more effectively, which often results in faster execution times.

3. **Code Clarity**: Explicit type declarations can serve as a form of documentation; they make it easier to understand how the program is supposed to work.

4. **Tooling**: The compile-time type information enables better support from various tools like IDEs, linters, and auto-completers, which can make development easier and more efficient.

### Examples

Here are some examples to illustrate the difference:

In a statically-typed language like Go, you might declare a variable as follows:

```go
var x int = 42
```

Here, `x` is explicitly declared to be of type `int`, and attempting to assign a value of another type to `x` will result in a compile-time error:

```go
x = "hello"  // Compile-time error
```

In a dynamically-typed language like Python, the type of a variable is determined at runtime:

```python
x = 42  # x is an integer
x = "hello"  # x is now a string
```

In this Python example, `x` can be reassigned to a value of a different type without any issues. The type is checked at runtime, not compile-time.

Statically-typed languages include C, C++, Java, Go, Rust, and many others. Dynamically-typed languages include Python, Ruby, JavaScript, and others. Some languages like TypeScript offer a static type layer on top of a dynamically-typed language (JavaScript, in the case of TypeScript).

# Go

The Go programming language, often referred to as Golang, is defined by its specification, which outlines the syntax, semantics, and core libraries. Here's a broad overview of what defines the Go language:

### Design Principles
Go was designed to be simple, efficient, and effective for system-level programming. Its design principles focus on:
- Strong typing
- Simplicity and readability
- Compilation speed
- Built-in concurrency (goroutines and channels)
- A rich standard library

### Syntax
Go's syntax is designed to be clean and straightforward, inspired by languages like C and Algol. It employs braces for block structure and uses standard operators.

### Data Types
Go has a set of basic data types like integers, floating-point numbers, and strings. It also supports composite types like arrays, slices, structs, and maps.

### Concurrency Model
One of the defining features of Go is its concurrency model, which is built around goroutines (lightweight threads managed by Go runtime) and channels (for communication between goroutines).

### Standard Library
Go comes with a rich standard library that provides a wide array of utilities and packages, ranging from HTTP servers to text manipulation and file I/O.

### Memory Management
Go uses garbage collection for automatic memory management.

### Interfaces
Go uses interfaces to provide a flexible and composable mechanism for abstracting functionality. Unlike some object-oriented languages, Go does not support inheritance and method overloading, aiming for a simpler and more straightforward type system.

### Error Handling
Go eschews traditional exception-handling constructs in favor of explicit error return values. This feature encourages developers to directly confront error-handling logic.

### Tooling
Go also has a robust set of development tools, including but not limited to:
- `go build` for compiling
- `go test` for testing
- `go get` for package management
- `go fmt` for formatting code

### Open Source
Go is open-source and has a thriving community that contributes to its ecosystem.

### Specification and Implementation
The definitive specification of the Go language is maintained by its creators. The primary implementation is the Go compiler, `gc`, originally written in C but now in Go itself.

The Go programming language is used for a wide range of applications including web backends, network servers, data pipelines, and machine learning, among others.

# Static Type vs Dynamic Type Language

In programming, a language is said to be "statically typed" when the type of a variable is known at compile-time instead of at runtime. This means that you have to explicitly declare the type of the variable or it has to be inferred by the compiler before the program runs. Because the type information is available at compile-time, the compiler can catch type errors before the program executes. This is in contrast to dynamically-typed languages, where the type of a variable can be changed at runtime and type errors are generally caught during execution of the program.

# Code Review check

### int
- using anything other than int
- Why are you using an int32 or int64; why not just int?

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce) 
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses