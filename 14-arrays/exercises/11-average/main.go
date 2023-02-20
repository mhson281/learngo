// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

import ("fmt"; "os"; "strconv")

func main() {
    args := os.Args[1:]
    if len(args) > 5 || len(args) < 1{
        fmt.Println("Please tell me numbers (maximum 5 numbers).")
        return
    }

    var (
        sum float64
        nums [5]float64
        numCount float64
    )
    for index, value := range args{
        if n, err := strconv.ParseFloat(value, 64); err == nil {
            nums[index] = n
            numCount++
            sum += n
        } else {
            continue
        }
    }

    fmt.Printf("Your number: %v\n", nums)
    fmt.Printf("Average: %.0f\n", sum / numCount)
}
