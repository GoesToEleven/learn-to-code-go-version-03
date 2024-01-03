# "If you want to build a ship, don’t drum up the men to gather wood, divide the work, and give orders. Instead, teach them to yearn for the vast and endless sea."
― Antoine de Saint-Exupéry

### "Every array in Go is just a slice waiting to happen."
- Bill Kennedy
(you use slicing on an array to get a slice)

### "A function cannot dictate how it works with data. The data tells the API how it's supposed to work."
- Bill Kennedy

# Video
[SOON TO BE PUBLISHED Here is the video on youtube]()

# Takeaways
1. Slices

# Table of Contents

1. [Slice fundamentals](slice-fundamentals)
1. [Three categories of types in Go](three-categories-of-types-in-go)
1. [Remember value semantics](remember-value-semantics)
1. [Fun fact about empty slice](fun-fact-about-empty-slice)
1. [Append is a value semantic mutation API](append-is-a-value-semantic-mutation-api)
1. [Memory leak scenarios](memory-leak-scenarios)
1. [Range clause](range-clause)
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
   - That declares a slice and sets it to its zero value which is `nil`
   - numbers is a `nil` slice

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
   - Length sets the total number of initialized elements that we can read and write to in a slice. 
   - Capacity sets the size of the underlying array. This provides us with capacity for future growth. If we exceed capacity, a new underlying array will be created. However, if we initially create a large enough underlying array, we can avoid this small performance cost of a new underlying array having to be created.
   - Don't preallocate `cap` unless you know the size needed.

3. **Length and Capacity**:
   - `len(slice)`: Returns the number of items in the slice.
   - `cap(slice)`: Returns the capacity of the slice, i.e., the number of elements in the underlying array starting from the first element of the slice.

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

# Fun fact about empty slice
- An empty slice points to an empty struct, not an underlying array.
- The struct{} with no fields is called an ***empty struct***. It has a size of zero bytes and is often used in scenarios where you need a placeholder data type or a signal type without carrying any information.

# Append is a value semantic mutation API
- Append is a beautiful example of a value semantic mutation API. [Append is a builtin function.](https://go.dev/ref/spec#Appending_and_copying_slices) When you use `append` to append to a slice, `append` gets its own copy of the value (the slice) that is passed into append. This is value semantics. The value is passed into `append.` As a strong rule of thumb, we use value semantics for builtin types and reference types (like slices). The `append` function then looks to see if `len` is equal to `cap`. If this is the case, `append` creates a new underlying array. Then `append` returns a slice and this slice is assigned to a new variable. Again, this is value semantics. The value is assigned to a new variable. The stack associated with `append` then self-cleans. The underlying array still exists. The slice assigned when `append` returned points to this underlying array.

# Memory leak scenarios
Memory leaks in Go can occur when there's a value on the heap and there's still some reference to it, thus not allowing the garbage collector to free up that memory. Causes of memory leaks in Go include:
1. goroutines - Goroutine leaking or blocking; the goroutine didn't terminate.
1. maps - You have to delete keys when entries are no longer needed.
1. append - The data going in is not being replaced on the way going out.
1. close - Forgetting to close a resource.

# Slicing a slice
- syntax [inclusive:exclusive:cap]

# Range clause

The `range` clause provides a convenient way to iterate over slices, arrays, maps, strings, and channels. When it comes to slices, the `range` clause allows you to easily loop over each element in the slice.

Let's break down the usage and characteristics of the `range` clause with slices:

1. **Basic Usage**:
   When you use `range` with a slice, it returns two values for each iteration:
   - The index of the current element.
   - A copy of the element at that index.

   Here's a simple example:
   ```go
   nums := []int{2, 3, 5, 7, 11, 13}
   for i, v := range nums {
       fmt.Printf("index %d: value %d\n", i, v)
   }
   ```

2. **Ignoring the Index or Value**:

   - To get only the index:
     ```go
     for i := range nums {
         fmt.Println(i)
     }
     ```
   - To get only the value:
     ```go
     // use the blank identifier `_` to ignore the unwanted value:
     for _, v := range nums {
         fmt.Println(v)
     }
     ```

3. **Working with a Copy**:
   It's crucial to understand that `range` provides a copy of the slice's element on each iteration. Therefore, modifying the value (`v` in the examples) does not modify the original slice. If you want to modify the original slice, you'd use the index to access and update the slice.

   ```go
   for i := range nums {
       nums[i] = nums[i] * 2
   }
   ```

4. **Range and Slice Headers**:
   When you're iterating over a slice using `range`, you're actually working with the slice's underlying array. This means that if you're also modifying the slice (e.g., appending elements) within the loop, you need to be careful because the behavior could be unexpected.

5. **Usage with Multi-dimensional Slices**:
   If you're working with a multi-dimensional slice (like a slice of slices), you can use nested loops to iterate through all the elements:

   ```go
   matrix := [][]int{
       {1, 2, 3},
       {4, 5, 6},
       {7, 8, 9},
   }
   
   for i, row := range matrix {
       for j, val := range row {
           fmt.Printf("element [%d][%d]: %d\n", i, j, val)
       }
   }
   ```

In summary, the `range` clause in Go provides a concise and readable way to iterate over slices. It gives you both the index and the value, but remember, the value is a copy, so you'll need to use the index if you wish to modify the original slice's elements.

# Code review check

### Rule of thumb: use value semantics to move around:
- builtin types
- reference types

### Use var for zero value
- don't user empty literal construction for zero value - don't do this:
```go
u := User{}         // zero value for user
xs := []string{}    // not a zero value; it's an empty slice 
```
- do this for zero value:
```go
var u User         // zero value for user
var xs []string    // zero value for slice of string 
```

### Preallocating cap
- As a rule of thumb, don't preallocate `cap` unless you know the size needed.
- Priorities: integrity first, resource usage second

### Append assigns to new variable
- Code smell: append to an array, then assign to a new variable
- This could result in holding references to old backing arrays the are no longer needed - memory leak.

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
