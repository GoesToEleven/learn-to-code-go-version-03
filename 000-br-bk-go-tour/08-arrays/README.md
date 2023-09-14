# ""

# Video
[SOON TO BE PUBLISHED Here is the video on youtube]()

# Takeaways
1. Functions and modularity

# Table of Contents

1. [Mechanics](#mechanics)
1. [The for range clause](#the-for-range-clause)
1. []()
1. []()
1. []()
1. []()
1. [Pointer and value semantics](#pointer-and-value-semantics)
1. []()
1. []()
1. []()
1. [#Code review check](#code-review-check)

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

# Pointer and value semantics

# Code review check

### N
- n

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
