# "If you want to build a ship, don’t drum up the men to gather wood, divide the work, and give orders. Instead, teach them to yearn for the vast and endless sea."
― Antoine de Saint-Exupéry

# Video
[SOON TO BE PUBLISHED Here is the video on youtube]()

# Takeaways
1. Slices

# Table of Contents

1. [Slice fundamentals](slice-fundamentals)
1. [Three categories of types in Go](three-categories-of-types-in-go)
1. [Remember value semantics](remember-value-semantics)
1. []()
1. []()
1. []()

# Slice fundamentals

A slice is a more flexible alternative to an array. While arrays have a fixed size, slices can grow or shrink dynamically. A slice is built on top of an array. A slice references a segment of an underlying array. A slice is a three word data structure (24 bytes) and includes a pointer to the array, a length, and a capacity.

```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
You can find this here: 
[src/runtime/slice.go](https://cs.opensource.google/go/go/+/refs/tags/go1.21.1:src/runtime/slice.go)

Here's a breakdown of slices in Go:

1. **Definition**: Unlike arrays, slices don't have a **fixed** size. Here's how you declare a slice:
   ```go
   var numbers []int
   ```

2. **Initialization**:
   - Using a literal:
     ```go
     numbers := []int{1, 2, 3, 4, 5}
     ```
   - Using the built-in `make` function:
     ```go
     numbers := make([]int, 5)      // Initializes a slice of length 5 and capacity 5, with all values set to 0
     numbers := make([]int, 3, 5)  // Initializes a slice of length 3 and capacity 5, with the first three values set to 0
     ```

3. **Length and Capacity**:
   - `len(slice)`: Returns the number of items in the slice.
   - `cap(slice)`: Returns the capacity of the slice, i.e., the number of elements in the underlying array starting from the first element of the slice.

Length sets the total number of initialized elements that we can read and write to in a slice.
Capacity sets the size of the underlying array. This provides us with capacity for future growth. If we exceed capacity, a new underlying array will be created. However, if we initially create a large enough underlying array, we can avoid this small performance cost of a new underlying array having to be created.

4. **Sub-slicing**:
   ```go
   numbers := []int{1, 2, 3, 4, 5}
   sub := numbers[1:4]  // sub is now [2, 3, 4]
   ```

5. **Appending**:
   You can grow a slice by appending values using the built-in `append` function:
   ```go
   numbers := []int{1, 2, 3}
   numbers = append(numbers, 4, 5, 6)  // numbers is now [1, 2, 3, 4, 5, 6]
   ```

   If the underlying array can't accommodate the new values, `append` will allocate a new array for the slice.

6. **Copying**:
   Use the built-in `copy` function to copy elements from one slice to another:
   ```go
   src := []int{1, 2, 3}
   dst := make([]int, len(src))
   copy(dst, src)
   ```

7. **Underlying Array**:
   When you create a slice, it's associated with an underlying array. Multiple slices can reference the same array. Modifying the elements of one slice can change the values in another slice if they share the same underlying array.

8. **Nil Slices**:
   A slice that hasn't been allocated with any storage has a value of `nil`, a length of 0, and no capacity. `nil` is the zero value for pointers.

9. **Range**:
   You can iterate over a slice using the `range` keyword:
   ```go
   numbers := []int{1, 2, 3, 4, 5}
   for index, value := range numbers {
       fmt.Println(index, value)
   }
   ```

Slices offer a more dynamic and flexible way to manage sequences of data compared to arrays. They encapsulate the power of dynamic-sized collections while maintaining the simplicity and efficiency that Go is known for. Always be mindful of the underlying array and how slices can share data, as this can lead to unexpected behavior if not understood correctly.

# Three categories of types in Go
- builtin: numeric, string, bool, array
- reference: slices, maps, channels, pointers
- struct

# Remember value semantics
When you pass a slice into `fmt.Println`, you are crossing a program boundary. The slice is being passed by value into the `Println` function. `Println` is receiving the value of the slice as an argument. `Println` then assigns that value to its own variable. This variable stores a slice data structure which points to the same underlying array as the value that was passed into `Println`.
```go
xi := []int{43,43,44}
fmt.Println(xi)
```
Rule of thumb: use value semantics to move around:
- builtin types
- reference types

# Code review check
Rule of thumb: use value semantics to move around:
- builtin types
- reference types

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
