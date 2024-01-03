
# "Order the fields in your struct in a logical way that makes sense for **readability** and not in terms of padding. But if we ever found ourselves in a situation where we needed to reduce the amount of memory we were allocating, now this understanding of padding can allow us to do some form of micro-optimization." ~ Bill Kennedy

### "Here's the reality - I don't care about any of this stuff." ~ A friend who is a Go genius

# Video
    - [video #7 on youtube](https://youtu.be/sKLhcphr40E) 
    - [video #8 on youtube](https://youtu.be/yJeQ6OyU8dc) 

# Takeaways
1. Language mechanics
2. Optimize for readability
3. Use value semantics as much as possible - more readable; less heap allocations.
4. The best way to get rid of your padding is to order your fields from largest to smallest.
5. Method sets are the methods used for INTERFACE IMPLEMENTATION: T, T and *T, T & *T 
5. These are guidelines. There are exceptions to everything.

# Table of Contents

1. [Struct Literal, Anonymous Struct, Zero Value](#struct-literal-anonymous-struct-zero-value)
2. [Methods](#methods)
3. [Method Set Preview](#method-set-preview)
4. [Embedding Types & Inner-Type Promotion](#embedding-types-innertype-promotion)
5. [Composite Types & Aggregate Types](#composite-types)
6. [Mechanical Sympathy, Alignment, & Padding Bytes](#mechanical-sympathy)
7. [CPU Cycles & CPU Operations Per Cycle](#cpu-cycles)
8. [Understanding The Alignof Function](#understanding-the-alignof-function)
9. [Field Alignment Analysis Tool](#field-alignment-analysis-tool)
10. [Code review](#code-review)
11. [Coupons for Go courses](#coupons-for-go-courses)

# Struct Literal Anonymous Struct Zero Value

A `struct` is a composite data type that groups together variables under a single name for more organized and effective data management.

```go
type Person struct {
    Name string
    Age  int
}
```

### Struct Literal

A struct literal is an expression that creates a new struct value and fills in its fields. 

You can create a new `Person` instance using a struct literal like this:

```go
p1 := Person{"James", 30}  // Without field names
p2 := Person{Name: "Jenny", Age: 27}  // With field names

p3 := Person{
    Name: "Q", 
    Age: 72,    // Note trailing comma
}  

```

### Anonymous Struct

Anonymous structs are structs that don't have a formal type definition.

Example:

```go
point := struct {
    X, Y int
}{10, 20}

fmt.Println(point.X, point.Y)  // Output: 10 20
```

In this example, `point` is variable that stores an anonymous sturct. Notice that the struct isn't defined as a type like the struct `people` above.

### Struct Initialized to its Zero Value

If a struct is not explicitly initialized, its fields are given their zero values. For numeric types, the zero value is `0`; for strings, it's an empty string `""`; for booleans, it's `false`; and for pointers, slices, and maps, it's `nil`.


```go
var p4 Person  // Initialized to zero value
fmt.Println(p4.Name)  // Output: "" (zero value for string)
fmt.Println(p4.Age)   // Output: 0 (zero value for int)
```

In this example, `p4` is of type `Person` with the struct initialized to its zero value, so `p.Name` will be an empty string and `p.Age` will be `0`.

### Struct Initialized to its Zero Value - empty struct literal

Some people use an 'empty struct literal' to initialize a struct to its zero value, but generally speaking, only use the 'var' idiom, as above with `var p4 person` to create a struct value initialized to its zero value.

**=This is not idiomatic:**
```go
p5 := person{}
```

# Methods
Methods are functions associated with a particular type. You can call these methods using "dot" notation from values of a type that has methods. Methods allow us to define behaviors specific to types, similar to object-oriented programming paradigms in languages like Java or Python. The syntax for methods involves specifying a **"receiver"** which indicates the type with which the method should be associated. The receiver is basically a special kind of argument that receives the value (or a pointer to the value) on which the method is invoked.

Let's define some methods for the `Person` struct:

```go
type Person struct {
    Name string
    Age  int
}
```

### Value Receiver

A value receiver gets a copy of the struct that it operates on.

```go
func (p Person) Describe() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}
```

Here, `Describe` and `IsAdult` are methods associated with the `Person` struct. Each takes a receiver `p` of type `Person`. In these methods, `p` is a copy of the original `Person` instance.

To use these methods:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Describe() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}

func main() {
    john := Person{"John", 28}
    fmt.Println(john.Describe()) // Output: Name: John, Age: 28
    fmt.Println(john.IsAdult())  // Output: true
}
```

### Pointer Receiver

Using a pointer receiver allows the method to modify the original struct value. It also avoids making a copy of the struct, which might be more efficient for large structs. **Please note, just so this is clear in your head, everything in Go is pass by value.** With a pointer receiver, the method receives a value which is a pointer to an address, then it a makes a copy of that pointer to an address and assigns it to the receiver's parameter. In the `HaveBirthday` method, n these methods, `p` is a copy of the address pointing to the `Person` instance.

```go
func (p *Person) HaveBirthday() {
    p.Age++
}
```

Here, `HaveBirthday` has a pointer receiver (`*Person`), allowing it to modify the original `Person` instance.

Example usage:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Describe() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}

func (p *Person) HaveBirthday() {
    p.Age++
}

func main() {
    john := Person{"John", 28}
    fmt.Println(john.Describe()) // Output: Name: John, Age: 28

    john.HaveBirthday()
    fmt.Println(john.Describe()) // Output: Name: John, Age: 29
}
```

### When to use Value Receivers and Pointer Receivers

These are general guidelines:

- Use a value receiver when you don't need to modify the receiver or when you want to work with a copy of the data.
    - You can also return a modified receiver (the value of the struct type to which the method is attached). This allows you to modify the receiver, and then return it.
    - **Use value receivers as much as possible.** Switch to pointer receivers only if absolutely necessary. ***Value semantics is more readable and also reduces allocations on the heap.***
    - Using value receivers is also known as **value semantics.**
- Use a pointer receiver for methods that need to modify the receiver or when you want to avoid the cost of copying large structs.
    - Using pointer receivers is also known as **pointer semantics**.

### Method set preview

1. When we learn about **"method sets"**, we will see that ***THE METHOD SET OF A TYPE DETERMINES WHAT INTERFACES THE TYPE IMPLEMENTS.*** 
2. The method set of a type is influenced by whether the methods use value semantics (value receivers) or pointer semantics (pointer receivers).

#### Method Sets for Value Types

If you have a type `T`, the method set of the value type `T` includes all methods that have a value receiver of type `T`. This set of methods is the set used to determine which interfaces the the value of type `T` implements.

For example, consider the following `Person` struct and methods:

```go
type Person struct {
	Name string
	Age  int
}

// Value receiver
func (p Person) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// Value receiver
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Pointer receiver
func (p *Person) RunsPtr() bool {
	fmt.Println("Printing from in RunsPtr")
	p.Age++
	return true
}

// IMPORTANT
// Notice that both value semantics & pointer semantics run methods
// with both value receivers and pointer receivers.
// Methods sets are important for determining INTERFACE IMPLEMENTATION.

func PtrValSem() {
	fmt.Printf("VALUE SEMANTICS \n \n")

	p1 := Person{"James", 32}
	fmt.Println(p1.Describe())
	fmt.Println(p1.IsAdult())

	fmt.Println()

	fmt.Println(p1.Age)
	fmt.Println(p1.RunsPtr())
	fmt.Println(p1.Age)

	fmt.Printf("\nPOINTER SEMANTICS \n \n")

	p2 := &Person{"Jenny", 27}
	fmt.Println(p2.Describe())
	fmt.Println(p2.IsAdult())

	fmt.Println()

	fmt.Println(p2.Age)
	fmt.Println(p2.RunsPtr())
	fmt.Println(p2.Age)
}

```

Here, the method set of the value type `Person` ONLY INCLUDES VALUE RECEIVERS `Describe` and `IsAdult`:

```go
p1 := Person{"James", 32}
```

#### Method Sets for Pointer Types

For a pointer type `*T`, the method set includes all methods with either value receivers or pointer receivers of type `T`. This means that you can call both pointer-receiver methods and value-receiver methods on a pointer of type `*T`.

In the code above, the method set of the value type `*Person` INCLUDES VALUE RECEIVERS AND POINTER RECEIVERS `Describe`, `IsAdult`, and `RunsPtr`:

```go
p2 := &Person{"Jenny", 27}
```

***THE METHOD SET OF A VALUE DETERMINES INTERFACE IMPLEMENTATION.***

### In Relation to Interfaces

Understanding method sets becomes particularly important when working with interfaces. A type must have all the methods specified in an interface to implement that interface. Whether the methods use value receivers or pointer receivers will affect interface implementation.

For example, consider the interface:

```go
type Describer interface {
    Describe() string
}
```

In this case, both `Person` and `*Person` implement the `Describer` interface because `Describe` is a value-receiver method.

But if the interface were:

```go
type Runner interface {
    RunsPtr()
}
```

Only `*Person` would implement the `Runner` interface because `RunsPtr` uses a pointer receiver.

Understanding method sets and the distinctions between value and pointer semantics can help you design better Go programs, especially when interfaces are involved.

# Embedding Types InnerType Promotion
You can include one struct type within another by **embedding it as a field without a field name.** This is known as "embedding" and it provides a way to reuse code and model is-a or has-a relationships. When you embed a type in another struct, the fields and methods of the embedded **"inner"** type become accessible as if they were part of the **"outer"** type. This is often referred to as "promotion."

### Basic Embedding
Here's a simple example to illustrate embedding:

```go
package main

import "fmt"

type Address struct {
    Street string
    City   string
    State  string
    Zip    string
}

type Person struct {
    Name    string
    Age     int
    Address // Embedding Address struct
}

func main() {
    p := Person{
        Name: "John",
        Age:  30,
        Address: Address{
            Street: "123 Main St",
            City:   "Anytown",
            State:  "CA",
            Zip:    "12345",
        },
    }

    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
    fmt.Println("Street:", p.Street) // Note the promoted field
    fmt.Println("City:", p.City)       // Note the promoted field
}
```

In the above example, `Person` embeds `Address`. This allows us to directly access the `Street` and `City` fields on a `Person` object, as if they were fields of `Person` itself.

### Method Promotion
In addition to fields, methods are also promoted. If the embedded type has any methods, those become accessible as methods of the outer type.

```go
package main

import "fmt"

type Writer struct{}

// Embedding Writer
// An 'Advanced Writer' is also a 'Writer'
type AdvancedWriter struct {
    Writer 
}

// This is attached to 'Writer'
// It will be availabled to an 'Advanced Writer'
func (w Writer) Write(p []byte) (n int, err error) {
    n = len(p)
    fmt.Println(string(p))
    return n, nil
}


func main() {
    aw := AdvancedWriter{Writer{}}

    // The 'Advanced Writer' can Use the 
    // Write method of the embedded 'Writer'
    aw.Write([]byte("Hello, World!")) 
}
```

### Overriding Methods and Fields

You can also override methods and fields of the embedded type by declaring them in the outer type:

```go
package main

import "fmt"

type Animal struct{}

// Every 'Dog' is also an 'Animal'
type Dog struct {
    Animal 
}

func (a Animal) Speak() {
    fmt.Println("Animal speaks.")
}

// Override 'Animal' Speak method
func (d Dog) Speak() {
    fmt.Println("Woof woof!")
}

func main() {
    d := Dog{}

    // Output will be "Woof woof!"
    d.Speak() 
}
```

In this example, we have a `Dog` struct that embeds an `Animal` struct. Both `Animal` and `Dog` have a method named `Speak()`. When we call `Speak()` on a `Dog` object, it uses the method defined in the `Dog` struct, overriding the one from `Animal`.

### Limitations and Ambiguity

When you have two embedded types with fields or methods having the same name, you can't access them directly from the outer struct to avoid ambiguity. In such cases, you'll have to specify which embedded type's field or method you want to access:

```go
package main

import "fmt"

type Writer struct {
    Name string
}

func (w Writer) Write() {
    fmt.Println("Writer is writing.")
}

type Reader struct {
    Name string
}

func (r Reader) Read() {
    fmt.Println("Reader is reading.")
}

type WriterReader struct {
    Writer
    Reader
}

func main() {
    wr := WriterReader{
        Writer: Writer{Name: "Writer"},
        Reader: Reader{Name: "Reader"},
    }

    // Ambiguous field
    // This will throw an error
    // fmt.Println(wr.Name) 

    // Disambiguate
    fmt.Println(wr.Writer.Name)
    fmt.Println(wr.Reader.Name)

    // Methods aren't ambiguous
    wr.Write()
    wr.Read()
}
```

In this example, both `Writer` and `Reader` have a field named `Name`. If you try to access `Name` directly from `WriterReader`, it will result in a compilation error due to ambiguity. You have to specify whether you mean `Writer.Name` or `Reader.Name`.

So that's a quick overview of field embedding and inner type promotion in Go. It's a very powerful feature and it's used extensively in idiomatic Go code.

# Composite Types & Aggregate Types

Structs are known as composite types. 

The terms "composite types" and "aggregate types" are often used interchangeably. However, it's worth breaking down the nuances of these two terms to clarify the distinction.

### Composite Types

The term "composite types" refers to types that can be composed of other types. In Go, these are:

- Arrays
- Slices
- Maps
- Structs
- Channels

These types allow you to bundle multiple values, possibly of multiple types, into a single variable.

### Aggregate Types

The term "aggregate types" is not strictly defined in Go's language specification. However, in general computer science terms, an "aggregate type" is a type that groups together variables under a single name. In C or C++, for example, this would be a `struct` or an `array`.

In the context of Go, you might consider arrays and structs as aggregate types because they "aggregate" multiple values of the same or different types, respectively.

While all aggregate types (in the general sense) could be considered composite types in Go, not all composite types would traditionally be considered aggregate types. For example, channels and maps in Go could be considered composite but not aggregate, as they don't just bundle existing types together—they provide additional functionality and behavior.

In practice, the terms are often used interchangeably, and the specific meaning is usually clear from the context in which they are used.

# Mechanical Sympathy, Alignment, & Padding Bytes

### Mechanical Sympathy

Mechanical sympathy refers to writing code that is sympathetic to the underlying hardware it runs on. In other words, it's about understanding enough of the hardware's behavior to write software that runs efficiently on that hardware. For Go programmers, mechanical sympathy often entails understanding the machine-level implications of their code: how data is laid out in memory, how the Go runtime schedules Goroutines onto OS threads, how network packets are transmitted over the wire, etc.

### Alignment

In computer memory, data is usually read in "chunks" that align with the architecture's word size. For example, on a 64-bit machine, data is most efficiently accessed when it is aligned on 64-bit (8-byte) boundaries. Misaligned access usually results in slower performance due to additional CPU cycles spent in rearranging the data.

In Go, the language runtime and compiler handle most of the alignment issues, but you should still be aware of it, especially when you're trying to optimize data structures. The `unsafe.Alignof` function can help you understand the alignment of different types in Go.

```go
import (
	"fmt"
	"unsafe"
)

type Example struct {
	A bool    // 1 byte
	B int16   // 2 bytes
	C float64 // 8 bytes
}

func main() {
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int16(0)))
	fmt.Println("Alignment of float64:", unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of Example:", unsafe.Alignof(Example{}))
}
```

### Padding Bytes

To maintain alignment, compilers may insert "padding" bytes between fields of a struct. Padding bytes are essentially unused spaces that exist to ensure that the data structure aligns properly with the machine's word size. This can lead to larger-than-expected memory consumption, which can be critical in systems that are memory-constrained.

You can sometimes reduce the size of a struct by reordering its fields in a way that minimizes padding. Go's language specification does not guarantee the layout of the struct fields in memory, but in practice, you'll often find that the fields are laid out in the order they are defined.

```go
// Without optimization
type Unoptimized struct {
	A bool    // 1 byte
	B float64 // 8 bytes
	C int32   // 4 bytes
	// Padding: 7 bytes for A, 4 bytes for C (assuming 64-bit machine)
}

// With optimization
type Optimized struct {
	B float64 // 8 bytes
	C int32   // 4 bytes
	A bool    // 1 byte
	// Padding: 3 bytes for A (assuming 64-bit machine)
}
```

Understanding mechanical sympathy, alignment, and padding bytes can help you write Go code that runs more efficiently, especially in system-critical or resource-constrained environments.

# CPU Cycles & CPU Operations Per Cycle

The terms "CPU cycles" and "CPU operations" are often used interchangeably but they have distinct meanings in the context of computer architecture and performance analysis.

### CPU Cycles

A CPU cycle is the smallest unit of time for a CPU and is measured by the clock rate, often represented in gigahertz (GHz). For example, a 3 GHz CPU completes 3 billion cycles per second. During each cycle, the CPU can perform a certain amount of work, such as fetching instructions from memory, decoding instructions, executing instructions, or writing data back to memory.

However, not all actions can be completed in a single CPU cycle. Modern CPUs have multiple stages in their pipeline to fetch, decode, execute, and retire instructions. This allows them to work on multiple instructions simultaneously but also means that individual instructions may take multiple cycles to complete.

### CPU Operations

CPU operations refer to the actual tasks that a CPU performs, typically represented as instructions in machine code. These operations include arithmetic calculations, data movement, logical operations, and control flow changes, among others. Each operation may correspond to a specific machine language instruction or a set of instructions that the CPU can execute.

The relationship between CPU cycles and operations can be complex:

1. **CPI (Cycles Per Instruction)**: Not all operations are completed in one cycle. Some might require multiple cycles, depending on the complexity of the operation and the CPU architecture. The CPI metric gives an average number of cycles needed for an instruction.

2. **Superscalar Architecture**: Modern CPUs can often execute multiple instructions per cycle, thanks to features like multiple execution units and out-of-order execution.

3. **Instruction Pipelining**: Modern CPUs use pipelining to break down each operation into smaller steps, allowing the CPU to work on different stages of multiple instructions simultaneously.

4. **Instruction Set Architecture (ISA)**: Different CPUs may have different sets of operations, and some may even have specialized instructions that perform a complex task in a single operation, affecting the operation-to-cycle ratio.

### Summary

- **CPU Cycles**: A unit of time for the CPU, measured in clock speed (GHz). Defines how many cycles the CPU has in one second to perform tasks.
  
- **CPU Operations**: The actual tasks or instructions that a CPU performs. These can vary in complexity and may take multiple cycles to complete.

Understanding both concepts is crucial for performance tuning and optimization, both in software and hardware design.

# Understanding The Alignof Function

The `unsafe.Alignof` function in Go's `unsafe` package returns the alignment of a type in bytes. The alignment is the smallest number `n` such that an address `a` for a variable of that type will always be a multiple of `n`. Understanding alignment can be crucial for low-level programming, performance tuning, and interacting with C code using cgo.

Here's a basic example:

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int32
	var y int64
	var z float64

	fmt.Println("Alignment of int32:", unsafe.Alignof(x))
	fmt.Println("Alignment of int64:", unsafe.Alignof(y))
	fmt.Println("Alignment of float64:", unsafe.Alignof(z))
}
```

This will usually output:

```
Alignment of int32: 4
Alignment of int64: 8
Alignment of float64: 8
```

In this example, the alignment of `int32` is 4 bytes, `int64` is 8 bytes, and `float64` is 8 bytes as well. This means that, in memory, the starting address of a variable of type `int32` will be a multiple of 4, for `int64` it will be a multiple of 8, and for `float64`, it will be a multiple of 8.

On a 64-bit machine, that is a machine that has a word size of 64 bits or 8 bytes, this means is that 

- The int32 needs to be stored at byte location 0 or 4 in a word. 
- The int64 and float64 need to be stored at byte location 0 in a word.

A variable's starting address in memory must align with its alignment requirement. That is, the starting address should be a multiple of the alignment value. This ensures that the CPU can access the data efficiently, often in a single cycle, without having to deal with misalignment penalties.

### To further explain this example

1. **int32**: If a variable of type `int32` has an alignment of 4 bytes, it means that the starting address of this variable in memory would be a multiple of 4 (e.g., 0, 4, 8, 12, ...).

2. **int64 and float64**: Similarly, if a variable of type `int64` or `float64` has an alignment of 8 bytes, the starting address in memory would be a multiple of 8 (e.g., 0, 8, 16, 24, ...).

### Illustrations & Explanations #1

```go
	type Example struct {
		A bool    // 1 byte
		B int16   // 2 bytes
		C float32 // 4 bytes
	}
	// padding: 1 byte for A
```

In the `Example` struct, padding might be introduced to ensure that each field aligns with memory locations that are multiples of their respective sizes. When needed, the alignment requirement is typically the same as the size of the type for built-in types.

Here's how the memory layout might look, assuming the machine architecture requires alignment:

1. **Field `A` (bool)**: Takes 1 byte of memory, but will likely have 1 byte of padding after it to ensure that `B` starts at an address that is a multiple of 2.
  
   - `A`: 1 byte
   - Padding after `A`: 1 byte

2. **Field `B` (int16)**: Takes 2 bytes of memory and is aligned on a 2-byte boundary.

   - `B`: 2 bytes

3. **Field `C` (float32)**: Takes 4 bytes of memory and would typically be aligned on a 4-byte boundary.

   - `C`: 4 bytes

So, the struct might have 1 byte of padding after the `bool` field `A` to ensure that `B` starts at an address that is a multiple of 2.

Total Size: 1 byte (A) + 1 byte (padding) + 2 bytes (B) + 4 bytes (C) = 8 bytes

Please note that the Go language specification doesn't guarantee the layout of structs in memory, so the actual layout might differ depending on compiler optimizations and the target architecture. However, the layout above is commonly seen in practice on systems that require data alignment.


### Illustrations & Explanations #2

```go
	type Example struct {
		A bool    // 1 byte
		B int32   // 4 bytes
		C float32 // 4 bytes
	}
	// padding: 3 bytes for A
```

Let's examine the potential padding:

1. **Field `A` (bool)**: Takes 1 byte of memory, but will likely have 3 bytes of padding after it to ensure that `B` starts at an address that is a multiple of 4.
  
   - `A`: 1 byte
   - Padding after `A`: 3 bytes

2. **Field `B` (int32)**: Takes 4 bytes of memory and would typically be aligned on a 4-byte boundary.
  
   - `B`: 4 bytes

3. **Field `C` (float32)**: Tkes 4 bytes of memory and would typically be aligned on a 4-byte boundary.
  
   - `C`: 4 bytes

Here's how the struct would likely be laid out in memory, assuming the machine architecture requires alignment:

Memory layout:

```
+------+------+------+------+------+------+------+------+
|   A  | pad  | pad  | pad  |   B (4 bytes)    |   C (4 bytes)    |
+------+------+------+------+------+------+------+------+
1 byte   3 bytes       4 bytes          4 bytes
```

Total Size: 1 byte (A) + 3 bytes (padding) + 4 bytes (B) + 4 bytes (C) = 12 bytes

Again, the Go language specification doesn't guarantee specific memory layouts for structs, so the actual layout might differ depending on compiler optimizations and the target architecture. However, this layout would be typical for systems that require data to be aligned.

### Illustrations & Explanations #3

```go
	type Example struct {
		A bool    // 1 byte
		B int32   // 4 bytes
		C float64 // 8 bytes
	}
```

The padding and alignment would look like this:

1. **Field `A` (bool)**: Takes 1 byte of memory, but will likely have 3 bytes of padding after it to ensure that `B` starts at an address that is a multiple of 4.

    - `A`: 1 byte
    - Padding after `A`: 3 bytes

2. **Field `B` (int32)**: Takes 4 bytes of memory and would be aligned on a 4-byte boundary.

    - `B`: 4 bytes

3. **Field `C` (float64)**: Takes 8 bytes of memory and would typically be aligned on an 8-byte boundary. 

    - `C`: 8 bytes

Here's how the struct would likely be laid out in memory:

```
+------+------+------+------+------+------+------+------+------+------+
|   A  | pad  | pad  | pad  |   B (4 bytes)    |          C (8 bytes)          |
+------+------+------+------+------+------+------+------+------+------+
1 byte   3 bytes       4 bytes                          8 bytes
```

Total Size: 1 byte (A) + 3 bytes (padding) + 4 bytes (B) + 8 bytes (C) = 16 bytes

### Illustrations & Explanations #4 - Unoptimized

```go
type Unoptimized struct {
		A bool    // 1 byte
		B float64 // 8 bytes
		C int32   // 4 bytes
		// Padding: 7 bytes for A, 4 bytes for C (assuming 64-bit machine)
	}
```

In the `Unoptimized` struct, padding and alignment might look like this based on typical alignment requirements for the respective types:

1. **Field `A` (bool)**: Takes 1 byte of memory, but will likely have 7 bytes of padding after it to ensure that `B` starts at an address that is a multiple of 8.

    - `A`: 1 byte
    - Padding after `A`: 7 bytes

2. **Field `B` (float64)**: Takes 8 bytes of memory and would typically be aligned on an 8-byte boundary.

    - `B`: 8 bytes

3. **Field `C` (int32)**: Takes 4 bytes of memory and would typically be aligned on a 4-byte boundary.

    - `C`: 4 bytes
    - Padding after `C`: To align the size of the entire struct to a multiple of the largest field (float64, which has an alignment requirement of 8 bytes), you would have 4 bytes of padding at the end of the struct.

Here's how the struct would likely be laid out in memory:

Memory layout:

```
+------+------+------+------+------+------+------+------+------+------+------+------+------+
|   A  | pad  | pad  | pad  | pad  | pad  | pad  | pad  |          B (8 bytes)             | C (4 bytes) | pad  | pad  | pad  | pad  |
+------+------+------+------+------+------+------+------+------+------+------+------+------+
1 byte   7 bytes                         8 bytes                           4 bytes 4 bytes
```

Total Size: 1 byte (A) + 7 bytes (padding) + 8 bytes (B) + 4 bytes (C) + 4 bytes (padding) = 24 bytes

```bash
EXAMPLE 4 - Unoptimized
Alignment of bool: 1
Alignment of float64: 8
Alignment of int32: 4
Alignment of Unoptimized: 8
Size of Unoptimized: 24
```
***The entire struct also has to fall on an alignment for the largest field that is in that struct***

### Illustrations & Explanations #4 - Optimized

```go
	type Optimized struct {
		B float64 // 8 bytes
		C int32   // 4 bytes
		A bool    // 1 byte
		// Padding: A has 3 bytes (assuming 64-bit machine)
	}
```

In the `Optimized` struct, we've rearranged the fields to minimize padding. Let's examine the alignment and padding based on typical alignment requirements:

1. **Field `B` (float64)**: Takes 8 bytes of memory and would typically be aligned on an 8-byte boundary.
  
    - `B`: 8 bytes

2. **Field `C` (int32)**: Takes 4 bytes of memory and would typically be aligned on a 4-byte boundary. In this case, it would also naturally fall into an address that is a multiple of 4, assuming `B` started at an 8-byte-aligned address.
  
    - `C`: 4 bytes

3. **Field `A` (bool)**: Takes 1 byte of memory. Given that `C` ends on a 4-byte boundary, `A` would start at the next byte.
  
    - `A`: 1 byte
    - Padding after `A`: To align the size of the entire struct to a multiple of the largest field (float64, which has an alignment requirement of 8 bytes), you would have 3 bytes of padding at the end of the struct.

Here's how the struct would likely be laid out in memory:

Memory layout:

```
+------+------+------+------+------+------+------+------+
|                    B (8 bytes)                    |        C (4 bytes)         |   A  | pad  | pad  | pad  |
+------+------+------+------+------+------+------+------+
                        8 bytes                               4 bytes          1 byte   3 bytes
```

Total Size: 8 bytes (B) + 4 bytes (C) + 1 byte (A) + 3 bytes (padding) = 16 bytes

By rearranging the fields like this, we have minimized the padding needed, and the size of the struct is now 16 bytes instead of 24 bytes, as it was in the previous `Unoptimized` version. This is a typical optimization for systems that require data to be aligned.

### Why is this Important?

1. **Performance**: Misaligned reads and writes can be slower because they may require multiple CPU cycles and additional operations to complete.

2. **Hardware Requirements**: Some hardware architectures do not even support misaligned access and will trigger a runtime fault if attempted.

3. **Interoperability**: When interfacing with low-level languages like C or when doing system programming, knowing the alignment is essential for data structure compatibility.

4. **Memory Layout**: Understanding alignment can help you design data structures that require less padding, thereby using memory more efficiently.

### Caveats

- Using `unsafe` should be done carefully, as incorrect use can lead to various problems like undefined behavior and difficult-to-debug issues.
  
- Go does a good job of handling alignment automatically in most cases, so you typically don't need to worry about it for day-to-day coding tasks.

- The result of `unsafe.Alignof` is Go-implementation and architecture-specific. Therefore, it might differ when you switch from one Go implementation or CPU architecture to another.

Understanding alignment via `unsafe.Alignof` can provide insights into the memory layout and potential performance characteristics of your Go programs, but it's generally not something you need to worry about unless you are working on very low-level or performance-critical code.

# Field Alignment Analysis Tool

### Tool location & instructions

You can find the **most recent version** of this tool here:
[https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment	)

To see how to run this tool, look in this folder of my github repo, open up main.go, and read the comments:

learn-to-code-go-version-03/000-br-bk-go-tour/05-field-alignment-analyzer
[https://github.com/GoesToEleven/learn-to-code-go-version-03/tree/main/000-br-bk-go-tour/05-field-alignment-analyzer](https://github.com/GoesToEleven/learn-to-code-go-version-03/tree/main/000-br-bk-go-tour/05-field-alignment-analyzer)

# Code review check

### struct fields largest to smallest
- Why are you micro-optimizing?
- Optimize for readability first

### empty literal struct
- use var instead: var p person

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses