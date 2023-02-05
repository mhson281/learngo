// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

import (
    "os"
    "strconv"
    "fmt"
) 


func main() {
    if len(os.Args) != 2 {
        fmt.Println("Pick a number")
        return
    }

    num, err := strconv.Atoi(os.Args[1])

    if err != nil || num <= 0{
        fmt.Printf("%q is not a number or it is less than 0", os.Args[1])
    } else if num % 2 == 0 && num % 8 == 0 {
        fmt.Printf("%d is an even number and it's divisible by 8", num)
    } else if num % 2 == 0 {
        fmt.Printf("%d is an even number", num)
    } else {
        fmt.Printf("%d is an odd number", num)
    } 
}
