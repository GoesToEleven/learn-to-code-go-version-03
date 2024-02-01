package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// Regular expression to validate input format (XdY[+-]N)
	diceExpr := regexp.MustCompile(`^(\d+)d(\d+)([+-]\d+)?$`)

	fmt.Println("Enter dice expression (XdY[+-]N): ")
	var input string
	fmt.Scanln(&input)

	// Validate input format
	if !diceExpr.MatchString(input) {
		fmt.Println("Invalid input format. Please enter XdY[+-]N where X is the number of dice, Y is the number of sides, and N is the optional modifier.")
		return
	}

	// Extract number of dice, sides, and modifier (if present) from input
	matches := diceExpr.FindStringSubmatch(input)
	numDice, _ := strconv.Atoi(matches[1])
	numSides, _ := strconv.Atoi(matches[2])
	modifier := 0
	if len(matches) > 3 {
		modifier, _ = strconv.Atoi(matches[3])
	}

	// Seed random number generator
	rand.Seed(int64(time.Now().UnixNano()))

	// Roll dice and sum results
	var total int
	for i := 0; i < numDice; i++ {
		roll := rand.Intn(numSides) + 1
		fmt.Printf("Die %d: %d\n", i+1, roll)
		total += roll
	}

	// Apply modifier and print total
	total += modifier
	fmt.Printf("Total: %d (including modifier of %d)\n", total, modifier)
}
