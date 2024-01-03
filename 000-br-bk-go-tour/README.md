# Welcome

This repo and [the videos on youtube](https://www.youtube.com/playlist?list=PLSak_q1UXfPqSyH2r5DnCXUJKIlbrLVGn) are here to help you learn the Go programming language!

# Table of Contents
1. [Strings, unicode, UTF-8](/000-br-bk-go-tour/01-string-unicode-utf8/) 
    - [video on youtube](https://www.youtube.com/watch?v=S3BHZv6OrJg)
1. [Time package & time formatting](/000-br-bk-go-tour/02-time-pkg/) 
    - [video #1 on youtube](https://youtu.be/HBtu9Dsjp80)
    - [video #2 on youtube](https://youtu.be/ut_REn0xFPM) 
1. [Git & GitHub: Creating Your First Pull Request](/000-br-bk-go-tour/02b-git-github/Todd-McLeod-Course-Learn-Git-GitHub-Beginner-Intermediate-Concepts.pdf)
    - [video on youtube](https://youtu.be/VZeOcX2DPwo) 
1. [Design Guidelines & Code Reviews](/000-br-bk-go-tour/02c-design-guides)
    - [video on youtube](https://youtu.be/WkQFrctSDsc) 
1. [Variables, values, & types](/000-br-bk-go-tour/03-variables/) 
    - [video #5 on youtube](https://www.youtube.com/watch?v=AmiomWHD6F8) 
    - [video #6 on youtube](https://youtu.be/wpS3T83IRfc) 
1. [Struct, padding bytes, methods sets](/000-br-bk-go-tour/04a-struct-types/) 
    - [video #7 on youtube](https://youtu.be/sKLhcphr40E) 
    - [video #8 on youtube](https://youtu.be/yJeQ6OyU8dc) 
1. [Pointers, nil, stack, heap, escape analysis](/000-br-bk-go-tour/05-pointers) 
    - [video #9 on youtube](https://youtu.be/QKHLjA6e7w8) 
    - [video #10 on youtube](https://youtu.be/3Z9jEtDYa8I) 
    - [video #11 on youtube](https://youtu.be/SXFfUGN8Maw) 
1. [Typed and untyped constants, iota, and math big](/000-br-bk-go-tour/06-constants) 
    - [video soon to be published]
    - [video # on youtube]() 
1. [functions, types, anonymous funcs, closure](/000-br-bk-go-tour/07-functions) 
    - [video soon to be published]
1. [Arrays, for range clause, semantics](/000-br-bk-go-tour/08-arrays) 
    - [video soon to be published]
1. [Slices](/000-br-bk-go-tour/09-slices) 
    - [video soon to be published]
1. [Coupons for Go courses](coupons-for-go-courses)

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce) 
    - coupon: toddmcleod
1. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# Code review check

### Code review guidelines
- Integrity
- Readability
- Simplicity
- Performance
[readme: Strings, unicode, UTF-8](/000-br-bk-go-tour/01-string-unicode-utf8/)

### int
- using anything other than int
- Why are you using an int32 or int64; why not just int?
[readme: Variables, values, & types](/000-br-bk-go-tour/03-variables-01/)

### struct fields largest to smallest
- Why are you micro-optimizing?
- Optimize for readability first
### empty literal struct
- use var instead: var p person
[readme: Struct, padding bytes, methods sets](/000-br-bk-go-tour/04a-struct-types/)

### Returning a pointer
- If you are returning a pointer, make it readable:
[readme: Pointers, nil, stack, heap, escape analysis](/000-br-bk-go-tour/05-pointers)

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

### typed constants decrease precision and flexibility
- only use typed constants if you have a specific reason
[readme: Typed and untyped constants, iota, and math big](/000-br-bk-go-tour/06-constants)

### named returns decrease readability
- an empty `return` is not idiomatic
[readme: functions, types, anonymous funcs, closure](/000-br-bk-go-tour/07-functions)

### Use value semantics for builtin types
- Builtin types
    - bool
    - numerics
    - strings
- Exception: 
    - use pointers to represent "null", eg, no value, eg, a database not even returning a field
    - you can have a pointer to an int which is nil, because nil is the zero value of a pointer; whereas if you just had an int, you would have the zero value of zero.
[Arrays, for range clause, semantics](/000-br-bk-go-tour/08-arrays)

### Three types of latencies
- internal latency
- external latency
- data latency
 
### Clarity, readability, and abstraction
- Eliminate needless abstraction. Focus on readability. Write code that can be understood by average developers.
    - "I've seen go code where the author has you going down a rabbit hole of functions calling other functions, those functions are in different files, and each of those functions are only five to ten lines, and are only used once... If it's only used once - why create a new function? Very hard to follow that code. It seems like they think there is a size limit, that a function that has more than 10 lines in it is bad." ~ Richard M.

### Clarity, readability, and nesting
- "One of my largest concerns for design is nesting. Nesting with for loops, if else if, switch statements. When I see at the end of a chunk of code five levels of curly braces -- I need to simplify that. Another is nested functions." ~ Richard M. 

[Design Guidelines & Code Reviews](/000-br-bk-go-tour/02c-design-guides)