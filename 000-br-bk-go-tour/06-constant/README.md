
# "If I miss one day of practice, I notice it. If I miss two days, the critics notice it. If I miss three days, the audience notices it"
~ Ignacy (Jan) Paderewski, a French philosopher and mathematician. 

## "An untyped constant has a default type which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a short variable declaration such as i := 0 where there is no explicit type. The default type of an untyped constant is bool, rune, int, float64, complex128, or string respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant."
~ [Go language specification](https://go.dev/ref/spec#Constants)

## "Implicit conversion is one of the features that make Go both strongly typed and yet convenient to use."

# Video
[SOON TO BE PUBLISHED Here is the video on YouTube]()

# Takeaways
1. Constants can be typed and untyped
1. Untyped constants are also called "constants of a kind"
1. Untyped constants can be implicitly converted at compile time
1. Numeric literals are untyped constants
1. Typed constants limit precision

# Table of Contents

1. [Mechanics](#mechanics)
1. [Understanding constants](#understanding-constants)
1. [Understanding "kind"](#understanding-kind)
1. [Numeric literals are untyped constants](#numeric-literals-are-untyped-constants)
1. [Constants have high precision](#constants-have-high-precision)
1. [Typed constants limit precision](#typed-constants-limit-precision)
1. [Code review check](#code-review-check)

# Terminology

1. Constant of a kind
1. Constant of a type
1. Kind promotion

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

Numbers like `3` or `8` are called numeric literals. 

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

### Mantissa

The mantissa is a component of a number in scientific notation or a floating-point number, used along with the exponent to represent the number. Scientific notation represents numbers as a base number and an exponent. In the case of decimal scientific notation, for example, the number \(3.45 \times 10^2\) has a mantissa of \(3.45\) and an exponent of \(2\).

In the context of floating-point representation in computers, a number \(x\) is typically represented as:

\[
x = (-1)^{\text{sign}} \times 1.\text{mantissa} \times 2^{\text{exponent}}
\]

- The **sign** determines whether the number is positive or negative.
- The **mantissa** holds the significant digits of the number. Note that the term "mantissa" in the context of floating-point numbers is sometimes also called "significand."
- The **exponent** scales the number by a power of 2.

For example, in IEEE 754 floating-point standard for base-2 (binary) numbers:

- For single-precision (float32 in Go), the mantissa is 23 bits.
- For double-precision (float64 in Go), the mantissa is 52 bits.

The mantissa determines the precision of the number: the more bits in the mantissa, the more precise the number can be. However, the number of bits in the mantissa is fixed for each floating-point type (float32, float64, etc.), which puts a limitation on how precisely numbers can be represented.

In the context of the Go programming language, when the specification mentions that floating-point constants should be represented with a mantissa of "at least 256 bits," it means that the constants can hold a very high degree of precision, much higher than the typical float32 or float64 types, until they are assigned to a variable or used in an expression that requires a specific type.

# Code review check

### R
- If 


# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce)  
    - coupon: toddmcleod
2. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# "The very essence of success is practice."
#### ~ Ignacy (Jan) Paderewski

![Ignacy (Jan) Paderewski](https://github.com/GoesToEleven/learn-to-code-go-version-03/blob/main/000-br-bk-go-tour/06-garbage-collector/images/Ignacy_Jan_Paderewski_-_Project_Gutenberg_eText_15604.png)

source: https://upload.wikimedia.org/wikipedia/commons/d/d4/Ignacy_Jan_Paderewski_-_Project_Gutenberg_eText_15604.png