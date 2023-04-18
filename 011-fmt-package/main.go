package main

import (
	"fmt"
	"mymodule/011-fmt-package/poker"
)

func main() {

	fmt.Println("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")

	for i := 1; i < 4; i++ {
		fmt.Printf("\nHand number: %v\n", i)
		poker.Deal()
	}

	// see emojis
	// Windows logo key + .
}
