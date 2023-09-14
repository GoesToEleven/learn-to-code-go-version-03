# ""

# Video
[SOON TO BE PUBLISHED Here is the video on YouTube]()

# Takeaways
1. Functions and modularity

# Table of Contents

1. [Mechanics](#mechanics)
1. [Looping considerations](#looping-considerations)
1. [Pointer and value semantics](#pointer-and-value-semantics)
1. []()
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

# Looping considerations

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
