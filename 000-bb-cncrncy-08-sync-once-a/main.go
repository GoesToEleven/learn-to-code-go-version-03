package main

import (
	"fmt"
	"sync"
)

type Payment struct {
	From   string
	To     string
	Amount float64

	once sync.Once
}

func (p *Payment) process() {
	fmt.Printf("%s --> $%.2f --> %s\n", p.From, p.Amount, p.To)
}

func (p *Payment) Process() {
	p.once.Do(p.process)
}

func main() {
	p := Payment{
		From:   "Bugs Bunny",
		To:     "Walt",
		Amount: 43.78,
	}
	p.Process()
	p.Process() // indempotence
}

/*
In programming, indempotence refers to a property or characteristic of an operation or function where applying the operation multiple times produces the same result as applying it once. In other words, if you repeat an idempotent operation or function with the same input, you will get the same output without any additional side effects or changes to the system.

Idempotence is an important concept in various areas of computer science, including distributed systems, web development, and network protocols. It allows developers to design systems that can safely handle repeated requests or retries without causing unintended consequences or inconsistencies.

To illustrate this concept, consider an API endpoint that creates a new user account. If the creation of a user account is an idempotent operation, calling the endpoint multiple times with the same set of data will result in creating a single user account. Subsequent requests with the same data will have no effect, and the system will remain in the same state.

Idempotence is especially crucial in scenarios where failures or network issues can cause duplicate requests. By ensuring that operations are idempotent, developers can build robust and reliable systems that can tolerate these situations without producing incorrect results or unintended side effects.

It's important to note that not all operations or functions are idempotent. Some operations may have side effects or produce different outcomes when applied multiple times, such as updating a counter or generating a unique identifier. Therefore, understanding whether an operation is idempotent or not is essential for designing and implementing reliable and predictable systems.
*/


