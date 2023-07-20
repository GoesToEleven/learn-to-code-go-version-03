package main

import "fmt"

func main() {
	// Create a slice to store the questions
	var questions []string

	// Add questions to the slice
	questions = append(questions, "Put these in the correct order: centisecond, second, nanosecond, millisecond, decisecond, microsecond")
	questions = append(questions, "What is the difference between make([]int, 10), make([]int, 0, 10), make([]int, 10, 10)")
	questions = append(questions, "Explain these types in Go: builtin, reference, user defined")
	questions = append(questions, "Explain value semantics and pointer semantics. What are rules-of-thumb for using one versus the other?")
	questions = append(questions, "What types can a map use as a key in the Go programming language?")
	questions = append(questions, "Explain how a linked list works.")
	questions = append(questions, "What does it mean to write idiomatic Go code?")
	questions = append(questions, "Who is Rob Pike?")
	questions = append(questions, "Who is Ken Thompson?")
	questions = append(questions, "Who is Robert Griesemer")
	questions = append(questions, "What is the short-declaration operator and when do you use it?")
	questions = append(questions, "What is the var keyword used for and when do you use it?")
	questions = append(questions, "How are arrays and slices different in Go? What are rules-of-thumb around arrays and slices in Go?")
	questions = append(questions, "What are pointers?")
	questions = append(questions, "How does a computer work?")
	questions = append(questions, "What is iota?")
	questions = append(questions, "What is a Goroutine?")
	questions = append(questions, "What is the difference between concurrency and parallelism?")
	questions = append(questions, "What is a compiler?")
	questions = append(questions, "What is the stack?")
	questions = append(questions, "What is the heap?")
	questions = append(questions, "What is escape analysis?")
	questions = append(questions, "Which do you choose: performance or readability?")
	questions = append(questions, "Which do you choose: brevity or readability?")
	questions = append(questions, "What is a garbage collector?")
	questions = append(questions, "Is Go an Object Oriented language?")
	questions = append(questions, "Is Go a Data Oriented language?")
	questions = append(questions, "How does Go handle error handling compared to other languages?")
	questions = append(questions, "Why was Go created, and who created it?")
	questions = append(questions, "Can you explain what is meant by 'Go is strongly typed?'")
	questions = append(questions, "Can you give an example of a problem you've solved using Go? What challenges did you encounter, and how did you overcome them?")
	questions = append(questions, "Can you describe a time when you used channels and goroutines to solve a concurrency problem?")
	questions = append(questions, "Have you worked on a large-scale project using Go? If so, how did Go's features contribute to the project's success?")
	questions = append(questions, "What tools do you typically use for testing Go code?")
	questions = append(questions, "How would you set up a CI/CD pipeline for a Go project?")
	questions = append(questions, "Can you explain how package management works in Go?")
	questions = append(questions, `What's problematic with this code:

	func slice() {
		xi := make([]int, 10)
		for i := 0; i < 10; i++ {
			xi1 = append(xi1, i)
		}
		fmt.Println(xi1)
	}		
	`)
	questions = append(questions, "What is your IDE for Go, and why?")
	questions = append(questions, "Can you argue in favor of your perspective, and accept when your perspective isn't chosen?")
	questions = append(questions, "What do you do to stay current with programming?")
	questions = append(questions, `Do you recognize any of these, and if so, what are they and have you worked with them:
	cobra
	viper
	opentelemetry
	jaeger
	docker
	chi
	zshell
	`)
	questions = append(questions, "What is a multiplexer?")
	questions = append(questions, "What is the difference between TCP and UDP?")
	questions = append(questions, "Who is Caleb Doxsey?")
	questions = append(questions, "Tell me about containers and container orchestration.")
	questions = append(questions, "What made Docker innovative?")
	questions = append(questions, "Tell us about your soft-skills.")

	// Print out each question
	for i, question := range questions {
		fmt.Printf("Question %d: %s\n", i+1, question)
	}

	sliceDiff()
}

func sliceDiff() {
	xi1 := make([]int, 10)
	xi2 := make([]int, 0, 10)
	xi3 := make([]int, 10, 10)

	for i := 0; i < 10; i++ {
		xi1 = append(xi1, i)
		xi2 = append(xi2, i)
		xi3 = append(xi3, i)
	}
	fmt.Println(xi1)
	fmt.Println(xi2)
	fmt.Println(xi3)

}
