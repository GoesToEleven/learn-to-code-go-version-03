package tokenizer

import "fmt"

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := Tokenize(text)
	fmt.Println(tokens)

	//Output:
	// [who s on first]
}