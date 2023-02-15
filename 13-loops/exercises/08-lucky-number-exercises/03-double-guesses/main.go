// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

You have two chance to guess the random number

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
    guess2, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess1 <= 0 || guess2 <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

    var bestGuess int
    if guess2 >= guess1 {
        bestGuess = guess2
    } else {
        bestGuess = guess1
    }

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(bestGuess) + 1

		if n == guess1 || n == guess2 {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉  YOU WIN!")
			case 1:
				fmt.Println("🎉  YOU'RE AWESOME!")
			case 2:
				fmt.Println("🎉  PERFECT!")
			}
			return
		}
	}

	msg := "%s Try again?\n"

	switch rand.Intn(2) {
	case 0:
		fmt.Printf(msg, "☠️  YOU LOST...")
	case 1:
		fmt.Printf(msg, "☠️  JUST A BAD LUCK...")
	}

}
