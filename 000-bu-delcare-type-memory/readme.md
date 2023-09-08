# Notes - Variables

### Type
- type tells us two things
    - size: amount of memory used
    - representation: what is stored there
- three classes of types
    - builtin: numeric, string, bool
    - reference
    - struct
- initialize = declaring and assigning
- var
    - use for zero value
- short declaration operator
    - inside code block
- conversion over casting
    - conversion has more safety than casting
    - conversion allocates new memory and copies values in

# Types & Memory
`bool` uses 1 byte of memory
`int` uses 1-8 bytes of memory
`string` uses 2 bytes of memory
`slice` uses 3 bytes of memory

# Declare, Assign, Initialize, Zero Value
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

# Built-in, Reference, & Struct types

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

# String is a built-in type

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

# Value Semantics & Pointer Semantics 

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