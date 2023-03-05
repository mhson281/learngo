// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

import (
    "fmt"
    "strings"
    "strconv"
)
func main() {
	data := "2 4 6 1 3 5"
    dataSlice := strings.Fields(data)

    var nums []int 
    
    for _, v := range dataSlice{
        n,_ := strconv.Atoi(v)
        nums = append(nums, n)
    }

    var evens, odds []int
    evens = nums[0:3]
    odds = nums[3:]
    fmt.Printf("nums            : %v\n", dataSlice)
    fmt.Printf("evens           : %v\n", evens)
    fmt.Printf("odds            : %v\n", odds)
    fmt.Printf("middles         : %v\n", nums[2:4])
    fmt.Printf("first 2         : %v\n", nums[:2])
    fmt.Printf("last 2          : %v\n", nums[4:])

    fmt.Printf("evens last 1    : %v\n", evens[len(evens)-1:])
    fmt.Printf("odds last 2     : %v\n", odds[len(odds)-2:])
}
