# Table of Contents

1. [Composite Types & Aggregate Types](# Composite Types & Aggregate Types)

# Composite Types & Aggregate Types

In the context of the Go programming language, the terms "composite types" and "aggregate types" are often used interchangeably. However, it's worth breaking down the nuances of these two terms to clarify the distinction.

### Composite Types in Go

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

### Summary

In summary, while all aggregate types (in the general sense) could be considered composite types in Go, not all composite types would traditionally be considered aggregate types. For example, channels and maps in Go could be considered composite but not aggregate, as they don't just bundle existing types together—they provide additional functionality and behavior.

In practice, the terms are often used interchangeably, and the specific meaning is usually clear from the context in which they are used.