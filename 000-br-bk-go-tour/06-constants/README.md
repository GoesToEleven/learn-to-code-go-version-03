
# "If I miss one day of practice, I notice it. If I miss two days, the critics notice it. If I miss three days, the audience notices it."
~ Ignacy (Jan) Paderewski, pianist and politician, he signed the Treaty of Versailles, which ended World War I.

### "Implicit conversion is one of the features that make Go both strongly typed and convenient to use."

#### "An UNTYPED CONSTANT has a DEFAULT TYPE which is the type to which the constant is implicitly converted in contexts where a typed value is required - for instance, in a short variable declaration such as i := 0 where there is no explicit type. The default type of an untyped constant is bool, rune, int, float64, complex128, or string respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant."
~ [Go language specification](https://go.dev/ref/spec#Constants)


# Video
[SOON TO BE PUBLISHED Here is the video on youtube]()

# Takeaways
1. Constants can be typed and untyped
1. Untyped constants are also called "constants of a kind"
1. Untyped constants can be implicitly converted at compile time
1. Numeric literals are untyped constants
1. Typed constants limit precision
1. Scientific notation, mantissa, and e

# Table of Contents

1. [Mechanics](#mechanics)
1. [Understanding constants](#understanding-constants)
1. [Understanding "kind"](#understanding-kind)
1. [Numeric literals are untyped constants](#numeric-literals-are-untyped-constants)
1. [Constants have high precision](#constants-have-high-precision)
1. [Typed constants limit precision](#typed-constants-limit-precision)
1. [Scientific notation mantissa and e](#scientific-notation-mantissa-and-e)
1. [Go big with math big](#go-big-with-math-big)
1. [Width and precision](#width-and-precision)
1. [Arbitrary precision and math big](#arbitrary-precision-and-math-big)
1. [Code review check](#code-review-check)

# Terminology

1. Constant of a kind 	(untyped constant)
1. Constant of a type 	(typed constant)
1. Kind promotion 		(untyped --> typed)

# Mechanics

The Go language specification says, "There are boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants." Below are examples of each type of constant:

### 1. Boolean Constants
Boolean constants can be either `true` or `false`.

```go
const myTrue = true
const myFalse = false
```

### 2. Rune Constants
Rune constants represent a single Unicode code point enclosed in single quotes.

```go
const letterA = 'A'
const unicodeStar = '★'
```

### 3. Integer Constants
Integer constants are just numbers without fractional components. They can be positive, negative, or zero.

```go
const zero = 0
const positiveInt = 42
const negativeInt = -42
```

### 4. Floating-Point Constants
Floating-point constants are numbers that have a decimal point or are expressed in scientific notation.

```go
const pi = 3.14159
const avogadro = 6.022e23
```

### 5. Complex Constants
Complex constants are formed by adding a real and an imaginary component. The imaginary component ends in `i`.

```go
const complexNum = 1 + 2i
const anotherComplexNum = 3.14 - 2.7i
```

### 6. String Constants
String constants are sequences of characters enclosed in double quotes.

```go
const helloWorld = "Hello, world!"
const anotherString = "Go is fun."
```

Here they all are together in a complete Go program for context:

```go
package main

import "fmt"

func main() {
	// Boolean constant
	const myTrue = true

	// Rune constant
	const letterA = 'A'

	// Integer constants
	const zero = 0
	const positiveInt = 42

	// Floating-point constants
	const pi = 3.14159

	// Complex constants
	const complexNum = 1 + 2i

	// String constants
	const helloWorld = "Hello, world!"

	fmt.Println("Boolean:", myTrue)
	fmt.Println("Rune:", letterA)
	fmt.Println("Integer:", zero, positiveInt)
	fmt.Println("Floating-point:", pi)
	fmt.Println("Complex:", complexNum)
	fmt.Println("String:", helloWorld)
}
```

Each of these constants has a type that's inferred from its value, except that they are "untyped" until they are used in a context that requires a specific type. This allows them to be used more flexibly than variables.

# Understanding constants

A constant is a simple, unchanging value that is determined at compile time. Constants are created using the `const` keyword and can be of any basic data type, such as integer, float, string, or boolean.

Here's how you can define a constant:

```go
const Pi = 3.14159
```

In this example, `Pi` is a constant whose value is `3.14159`. Once a constant is defined, you can't modify its value. Attempting to do so will result in a compile-time error.

You can also define multiple constants in a single declaration:

```go
const (
  FirstName = "John"
  LastName  = "Doe"
)
```

### Typed and Untyped Constants

In Go, constants can be either typed or untyped.

- **Typed constants**: These have an explicit type specified. For example, you could specify that a constant should be of type `float64`:

  ```go
  const Pi float64 = 3.14159
  ```

  Typed constants can only be used where that specific type is expected.

- **Untyped constants**: These do not have an explicit type specified and can adapt their type based on the context in which they are used. For example:

  ```go
  const Pi = 3.14159
  ```

  Here, `Pi` is an untyped constant and can be used wherever a numeric type is expected.

### Constant Expressions

Constants can also be determined through expressions, but remember that these expressions are evaluated at compile-time:

```go
const TwoPi = Pi * 2
```

However, the expressions used to initialize a constant must be made up of either literals, previously defined constants, or built-in functions. They can't be determined at runtime or include variables, function calls (unless the function is a built-in), or complex computations.

### Enumeration Constants

You can use constants to create enumerations. The `iota` keyword is often used for this purpose. It starts at 0 and increments by 1 for each item in the constant list:

```go
const (
  Sunday = iota  // 0
  Monday         // 1
  Tuesday        // 2
  // ...
)
```

### Usage

Constants are generally used to make the code more readable and to avoid "magic numbers" or repeated values.

```go
const SecondsInMinute = 60

func calculateTotalSeconds(minutes int) int {
  return minutes * SecondsInMinute
}
```

By defining `SecondsInMinute` as a constant, the code becomes easier to read and maintain. If the value needs to change for some reason, you only have to update it in one place.

# Understanding kind

The concept of "kind" isn't explicitly defined in the language specification, but the notion is sometimes used informally to describe the underlying category of a type. For example, both `int` and `int32` are of the integer "kind." In this context, when you define an untyped constant, it's not just a value but a value with a certain "kind" based on its syntax and how it's used.

An untyped constant has a default type that can be implicitly used if the context demands it. However, unlike a typed constant, an untyped constant is flexible and can be implicitly converted to other types within its "kind."

For example:

```go
const x = 42 // untyped integer constant, "kind" integer
const y = 3.14 // untyped float constant, "kind" float

var i int = x // valid, x is of "kind" integer
var f float64 = y // valid, y is of "kind" float

var j int32 = x // also valid, x can be converted to int32
var g float32 = y // also valid, y can be converted to float32
```

In these examples, `x` and `y` are untyped constants of different kinds—integer and float, respectively. They can be assigned to variables of any type that matches their "kind."

The concept of "kind" is more explicitly modeled in Go's reflection package (`reflect`). The `reflect.Kind` type is an enumeration that categorizes the kinds of Go types, like `Int`, `Float`, `String`, etc. However, note that "kind" in the reflection package has a broader scope and encompasses more than just basic types.

So, when talking about "constants of a kind," you're generally referring to untyped constants that can adapt their type based on the context in which they are used, fitting into a broad category of similar types.

# Numeric literals are untyped constants

Numbers like `3` or `8` are called NUMERIC LITERALS - they are also UNTYPED CONSTANTS. 

Numeric literals are untyped constants ("constants of a kind"). This means that they do not have a fixed type until they are assigned to a variable or used in an expression that requires them to take on a specific type.

The compiler can implicitly convert these untyped constants to appropriate types depending on the context in which they are used. This is different from typed constants, which have a fixed type and cannot be implicitly converted to another type.

To demonstrate implicit type conversion, consider these examples:

1. Implicit typing of integer literals:

    ```go
    var i int = 30 // Here, 30 is an untyped integer literal implicitly converted to type int
    ```

2. Implicit typing of float literals:

    ```go
    var f float64 = 2.5 // Here, 2.5 is an untyped float literal implicitly converted to type float64
    ```

3. Mixing types in expressions:

    ```go
    var result float64 = 30 / 2.5 // Here, 30 is an untyped integer literal, 2.5 is an untyped float literal
                                  // Both are implicitly converted to float64 for the expression
    ```

4. Short variable declarations:

    ```go
    x := 30  // The type of x is inferred to be int
    y := 2.5 // The type of y is inferred to be float64
    ```

5. Function arguments:

    ```go
    func foo(x float64) {
        //...
    }

    foo(30)  // Here, 30 is an untyped integer literal implicitly converted to float64
    ```

In all these examples, the untyped numeric literals (`30`, `2.5`) are implicitly converted to appropriate types depending on the context in which they are used.

***Implicit conversion is one of the features that make Go both strongly typed and yet convenient to use***. It allows for a form of polymorphism at the level of basic types, where an untyped constant can be used wherever any type within its "kind" (integer, float, etc.) is expected.

# Constants have high precision

The high-precision arithmetic for constants in Go enhances code robustness and reliability in a few ways:

### 1. Compile-time Checking

The high precision of constant arithmetic allows for exact results during **compile-time evaluation.** This can catch potential errors, overflow, and underflow conditions early. This is beneficial for numerical algorithms and systems programming where precision is crucial.

```go
const val = 1 << 60 // No overflow, even though it exceeds an int32 or int64
var x int64 = val   // Compile-time error: constant 1152921504606846976 overflows int64
```

In this example, the Go compiler performs a compile-time check to make sure that the constant value fits within the `int64` data type and flags an error if it doesn't.

### 2. Simplified Code

Because constants can be computed with high precision, it can simplify code by allowing complex calculations to be **pre-computed at compile-time.**

```go
const PiSquared = 3.141592653589793 * 3.141592653589793
```

In this example, `PiSquared` is calculated at compile-time with high precision. There's no need for run-time calculation or approximation.

### 3. Type Safety and Inference

Go constants are untyped until they are explicitly given a type either through assignment or by being an operand in an expression requiring a type. This means that you can define a constant in one part of the code and use it in multiple places where different types are expected, without losing precision.

```go
const val = 1 << 30
var a int32 = val  // No problem: fits within int32
var b int64 = val  // No problem: fits within int64, and no loss of precision
```

In these examples, the same constant `val` can be used in different contexts without manual type conversions, making the code cleaner and less error-prone.

### 4. Enabling Advanced Algorithms

Certain algorithms require high precision, especially in fields like cryptography, scientific computing, and data analysis. While the high-precision arithmetic of Go is mostly used at compile-time for constants, having this capability can enable certain optimizations and guarantees about the behavior of the code.

```go
const Prime = 2305843009213693951  // A Mersenne prime number (2^61 - 1)
```

In this example, a large Mersenne prime can be precisely represented as a constant, useful in cryptographic algorithms.

### 5. Portability

Having constants with high precision reduces architecture-specific discrepancies, which is critical in systems programming. Code behavior remains consistent whether it's running on a 32-bit or 64-bit system, as long as you're relying on constants for certain calculations.

```go
const Mask = 1 << 35  // Same value, regardless of 32-bit or 64-bit architecture
```

In this example, `Mask` will have the same value irrespective of the underlying architecture.

These features contribute to making Go a more robust and reliable language for numerical computations and systems programming.

# Typed constants limit precision

In Go, typed constants are subject to the same size and precision limitations as variables of the same type. When you explicitly specify the type of a constant, you're basically telling the compiler to limit the constant's precision to what can be held by that type. This is in contrast to untyped constants, which have arbitrary precision until they're assigned to a variable or used in an expression that requires a type.

Here's an example to illustrate:

```go
const untypedConst = 1 << 100  // No type specified, so arbitrary precision is used
const typedConst int64 = 1 << 60  // Explicitly typed as int64
```

In this example:

- `untypedConst` can hold a large value because it is untyped and thus can have arbitrary precision up to what the Go language specification allows (at least 256 bits for integers).
  
- `typedConst`, however, is explicitly typed as `int64`, which means it can only hold values that fit within 64 bits. If you try to set it to a value that can't be held by an `int64`, you'll get a compile-time error.

Here's another example with floating-point numbers:

```go
const untypedFloat = 3.14159265358979323846264338327950288419716939937510582097494459
const typedFloat float64 = 3.141592653589793
```

- `untypedFloat` is untyped and therefore can hold more digits of precision than any built-in floating-point type. The Go language specification requires at least 256 bits of precision for the mantissa of floating-point constants.

- `typedFloat` is explicitly typed as `float64`, which means it has 53 bits of mantissa, so the precision is limited to what can be represented by a `float64`.

If you try to assign these constants to variables, the typed constants will maintain their precision limitations:

```go
var a = untypedConst  // Error if 'a' can't hold the value of 'untypedConst'
var b int64 = typedConst  // No error, fits within int64
```

By limiting the precision of a constant through typing, you could inadvertently introduce errors or limitations into your calculations, especially if those constants are used in mathematical expressions where high precision is required. Therefore, you generally only want to give a constant a type if you're sure that the limited precision is acceptable for your use case.

# Scientific notation mantissa and e

Scientific notation writes numbers like this:

- 3.45 * 10^2
- 3.45e2

### Mantissa

The mantissa is a component of a number in scientific notation or a floating-point number, used along with the exponent to represent the number. Scientific notation represents numbers as a base number and an exponent. In the case of decimal scientific notation, for example, the number (3.45 * 10^2) has a mantissa of 3.45 and an exponent of 2.

- The **mantissa** holds the significant digits of the number. Note that the term "mantissa" in the context of floating-point numbers is sometimes also called "significand."
- The **exponent** scales the number by a power of 2.

For example, in IEEE 754 floating-point standard for base-2 (binary) numbers:

- For single-precision (float32 in Go), the mantissa is 23 bits.
- For double-precision (float64 in Go), the mantissa is 52 bits.

The mantissa determines the precision of the number: the more bits in the mantissa, the more precise the number can be. However, the number of bits in the mantissa is fixed for each floating-point type (float32, float64, etc.), which puts a limitation on how precisely numbers can be represented.

In the context of the Go programming language, when the specification mentions that ***floating-point constants*** should be represented with a mantissa of "at least 256 bits," it means that the constants can hold a very high degree of precision, much higher than the float32 or float64 types, until they are assigned to a variable or used in an expression that requires a specific type.

### e notation

The "e" notation is a shorthand for expressing numbers in scientific notation. It is often used for representing very large or very small numbers more concisely. In this notation, the letter "e" signifies "times ten raised to the power of."

Here's how it works:

- (a * 10^b) is represented as (aeb)
  
For example:

- ( 3 * 10^4) would be written as ( 3e4 )
- ( 2.5 * 10^-3) would be written as ( 2.5e-3 )

The number before the "e" is the mantissa, and the number after the "e" is the exponent. The exponent indicates the power of 10 by which the mantissa should be multiplied.

Here are some examples in Go code:

```go
const avogadro := 6.022e23  // 6.022 × 10^23
const electronCharge := 1.602e-19  // 1.602 × 10^-19
```

The "e" notation can be used for both float32 and float64 types in Go, and the compiler will infer the type based on the context in which the constant is used, or you can specify the type explicitly.

# Go big with math big

The `math/big` package in Go provides arbitrary-precision arithmetic for integers, rational numbers, and floating-point numbers. This package is particularly useful when you need to perform calculations that exceed the limits of Go's built-in numeric types.

Here's a quick rundown of how to use the `math/big` package:

### Integer (`big.Int`)

#### Initialization

You can initialize a `big.Int` variable in several ways:

- Using the `NewInt` function
- Using `SetInt64` or `SetUint64`
- Using `SetString`

```go
x := big.NewInt(42)
y := new(big.Int).SetInt64(42)
z := new(big.Int).SetString("42", 10) // Base 10
```

#### Arithmetic Operations

The `big.Int` type provides methods for various arithmetic operations:

```go
// Addition
result := new(big.Int).Add(x, y)

// Subtraction
result = new(big.Int).Sub(x, y)

// Multiplication
result = new(big.Int).Mul(x, y)

// Division
result = new(big.Int).Div(x, y)
```

### Floating-Point (`big.Float`)

#### Initialization

Initialize `big.Float` similar to `big.Int`:

```go
f1 := big.NewFloat(3.14)
f2 := new(big.Float).SetFloat64(3.14)
```

#### Arithmetic Operations

```go
// Addition
resultFloat := new(big.Float).Add(f1, f2)

// Subtraction
resultFloat = new(big.Float).Sub(f1, f2)

// Multiplication
resultFloat = new(big.Float).Mul(f1, f2)

// Division
resultFloat = new(big.Float).Quo(f1, f2)
```

### Rational (`big.Rat`)

`big.Rat` can be used for rational numbers (fractions).

#### Initialization

```go
r := big.NewRat(22, 7)
```

#### Arithmetic Operations

```go
// Addition, Subtraction, Multiplication, and Division
resultRat := new(big.Rat).Add(r, r)
resultRat = new(big.Rat).Sub(r, r)
resultRat = new(big.Rat).Mul(r, r)
resultRat = new(big.Rat).Quo(r, r)
```

### Comparison

The `Cmp` method can be used to compare two big numbers:

```go
if x.Cmp(y) == 0 {
    fmt.Println("x and y are equal")
}
```

### Conversion

You can convert `big.Int`, `big.Float`, and `big.Rat` to standard Go types, but be careful of overflow and loss of precision:

```go
intValue := x.Int64()
floatValue, _ := f1.Float64()
```

This is a very basic overview, and the package provides a rich set of methods to work with big numbers. You can read the [official package documentation](https://golang.org/pkg/math/big/) for more details.

### Rational numbers, aka, fractions

Rational numbers are numbers that can be expressed as the quotient or fraction a/b  of two integers a and b, where b neq 0. In other words, a rational number is one that can be written as a simple fraction. Both the numerator a and the denominator b must be integers, and the denominator b cannot be zero.

### Examples of Rational Numbers:

1. frac{1}{2} — One-half is a rational number because it can be expressed as a fraction of two integers.
2. frac{5}{1} — Five is a rational number because 5 = frac{5}{1} .
3. frac{-3}{4} — Negative numbers can be rational too. Here, -frac{3}{4} is a rational number.
4. frac{0}{1} — Zero is a rational number because it can be expressed as frac{0}{1} .
5. 7 — Whole numbers are also rational numbers because they can be written as frac{7}{1} , frac{6}{1} , frac{5}{1} , etc.
6. frac{22}{7} — This is an approximation of pi and is rational, although pi itself is not rational.

### Examples of Numbers that are NOT Rational:

1. pi — Pi cannot be exactly expressed as a simple fraction, so it's not a rational number.
2. sqrt{2} — The square root of 2 is an irrational number because it cannot be exactly expressed as a fraction.
3. e — Euler's number is also irrational.

### In Go with `math/big`

If you want to represent rational numbers in Go, you can use the `big.Rat` type from the `math/big` package. Here is an example:

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Create new rational numbers
	r1 := big.NewRat(22, 7)
	r2 := big.NewRat(-3, 4)

	// Add r1 and r2
	sum := new(big.Rat).Add(r1, r2)

	// Print the rational numbers
	fmt.Println("r1:", r1.FloatString(5))  // 5 digits after the decimal point
	fmt.Println("r2:", r2.FloatString(5))
	fmt.Println("sum:", sum.FloatString(5))
}
```

In this example, we created two rational numbers frac{22}{7}  and -frac{3}{4}  and then added them together. The `FloatString` method is used to convert the rational number to a string with a specified number of digits after the decimal point.

# Width and precision

With floats, you can specify the precision you want using the `%f` format specifier with a number indicating the number of digits after the decimal point.

For example, you could use `%.15f` to print up to 15 significant digits:

```go
package main

import (
	"fmt"
)

func main() {
	x := 1.0 / 3.0
	fmt.Printf("%.15f\n", x)
}
```

You might also consider using the `strconv.FormatFloat` function from the standard library, which allows more control over the format:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 1.0 / 3.0
	s := strconv.FormatFloat(x, 'f', 15, 64)
	fmt.Println(s)
}
```

Both of these methods are limited by the inherent precision of `float64`. If you need more precision, you could use arbitrary-precision arithmetic like the `math/big` package's `big.Float` type, but keep in mind that this comes with a performance cost.

# Arbitrary precision and math big

The term "arbitrary precision" refers to the capability of handling numbers with an arbitrary number of digits, limited only by the available system memory, rather than a fixed number of digits that would be determined by the data type size. This is in contrast to standard fixed-precision numeric types like integers (`int`, `int64`, etc.) or floating-point numbers (`float32`, `float64`) in most programming languages, which have a fixed range and level of precision.

In the context of programming, arbitrary-precision arithmetic is often implemented using specialized software libraries that can handle very large integers or very accurate floating-point numbers. These are useful for applications requiring very high precision or the ability to handle very large numbers, such as cryptography, scientific computing, financial applications, and more.

For example, in Go, the `math/big` package provides arbitrary-precision arithmetic for integers (`big.Int`), floating-point numbers (`big.Float`), and rational numbers (`big.Rat`).

Here is a simple Go example using `big.Int` to calculate factorial of 100:

```go
package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func main() {
	fmt.Println(factorial(100))
}
```

The output will be a very large integer that couldn't be handled by Go's built-in integer types. Similarly, `big.Float` can be used to perform high-precision floating-point calculations.

With arbitrary precision, you can continue to get more and more accurate results (or handle larger and larger numbers) as long as you have the memory to store the calculations. However, arbitrary-precision arithmetic is typically much slower than fixed-precision arithmetic, due to the added complexity of handling numbers with a dynamic number of digits.

# Code review check

### Use untyped constants
- typed constants decrease precision and flexibility
    - only use typed constants if you have a specific reason

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# "The very essence of success is practice."
#### ~ Ignacy (Jan) Paderewski

![Ignacy (Jan) Paderewski](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/06-constants/image/ip.png)

source: https://en.wikipedia.org/wiki/Ignacy_Jan_Paderewski