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
    - [video soon to be published]
1. [Struct, padding bytes, methods sets](/000-br-bk-go-tour/04a-struct-types/) 
    - [video soon to be published]
1. [Pointers, nil, stack, heap, escape analysis](/000-br-bk-go-tour/05-pointers) 
    - [video soon to be published]
1. [Typed and untyped constants, iota, and math big](/000-br-bk-go-tour/06-constants) 
    - [video soon to be published]
1. [functions, types, anonymous funcs, closure](/000-br-bk-go-tour/07-functions) 
    - [video soon to be published]
1. [Coupons for Go courses](coupons-for-go-courses)

# Coupons for Go courses
1. [Bill Kennedy courses](https://courses.ardanlabs.com/order?ct=670e0200-1823-4916-8ff5-b2438450e2ce) 
    - coupon: toddmcleod
1. [Todd McLeod courses](https://www.udemy.com/course/learn-how-to-code/?referralCode=BE659D12A78B2C0DFFB0)
    - click on my user icon for more courses

# Code review check

[readme: Strings, unicode, UTF-8](/000-br-bk-go-tour/01-string-unicode-utf8/)
- Code review guidelines
    - Integrity
    - Readability
    - Simplicity
    - Performance

[readme: Variables, values, & types](/000-br-bk-go-tour/03-variables-01/)
- int
    - using anything other than int
    - Why are you using an int32 or int64; why not just int?

[readme: Struct, padding bytes, methods sets](/000-br-bk-go-tour/04a-struct-types/)
- struct fields largest to smallest
    - Why are you micro-optimizing?
    - Optimize for readability first
- empty literal struct
    - use var instead: var p person

[readme: Pointers, nil, stack, heap, escape analysis](/000-br-bk-go-tour/05-pointers)
- Returning a pointer
    - If you are returning a pointer, make it readable:

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

[readme: Typed and untyped constants, iota, and math big](/000-br-bk-go-tour/06-constants)
- typed constants decrease precision and flexibility
    - only use typed constants if you have a specific reason

[readme: functions, types, anonymous funcs, closure](/000-br-bk-go-tour/07-functions)
- named returns decrease readability
    - an empty `return` is not idiomatic