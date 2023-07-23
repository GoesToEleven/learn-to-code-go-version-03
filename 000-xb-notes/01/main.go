package main

import "fmt"

type person struct {
	name string
	age  int
}

// value semantics
func (p person) info() {
	fmt.Println("Name & age \t", p.name, p.age)
}

// DO NOT DO THIS: mixing semantics
// this is pointer semantics
// also, no setters getters in go
// this is just for example
func (p *person) setAge(age int) {
	p.age = age
}

func main() {

	p := person{
		name: "James",
	}
	p.info()
	p.setAge(32)
	fmt.Println("***********")
	// f1 gets its own copy of p because the method uses value semantics
	f1 := p.info
	f1()
	p.name = "Jenny"
	f1()     // prints James
	p.info() // prints Jenny
	fmt.Println("***********")

	// f2 gets the address of d because the method uses pointer semantics
	f2 := p.setAge
	f2(27)
	p.info()
	p.name = "Moneypenny"
	f2(27)
	p.info()
}

/*
The variable `f1` in your code, which refers to a method with a receiver, is a two-word data structure.

When you create a variable like `f1 := p.info` in Go, you're creating a "method value". This is a function that is "bound" to a specific receiver value, in this case `p`. The `f1` variable now holds the function `info` and the specific `p` it is associated with.

In the internal implementation, the method value is represented as a two-word object. The first word points to the code for the method, and the second word is a pointer or a copy of the receiver value, depending on whether the receiver is of a pointer type or a value type.

In the case of `f1 := p.info`, because `info` has a value receiver (not a pointer receiver), a copy of `p` is made when `f1` is created. This copy is stored in the second word of the internal data structure. This is why changes to `p` after the creation of `f1` don't affect the output of `f1()`.

Conversely, `f2 := p.setAge` is also a two-word object, but because `setAge` has a pointer receiver, the second word is a pointer to `p`. That means changes to `p` or through `f2` will affect the same person object.
*/

/*
When you have DOUBLE INDIRECTION as in the code above,

f1 is a two word data structure
the first word points to the code for the method
the second word points to a copy of the receiver value

f2 is a two word data structure
the first word points to the code for the method
the second word points to the data structure p

in the case of f2, because of DOUBLE INDIRECTION,
the data structure p ends up on the heap

TAKEAWAY
DECOUPLING YOUR CODE RESULTS IN ALLOCATIONS ON THE HEAP
this is not a big deal - we don't need the most blazing fast go program ever,
however, it is good to know the costs
decoupling can create productive allocations on the heap



*/
