// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// THIS EXERCISE IS OPTIONAL

// ---------------------------------------------------------
// EXERCISE: Print hexes
//
//  1. Print 0 to 9 in hexadecimal
//
//  2. Print 10 to 15 in hexadecimal
//
//  3. Print 17 in hexadecimal
//
//  4. Print 25 in hexadecimal
//
//  5. Print 50 in hexadecimal
//
//  6. Print 100 in hexadecimal
//
// EXPECTED OUTPUT
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15
//  17
//  25
//  50
//  100
//
// NOTES
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system
//
// ---------------------------------------------------------

func main() {
	fmt.Println(0x0, 0x1)
	fmt.Println(0xa, 0xf)
	fmt.Println(0x11)
	fmt.Println(0x19)
	fmt.Println(0x64)
}
