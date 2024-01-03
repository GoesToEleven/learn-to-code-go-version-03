# "If you can't fly, then run, if you can't run, then walk, if you can't walk, then crawl, but whatever you do you have to keep moving forward."
― Martin Luther King Jr

# Video
[SOON TO BE PUBLISHED Here is the video on youtube]()

# Takeaways
1. Arrays, order, type
1. Arrays are fixed in size
1. Size is part of an array's type
1. Arrays use contiguous blocks of memory
1. For range clause
1. Pass by value
1. Data semantics: value semantics and pointer semanatics
1. Go is wysiwyg

# Table of Contents

1. [Mechanics](#mechanics)
1. [The for range clause](#the-for-range-clause)
1. [Size is part of the type](#size-is-part-of-the-type)
1. [Data semantics](#data-semantics)
1. [Value semantics and pointer semantics](#value-semantics-and-pointer-semantics)
1. [Code review check](#code-review-check)

# Mechanics

Arrays are fundamental data structures in the Go programming language, as well as in many other languages. They are ordered collections of elements of the same type. Arrays in Go are simple but powerful, and they are valuable for storing multiple values pf the same type in a structured format. 

Here are some essential points about arrays in Go:

### Defining Arrays

To define an array in Go, you specify its size and type. Here's the basic syntax:

```go
var arrayName [size]Type
```

For example, to declare an array of integers with size 5, you would do:

```go
var numbers [5]int
```

### Initializing Arrays

You can initialize an array at the time of declaration:

```go
var numbers = [5]int{1, 2, 3, 4, 5}
```

Or you can initialize it using shorthand syntax:

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

If you want Go to count the elements for you, you can use an ellipsis `...`:

```go
numbers := [...]int{1, 2, 3, 4, 5}
```

### Accessing Array Elements

Elements in an array are accessed using zero-based indices:

```go
firstNumber := numbers[0]  // Access the first element
secondNumber := numbers[1] // Access the second element
```

### Modifying Array Elements

You can modify an existing value in an array using its index:

```go
numbers[0] = 10 // Set the first element to 10
```

### Array Length

You can find the length of an array using the `len()` function:

```go
length := len(numbers)
```

### Looping Through an Array

You can loop through an array using the `for` loop:

```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

Or using the `range` keyword:

```go
for i, v := range numbers {
    fmt.Println("Index:", i, "Value:", v)
}
```

### Limitations

- Arrays in Go are of fixed length, and the length is part of the type. Thus, `[5]int` and `[10]int` are distinct types.
- Arrays in Go are values. Assigning one array to another copies all the elements.
  
### Multi-dimensional Arrays

You can also define multi-dimensional arrays:

```go
var matrix [3][3]int // 3x3 array
```
### Summary

Arrays in Go are straightforward but rigid in their structure, which often leads Go developers to use slices for more flexible operations. However, arrays are still useful for performance-sensitive tasks or when you know the size won't change.

# The for range clause

The `for range` clause is a convenient way to iterate over arrays (as well as other iterable data structures like slices, maps, and channels). It offers a simplified syntax for looping and provides easy access to both the index and the value of each element in the array.

Here's a quick overview of how you can use `for range` with arrays.

### Getting Both Index and Value

You can obtain both the index and the value of each element in the array like this:

```go
arr := [5]int{1, 2, 3, 4, 5}
for i, v := range arr {
    fmt.Println(i, v)
}
```

In this example, the variables `i` and `v` get the index and value of each element in the array `arr`, respectively.

### Getting Only the Index

If you're only interested in the index, you can omit the second variable:

```go
arr := [5]int{1, 2, 3, 4, 5}
for i := range arr {
    fmt.Println(i)
}
```

### Getting Only the Value

If you're only interested in the value, you can use an "blank identifier" underscore `_` to ignore the index:

```go
arr := [5]int{1, 2, 3, 4, 5}
for _, v := range arr {
    fmt.Println(v)
}
```

### Modifying Array Elements

It's worth remembering that everyting in Go is PASS BY VALUE. When you use the `for range` loop, a copy of the element in the array is assigned to `v`. Modifying `v` will not affect the array. If you need to modify the array elements, you must use the index to access the array directly.

Here's an example:

```go
arr := [5]int{1, 2, 3, 4, 5}
for i := range arr {
    arr[i] = 777
}
```

This will set each element in the array `arr` to 777.

### Summary

The `for range` clause offers an idiomatic and readable way to iterate over arrays. It allows you to access both the index and value, or either one, depending on what you need for your specific task. It makes the code cleaner and more straightforward, but remember that everyting in Go is PASS BY VALUE. This means that the range clause gives you a new variable, `v` in our example, with an array element VALUE assigned to it. To modify an element in an array, reference that element by index position in the array.

# Size is part of the type

The size of an array is part of its type. 

This means that arrays of different sizes are considered to be of different types, even if they hold the same kind of elements. 

For example, an array of integers with a length of 3 (`[3]int`) is a different type from an array of integers with a length of 4 (`[4]int`):

```go
var array1 [3]int // An array of integers with a length of 3
var array2 [4]int // An array of integers with a length of 4
```

Because the size is part of the type, you can't assign an array of one size to a variable of an array type of a different size:

```go
array1 = array2 // This will result in a compilation error
```

This feature has both advantages and disadvantages:

### Advantages:

1. **Type Safety**: The size being part of the type allows the Go compiler to catch size-related errors during compile-time rather than run-time.
2. **Optimization**: Since the size is known at compile-time, various optimizations can be performed.
3. **Predictability**: When you're working with an array, you know exactly how much memory it will take up.

### Disadvantages:

1. **Inflexibility**: If you need to resize an array, you have to create a new one and copy the elements, which is not as efficient.
2. **Code Duplication**: Functions that work on arrays of different sizes will either need to be duplicated or require more complex interfaces to handle multiple sizes.

To overcome some of these limitations, Go also provides slices, which are built on top of arrays but offer more flexibility. Slices are not burdened by this limitation, as their size is not part of their type. This makes slices more versatile for many applications in Go, but they are not as performance-efficient as arrays in certain use-cases.

# Data semantics

In the context of the Go programming language, "data semantics" generally refers to the meaning and structure associated with the data types, variables, and expressions in the language. Understanding the semantics helps to reason about the behavior and purpose of Go code.

Here are some important aspects of data semantics in Go:

### Data Types
Go has several built-in data types, such as integers, floats, strings, and booleans. Additionally, Go supports complex types like arrays, slices, structs, and maps. Each data type has its own semantic meaning.

- **int**: Represents integer numbers.
- **float64**: Represents floating-point numbers.
- **string**: Represents a sequence of characters.
- **bool**: Represents a true or false value.
  
### Variables
In Go, variables must be declared before they are used. Variables hold data of a specific type.

```go
var i int = 42
var f float64 = 3.14
```

### Constants
Constants in Go are similar to variables, but their values cannot be changed once declared.

```go
const pi = 3.14159
```

### Type Conversion
Explicit type conversion is required in Go, emphasizing type safety.

```go
var i int = 42
var f float64 = float64(i)
```

### Operators
Go supports standard arithmetic, comparison, and logical operators. These operators are part of the language semantics, specifying how data should be manipulated.

```go
result := 10 + 20  // arithmetic operation
isEqual := (i == 42)  // comparison operation
```

### Control Structures
Data semantics also extend to how data is processed or manipulated through control structures like loops and conditionals.

```go
if i > 40 {
  // Do something
}
```

### Functions
Go has first-class functions, and they also play a part in data semantics, especially when considering function signatures and return types.

```go
func add(x int, y int) int {
  return x + y
}
```

### Pointers
Pointers in Go are a way to reference memory locations, which is different from directly referencing the data.

```go
var x int = 42
var p *int = &x
```

### Methods and Interfaces
Methods define the behavior of struct types, and interfaces define sets of methods that types can implement.

```go
type Geometry interface {
  Area() float64
}

type Square struct {
  Side float64
}

func (s Square) Area() float64 {
  return s.Side * s.Side
}
```

Understanding the semantics of these different constructs helps you to write effective and accurate Go programs.

# Value semantics and pointer semantics

Value semantics and pointer semantics refer to how data is passed around and manipulated in your program. Both have their use cases, advantages, and disadvantages.

### Value Semantics

When using value semantics, data is passed or assigned by value. This means that a copy of the actual data is created and used, making the original data and its copy independent of each other.

#### Advantages:
- **Immutability**: The original data remains unchanged, promoting safer and more predictable code.
- **Ease of Reasoning**: Because data is not shared, it's generally easier to understand and debug code that uses value semantics.

#### Disadvantages:
- **Performance**: Copying large data structures can be expensive in terms of memory and CPU usage.

#### Example:
```go
type person struct {
    age int
}

func main() {
    p1 := person{age: 55}
    p2 := p1        // Value is copied
    p2.age = 57     // Doesn't affect p1
}
```

### Pointer Semantics

When using pointer semantics, a pointer to the data is passed or assigned rather than the data itself. This means that changes made via the pointer are reflected in the original data.

#### Advantages:
- **Performance**: Passing a pointer is usually more efficient than making a copy, especially for large data structures.
- **Shared State**: Pointer semantics allow you to easily update shared state.

#### Disadvantages:
- **Complexity**: Code can become harder to understand and maintain, as multiple parts of the program can affect shared data.
- **Safety**: Incorrect use of pointers can lead to bugs, crashes, or security vulnerabilities.

#### Example:
```go
func main() {
    p1 := &person{age: 55}
    p2 := p1        // Both point to the same data
    p2.age = 57     // Affects both p1 and p2
}
```

### Choosing Between Value and Pointer Semantics

The choice between value and pointer semantics often depends on the specific use case:

- Use value semantics when you want to ensure that the data remains immutable and easy to reason about.
- Use pointer semantics when you need to modify the original data or when passing around large data structures for performance reasons.

In practice, Go's standard library and idiomatic code often combine both value and pointer semantics depending on the needs of the task at hand. It's important to carefully consider the implications of each to write effective and maintainable Go code.

# Code review check

### Use value semantics for builtin types
- Builtin types
    - bool
    - numerics
    - strings
- Exception: 
    - use pointers to represent "null", eg, no value, eg, a database not even returning a field
    - you can have a `*int` "pointer to an int" which is nil, because nil is the zero value of a pointer; whereas if you just had an int, you would have the zero value of zero.

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
1. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses