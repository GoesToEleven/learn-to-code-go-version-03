
# "I would have written a shorter letter, but I did not have the time."
#### ~ Blaise Pascal, a French philosopher and mathematician. 

# Video
    - [video #9 on youtube](https://youtu.be/QKHLjA6e7w8) 
    - [video #10 on youtube](https://youtu.be/3Z9jEtDYa8I) 
    - [video #11 on youtube](https://youtu.be/SXFfUGN8Maw) 

# Takeaways
1. &, *, *T
1. Understanding nil.
1. Values and pointers.
1. Value semantics and pointer semantics.
1. Everything in Go is pass by value.
1. Every single function gets its own copy of the data
1. Each goroutine has a stack.
1. Data is stored either on the stack or allocated on the heap.
1. The compiler determines if a variable should be on the stack or the heap with "escape analysis."
1. The stack is self-cleaning.
1. Heap allocation: a function shares memory up the call stack
1. Heap allocation: large values may be put on the heap
1. Heap allocation: a value's size is not known at compile time
1. Heap allocation: decoupling code with functions or interfaces
1. "Allocation" = constructed on the heap, not the stack.
1. The stack is contiguous and dynamic; it grows from 2k.
1. There are efficiency benefits to contigious data storage.
1. You can't point to data on another stack (that would be too inefficient) - shared data is on the heap.
1. go build -gcflags -m=2
1. //go:noinline

# Table of Contents

1. [Reference and dereference operators](#reference-and-dereference-operators)
1. [Understanding nil](#understanding-nil)
1. [Goroutine and stack relationship](#goroutine-and-stack-relationship)
1. [The stack and the heap](#the-stack-and-the-heap)
1. [Decoupling with functions and interfaces](#decoupling-with-functions-and-interfaces)
1. [CPU cores and hardware threads](#cpu-cores-and-hardware-threads)
1. [Value semantics and pointer semantics](#value-semantics-and-pointer-semantics)
1. [Factory functions not constructors](#factory-functions-not-constructors)
1. [Up and down the stack](#up-and-down-the-stack)
1. [Count the passes](#count-the-passes)
1. [Code review check](#code-review-check)
1. [Coupons for Go courses](#coupons-for-go-courses)

# Reference and dereference operators

In Go, the `&` and `*` operators are used to reference and dereference memory addresses, respectively.

### Address-of Operator (`&`)

The `&` operator is used to get the memory address of a variable. This operator generates a pointer that points to the variable.

Here's an example:

```go
package main

import "fmt"

func main() {
    x := 42
    p := &x // 'p' is a pointer to 'x'
    fmt.Println(p)  // Prints the memory address
}
```

### Dereferencing Operator (`*`)

The `*` operator is used to access the value that is stored at a particular memory address to which a pointer points. This is called dereferencing the pointer.

Here's an example:

```go
package main

import "fmt"

func main() {
    x := 42
    p := &x
    fmt.Println(*p)  // Dereferencing 'p' to get the value of 'x', prints 42
}
```

You can also modify the value stored at the address to which the pointer points by using the `*` operator on the left-hand side of an assignment:

```go
package main

import "fmt"

func main() {
    x := 42
    p := &x
    *p = 21  // Modifies the value of 'x' through 'p'
    fmt.Println(x)  // Prints 21
}
```

These operators are a fundamental part of working with pointers in Go, allowing for more flexible and sometimes more efficient data manipulation.

### Type notation for pointer types like *int

In the context of a type like `*int`, the asterisk `*` is not an operator but rather part of the ***type notation*** to indicate that variables of this type are pointers to integers. When used in a type declaration, the asterisk `*` specifies that the type is a pointer to another type.

In Go, if you have a variable of type `*int`, it means that this variable is a pointer that is meant to point to an integer in memory. You can dereference this pointer to get the integer value it points to, or you can modify the integer value by dereferencing the pointer on the left-hand side of an assignment.

Here's a simple example that demonstrates a variable of type `*int`:

```go
package main

import "fmt"

func main() {
    var x int = 42
    var p *int = &x  // 'p' is of type '*int', a pointer to an int
    
    fmt.Println(p)   // Print the memory address
    fmt.Println(*p)  // Dereference to get the value of 'x'

    *p = 21          // Modify the value of 'x' through the pointer 'p'
    fmt.Println(x)   // Print the modified value of 'x'
}
```

In this example, `p` is declared as a variable of type `*int`, and it stores the memory address of an `int` variable `x`. You can then dereference `p` to get or set the value of `x`.

# Understanding nil

In Go, `nil` is the zero value for pointers, slices, maps, channels, and interface types. The `nil` value serves as a useful indicator for "no value" or "not initialized."

Here's how `nil` is used in different contexts:

1. **Pointer**: When you declare a pointer variable without initializing it, its value will be `nil`. Attempting to dereference a `nil` pointer will result in a runtime panic.

    ```go
    var p *int
    if p == nil {
        fmt.Println("p is nil")
    }
    ```

2. **Slice**: A `nil` slice has a length and capacity of 0 and has no underlying array.

    ```go
    var s []int
    if s == nil {
        fmt.Println("s is nil")
    }
    ```

3. **Map**: A `nil` map is not initialized and any attempt to add keys to it will result in a runtime panic.

    ```go
    var m map[string]int
    if m == nil {
        fmt.Println("m is nil")
    }
    ```

4. **Channel**: A `nil` channel is useful for some channel operations, but you cannot send or receive from it.

    ```go
    var c chan int
    if c == nil {
        fmt.Println("c is nil")
    }
    ```

5. **Interface**: An interface value is `nil` if both its type and value are `nil`.

    ```go
    var i interface{}
    if i == nil {
        fmt.Println("i is nil")
    }
    ```

It's important to remember that `nil` can have different behaviors depending on the type it is associated with. For example, sending or receiving from a `nil` channel will block forever, while accessing a key in a `nil` map will not cause a panic but will return the zero value for the map's value type.

Knowing when and how to use `nil` in these various contexts is important for writing effective Go code.

# Goroutine and stack relationship

Each goroutine gets its own stack, which grows and shrinks as needed. The stack is a critical aspect of a goroutine, as it determines how function calls are managed, how local variables are stored, and how control is transferred among functions. Each goroutine starts with a small stack, typically 2KB in size, and this stack can grow or shrink dynamically as needed.

Here is how the stack is related to a goroutine:

### Function Calls

When a new function is called within a goroutine, a new stack frame is pushed onto that goroutine's stack. This stack frame contains information like local variables, return addresses, and a base pointer to the previous stack frame. When the function returns, its stack frame is popped off, and control is returned to the calling function.

### Local Variables

Each function's local variables are stored in its stack frame. These are specific to the goroutine and are not shared between different goroutines, providing a level of isolation. This makes goroutines safer to use in concurrent scenarios, compared to threads in languages that don't isolate stack memory between threads.

### Dynamic Sizing

The Go runtime monitors the stack size and automatically resizes it if necessary. If a function call would cause a stack overflow, Go can allocate a larger stack and copy the existing stack data to this new stack. This is all transparent to the developer, who doesn't have to worry about stack management. This is in contrast to languages like C, where stack overflow can result in undefined behavior.

### Cost-efficiency

Goroutines are lighter than threads in traditional systems partly because they start with a small stack that can grow and shrink dynamically. This enables you to spawn thousands or even millions of goroutines in the same program without running out of memory, which would not be practical with traditional threads that typically have a fixed and larger stack size.

### Context Switching

The stack is also crucial during context switching between different goroutines. When the Go scheduler switches from executing one goroutine to another, it saves the state of the current goroutine's stack so that it can be resumed later from where it left off.

### Stack Traces

When debugging or when an unrecoverable error occurs, the stack trace gives you a snapshot of the call sequence within the affected goroutine, allowing you to trace back the steps that led to the issue.

Understanding the relationship between goroutines and stacks is key to mastering Go's concurrency model. Knowing how the stack operates helps you write more efficient and safer code, particularly in concurrent and multi-threaded scenarios.

# The stack and the heap

As in many other languages, memory allocation can be broadly categorized into two types: stack and heap allocation. Both serve different purposes and have different characteristics, advantages, and disadvantages.

### Stack

1. **Speed**: Allocating memory on the stack is faster because it involves merely moving the stack pointer, as opposed to searching for a suitable block of memory, which is what heap allocation entails.

2. **Lifetime**: Memory allocated on the stack has a lifetime that is easily determined, as it's tied to the function call that allocated it. Once the function returns, its stack frame (and all memory associated with it) is popped off the stack, automatically freeing that memory.

3. **Size**: Each goroutine starts with a small stack that is typically 2KB in size, though it can grow (or shrink) dynamically as needed.

4. **Scoping**: Variables allocated on the stack are scoped to the function that allocated them. They are not accessible outside of that function.

5. **Concurrency**: Each goroutine has its own stack. This makes stack-allocated data inherently thread-safe (within the same goroutine).

6. **Management**: The stack is self-managed. Variables are allocated and deallocated automatically as functions push and pop stack frames. The stack is "self cleaning."

7. **Data**: The Go compiler does escape analysis at compile-time to decide whether a variable should live on the stack or heap.

### Heap

1. **Speed**: Heap allocation is slower than stack allocation due to the extra work involved in finding a suitable block of memory and due to the overhead of garbage collection.

2. **Lifetime**: Memory allocated on the heap has a dynamic lifetime. Its deallocation is managed by Go's garbage collector, not tied to the scope of any function.

3. **Size**: Heap size is much larger compared to the stack, limited only by the size of the addressable memory.

4. **Scoping**: Heap-allocated data can be accessed from any part of the program, provided you have a reference to it.

5. **Concurrency**: Data on the heap can be accessed by multiple goroutines, so synchronization mechanisms like mutexes may be needed.

6. **Management**: Memory management for the heap is more complex, relying on the garbage collector to identify and free unused objects.

7. **Data**: Complex data structures like large arrays, slices, maps, and channels are usually heap-allocated. Also, variables that escape the scope of the function they are defined in are allocated on the heap.

#### Summary

- **Stack**: Fast, automatically managed, function-scoped, and used for short-lived data.
- **Heap**: Slower, managed by the garbage collector, globally accessible, and used for long-lived data.

Understanding the differences between stack and heap allocation can help you write more efficient and effective Go programs.

# Decoupling with functions and interfaces

In Go programming, the concepts of decoupling code with functions or interfaces refer to the software design practices aimed at reducing the interdependencies among different parts of a program. By doing this, you make the system easier to understand, modify, test, and maintain. Below are the ways these two concepts help in achieving decoupling:

### Decoupling with Functions

1. **Single Responsibility Principle**: Writing functions that do one specific task makes it easier to reuse them, test them, and understand them.

2. **Encapsulation**: Functions can encapsulate behavior and allow you to abstract complex operations into single function calls, which can be used in multiple parts of your application. This way, you don't need to know the underlying details.

3. **Parameterization**: Functions often accept parameters, which allows them to be more flexible and generalized. This means that you can change the behavior of a function without changing the function's code, thereby reducing code coupling.

4. **Higher-Order Functions**: Go supports functions as first-class citizens, which means you can pass functions as arguments to other functions or return them. This can help you abstract out the varying parts of an algorithm or operation, leading to more decoupled and flexible code.

### Decoupling with Interfaces

1. **Abstraction**: Interfaces define a contract (methods that need to be implemented), not the implementation. This allows you to write code that depends on just the interface, not the concrete implementation, making it easier to swap out one implementation for another.

    ```go
    type Writer interface {
        Write([]byte) (int, error)
    }
    ```

    Any type that satisfies the `Writer` interface can be used where a `Writer` is expected, decoupling the code from specific implementations.

2. **Polymorphism**: When a function accepts an interface type as a parameter, it can operate on multiple types that satisfy that interface, allowing you to write more generic and reusable code.

    ```go
    func Save(w Writer, data []byte) error {
        // ... use w to write data
    }
    ```

3. **Composition**: Go supports composing interfaces from other interfaces, which allows you to build up abstractions without modifying existing code.

    ```go
    type ReadWriter interface {
        Reader
        Writer
    }
    ```

4. **Testing**: Using interfaces makes it easier to write unit tests. You can create mock implementations of the interface to test how functions behave with different types, without having to use the actual implementations.

By combining both functions and interfaces, you can create a clean architecture that is modular, testable, and easy to understand. Decoupling through functions and interfaces is one of the key principles in Go's design, enabling you to write robust and maintainable software.

# CPU cores and hardware threads

Understanding the relationship between a physical CPU, multiple cores on that CPU, and hardware threads requires diving into some computer architecture concepts. Here's a simplified breakdown:

### Physical CPU (Central Processing Unit)

A physical CPU, often just called a CPU, is an actual piece of hardware that you could theoretically hold in your hand. Older computers might have just one CPU, but these days even consumer-grade computers often have more than one, and servers may have many.

A CPU takes instructions from a computer's memory and performs operations based on those instructions, which can be anything from simple arithmetic to complex data manipulation and more.

### Cores

Each physical CPU can have multiple cores. A core is essentially a smaller CPU (or processor) built into a larger CPU. Each core can execute its instructions independently of the others, which means a multi-core processor can perform multiple tasks simultaneously.

For example, if you have a quad-core CPU, that means your physical CPU chip contains 4 independent cores that can execute tasks.

### Threads and Hardware Threads

Threads are the smallest sequence of programmed instructions that can be managed independently by the operating system scheduler. Software can use multiple threads to perform several tasks in parallel. Each core can handle a certain number of threads.

Some CPUs use a feature called "Simultaneous Multi-Threading" (SMT), known as "Hyper-Threading" in Intel parlance. This allows each core to handle multiple threads in a way that makes it appear as if each core is actually two (or more) "logical" cores to the operating system and applications. This is achieved by sharing certain resources of the core between the threads, making better use of idle resources and thus increasing throughput.

#### Example:
- **Physical CPU:** 1 chip
- **Cores:** 4 cores
- **Hardware Threads:** 8 threads (assuming each core can handle 2 threads due to Hyper-Threading)

In this example, the operating system would see 8 logical CPUs. However, the physical hardware only has 4 actual cores, and there's 1 physical CPU chip.

### Summary

- **Physical CPU:** The actual hardware component.  
- **Cores:** Subunits within a CPU, capable of executing instructions.  
- **Hardware Threads:** Even smaller units that share resources of a core to execute instructions, making it appear as if a single core is multiple logical cores.

In modern computing, these architectural elements work together to achieve high-performance multitasking.

# Value semantics and pointer semantics

The concept of value semantics and pointer semantics is important for understanding how data is stored, accessed, and manipulated.

### Value Semantics

Value semantics in Go mean that when you pass a variable across program boundaries, like to another function, you are actually passing a copy of its data rather than a reference to the data. This is true for all basic data types (like integers, floats, etc.) as well as more complex data structures like arrays and structs. This means changes to the copy do not affect the original value, unless the original value is a reference type. 

In Go, slices and maps are reference types because they internally contain pointers to data, so changes to elements within a slice or map will affect the original even when passed by value. However, the slice or map descriptor itself behaves with value semantics.

Here's an example using an integer:

```go
func modifyValue(x int) {
    x = 10
}

func main() {
    a := 5
    modifyValue(a)
    fmt.Println(a) // Output will be 5, not 10
}
```

Even with complex types like structs, value semantics are at play:

```go
type Person struct {
    Name string
    Age  int
}

func modifyStruct(p Person) {
    p.Name = "Changed"
}

func main() {
    person := Person{Name: "Original", Age: 30}
    modifyStruct(person)
    fmt.Println(person.Name) // Output will be "Original"
}
```

### Pointer Semantics

Pointer semantics allow you to manipulate the data directly by passing the memory address, also known as a reference, to the data. You can create a pointer using the `&` operator and dereference it using the `*` operator.

Here's an example using an integer:

```go
// modifyValue accepts a `pointer to an int` as an argument
func modifyValue(x *int) {
    *x = 10 // Dereference the address to work with the value stored there
}

func main() {
    a := 5
    modifyValue(&a) // Pass the address where a is stored
    fmt.Println(a) // Output will be 10
}
```

With structs:

```go
type Person struct {
    Name string
    Age  int
}

func modifyStruct(p *Person) {
    p.Name = "Changed"
}

func main() {
    person := Person{Name: "Original", Age: 30}
    modifyStruct(&person)
    fmt.Println(person.Name) // Output will be "Changed"
}
```

### Which to Use When

1. **Value Semantics**: Prefer value semantics when:
   - The data is small and easily copied.
   - You want to ensure that the function you are calling cannot modify the data.
   - The function does not need to maintain any state.

2. **Pointer Semantics**: Prefer pointer semantics when:
   - The data is large and you don't want the overhead of copying.
   - You want to modify the original data from within the function.
   - The object has internal state that needs to be maintained across function calls.

# Factory functions not constructors

A factory function in Go is a regular function that encapsulates the creation and initialization of a struct or type. This function is often used to hide the internal details of object creation and setup, making it easier to manage default values or do any additional setup before the object is ready for use.

Here's an example:

```go
type Animal struct {
    Name string
    Type string
}

// Factory function for creating an Animal
func NewAnimal(name, animalType string) Animal {
    return Animal{
        Name: name,
        Type: animalType,
    }
}

func main() {
    dog := NewAnimal("Fido", "Dog")
    fmt.Println(dog.Name, dog.Type) // Output: Fido Dog
}
```

### Constructors (Initialization Functions)

Though the term "constructor" is often used in Go conversations, it's important to note that Go does not have constructors in the traditional sense. Go doesn't have "constructors" in the way that languages like Java or C++ do, nor does it have a `new` or `this` keyword that performs special initialization on an object, but the language's idioms and features still allow for object initialization with factory functions, as detailed above.

### Differences

In programming, both factory functions and constructors serve the purpose of creating new objects, but they do so in slightly different ways and offer different advantages. Below are some of the differences between the two:

### Factory Function

A factory function is a plain function that returns an object. It doesn't require the use of the `new` keyword. Factory functions often return an object literal or an object composed of other objects. They can be simpler to use and understand, especially for those new to programming, because they are just functions.

#### Example in JavaScript:

```javascript
function createPerson(name, age) {
  return {
    name: name,
    age: age,
    sayHello: function() {
      console.log("Hello, my name is " + this.name);
    }
  };
}

const alice = createPerson('Alice', 30);
alice.sayHello(); // Output: "Hello, my name is Alice"
```

#### Advantages of Factory Functions:

1. **Simpler syntax**: No need to use `new`, `this`, or worry about `constructor` functions.
2. **Private variables**: It's easier to encapsulate private variables using closures.
3. **Can return any arbitrary object**: A factory function can produce an instance of an object that is an instance of any type.

### Constructor Function

A constructor function is used with the `new` keyword. It initializes the newly created object. In languages that have classes, like Java, Python, or more recent versions of JavaScript (ES6+), the constructor is a special method inside a class definition. 

#### Example in JavaScript (ES5):

```javascript
function Person(name, age) {
  this.name = name;
  this.age = age;
  this.sayHello = function() {
    console.log("Hello, my name is " + this.name);
  };
}

const bob = new Person('Bob', 40);
bob.sayHello(); // Output: "Hello, my name is Bob"
```

#### Disadvantages:

1. **Complexity**: Use of `this`, `new`, and prototype chains can make the code harder to read and understand.
2. **Lack of encapsulation**: Harder to keep private variables compared to factory functions.

# Up and down the stack

```go
func createUser() *user {
    u := user{"James", 32}
}
    // down the stack
    fmt.Println(&u)

    // up the stack
    return &u
``` 

# Count the passes
Count how many times the basketball is passed in this [VIDEO: Count the passes](https://www.youtube.com/watch?v=IGQmdoK_ZfY)

# Code review check

### Returning a pointer
- If you are returning a pointer, make it readable:

```go
// like this
func createUser() *user {
    u := user{"James", 32}
}
    return &u
``` 

```go
// NOT like this
func createUser() *user {
    u := &user{"James", 32}
}
    return u
``` 

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# "Kind words do not cost much. Yet they accomplish much."
#### ~ Blaise Pascal

![Blaise Pascal](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/05-pointers/images/bp.png)

source: https://upload.wikimedia.org/wikipedia/commons/1/1d/Pascal_Pajou_Louvre_RF2981.jpg