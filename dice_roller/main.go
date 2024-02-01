package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// Define flags
	adv := flag.Bool("adv", false, "Roll twice and take the higher")
	dis := flag.Bool("dis", false, "Roll twice and take the lower")

	// Parse command line arguments
	flag.Parse()

	// Regular expression to validate input format (XdY[+-]N)
	diceExpr := regexp.MustCompile(`^(\d+)d(\d+)([+-]\d+)?$`)

	fmt.Println("Enter dice expression ([number of dice]d[size of dice]+/-[modifier] ie:1d20+4): ")
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

	// Initialize total
	var total int

	// Check if adv or dis flag was set
	if *adv {
		// Roll twice and take the higher
		roll1 := rand.Intn(numSides) + 1
		roll2 := rand.Intn(numSides) + 1
		highestRoll := max(roll1, roll2)
		fmt.Printf("Die 1: %d\n", highestRoll)
		total += highestRoll
	} else if *dis {
		// Roll twice and take the lower
		roll1 := rand.Intn(numSides) + 1
		roll2 := rand.Intn(numSides) + 1
		lowestRoll := min(roll1, roll2)
		fmt.Printf("Die 1: %d\n", lowestRoll)
		total += lowestRoll
	} else {
		// Regular dice rolling
		for i := 0; i < numDice; i++ {
			roll := rand.Intn(numSides) + 1
			fmt.Printf("Die %d: %d\n", i+1, roll)
			total += roll
		}
	}

	// Apply modifier and print total
	total += modifier
	fmt.Printf("Total: %d (including modifier of %d)\n", total, modifier)
}

// Helper functions to find the maximum and minimum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
