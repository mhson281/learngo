// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

import (
    "fmt"
    "unicode/utf8"
)
func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

    for _,s := range words {
        fmt.Printf("%s has a length of %d and a rune count of %d\n", s, len(s), utf8.RuneCountInString(s))
		fmt.Printf("\tbytes   : %x \n", s)

        fmt.Print("\trunes   : ")
        for _, r := range s {
            fmt.Printf("%x ", r)
        }
        fmt.Println()
    }


	// Print the runes of the strings in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the strings as rune literals
	// Hint: Use for range

	// Print the first rune and its byte size of the strings
	// Hint: Use utf8.DecodeRuneInString

	// Print the last rune of the strings
	// Hint: Use utf8.DecodeLastRuneInString

	// Slice and print the first two runes of the strings

	// Slice and print the last two runes of the strings

	// Convert the string to []rune
	// Print the first and last two runes
}
