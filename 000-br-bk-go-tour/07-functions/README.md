# "Why did you create a new language?"
"Go was born out of frustration with existing languages and environments for the work we were doing at Google. Programming had become too difficult and the choice of languages was partly to blame. One had to choose either ***efficient compilation, efficient execution, or ease of programming;*** all three were not available in the same mainstream language. Programmers who could were choosing ease over safety and efficiency by moving to dynamically typed languages such as Python and JavaScript rather than C++ or, to a lesser extent, Java.

We were not alone in our concerns. After many years with a pretty quiet landscape for programming languages, Go was among the first of several new languages—Rust, Elixir, Swift, and more—that have made programming language development an active, almost mainstream field again.

Go addressed these issues by attempting to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aimed to be modern, with support for networked and multicore computing. Finally, working with Go is intended to be fast: it should take at most a few seconds to build a large executable on a single computer. To meet these goals required addressing a number of linguistic issues: an expressive but lightweight type system; concurrency and garbage collection; rigid dependency specification; and so on. These cannot be addressed well by libraries or tools; a new language was called for."
[Golang FAQ](https://go.dev/doc/faq#creating_a_new_language)

# Video
[SOON TO BE PUBLISHED Here is the video on YouTube]()

# Takeaways
1. Functions and modularity
1. Parameters and arguments
1. Syntax: func, receiver, identifier, [parameter(s)], [return(s)], code block
1. No named returns
1. C
1. C
1. C

# Table of Contents

1. [Function overview](#function-overview)
1. [](#)
1. [](#)
1. [](#)
1. [](#)
1. [](#)
1. [](#)
1. [Code review check](#code-review-check)

# Function overview

Functions are blocks of code that perform a specific task or a set of related tasks. Functions are a fundamental building block of programs and are used to encapsulate and organize code for better modularity and reusability. Here's a detailed explanation of functions in Go:

1. Function Declaration:
   To declare a function in Go, you use the `func` keyword followed by the function name, a list of parameters enclosed in parentheses, and an optional return type. The general syntax is:

   ```go
   func functionName(parameter1 type, parameter2 type, ...) return_type {
       // Function body
   }
   ```

   For example:

   ```go
   func add(x int, y int) int {
       return x + y
   }
   ```

2. Parameters:
   Functions can take zero or more parameters, which are the values passed to the function when it is called. Each parameter consists of a name and a type. In the example above, `add` takes two integer parameters, `x` and `y`.

3. Return Values:
   Functions can return zero or one value (or none, in which case you can omit the return type). The `return` statement is used to specify the value to be returned. In the example above, `add` returns an integer.

4. Multiple Return Values:
   Go allows functions to return multiple values. This is useful in cases where you want to return more than one result from a function. For example:

   ```go
   func divide(x, y float64) (float64, error) {
       if y == 0 {
           return 0, errors.New("division by zero")
       }
       return x / y, nil
   }
   ```

   In this example, the `divide` function returns both a floating-point result and an error (if the division by zero occurs).

5. Named Return Values:
   Go allows you to specify names for the return values in a function's signature. Named return values are automatically initialized and can be used as regular variables within the function. They are particularly useful when dealing with complex functions or error handling. For example:

   ```go
   func divide(x, y float64) (result float64, err error) {
       if y == 0 {
           err = errors.New("division by zero")
           return
       }
       result = x / y
       return
   }
   ```

6. Function Calls:
   To call a function in Go, you use the function name followed by a list of arguments enclosed in parentheses. If the function has multiple return values, you can capture them in variables. For example:

   ```go
   sum := add(3, 5)           // Call the add function and assign its result to sum
   result, err := divide(8, 2) // Call the divide function and capture both result and error
   ```

7. Function Variadic Parameters:
   Go supports variadic functions, which allow you to pass a variable number of arguments to a function. To define a variadic parameter, you use an ellipsis (`...`) followed by the parameter type. For example:

   ```go
   func calculateSum(nums ...int) int {
       sum := 0
       for _, num := range nums {
           sum += num
       }
       return sum
   }
   ```

   You can call this function with any number of integers.

8. Anonymous Functions (Closures):
   Go supports anonymous functions, also known as closures. These are functions without a name, and they can be assigned to variables and used as arguments to other functions. Closures are often used in Go for tasks like defining inline functions or for concurrent programming using goroutines.

   ```go
   add := func(x, y int) int {
       return x + y
   }

   result := add(3, 5)
   ```

9. Function as a Type:
   In Go, functions can also be used as types. This allows you to define functions that take other functions as arguments or return them as results. It's a powerful feature for implementing higher-order functions and callbacks.

   ```go
   type MathFunc func(int, int) int

   func operate(x, y int, op MathFunc) int {
       return op(x, y)
   }
   ```

   You can then pass functions as arguments to `operate`.

10. Defer and Panic:
    Go provides two special built-in functions called `defer` and `panic` for handling exceptional situations and resource management. `defer` is used to schedule a function call to be executed just before the function returns, while `panic` is used to trigger a run-time error and unwind the stack.

Functions in Go play a central role in structuring code and promoting clean, maintainable, and efficient software. Understanding how to declare, define, and use functions is crucial for writing effective Go programs.

# Code review check

### No named returns
- named returns decrease readability

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# "The very essence of success is practice."
#### ~ Ignacy (Jan) Paderewski

![Ignacy (Jan) Paderewski](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/06-constants/image/ip.png)

source: https://en.wikipedia.org/wiki/Ignacy_Jan_Paderewski