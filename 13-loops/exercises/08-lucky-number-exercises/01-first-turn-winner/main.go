// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

import ("time"; "math/rand"; "strconv"; "os"; "fmt")

const maxTurns = 5

func main() {
    rand.Seed(time.Now().UnixNano())

    if len(os.Args) != 2 {
        fmt.Println("You need to take a guess what the number is!!")
        return
    }

    args := os.Args[1:]
    guess, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Println("Not a number.")
        return
    } else if guess < 0 {
        fmt.Println("Please enter a number greater than 0")
        return
    }

    for turn := 1; turn <= maxTurns; turn++ {
        n := rand.Intn(guess + 1)

        if n != guess {
            continue
        }

        switch turn {
        case 1:
            fmt.Println("First turn winner!")
            fallthrough
        default:
            fmt.Println("You won!")
            return
        }
    } 

    fmt.Println("YOU LOST!!")
}
