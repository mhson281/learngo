// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments is not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

import ("fmt"; "os"; "strconv")

func main() {
    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Println("Please give me up to 5 numbers")
        return
    } else if len(args) > 5 {
        fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting 5 numbers...")
        return
    }

    var nums [5]int
    for i, v := range args {
        if n, err := strconv.Atoi(v); err == nil {
            nums[i] = n
        } else {
            nums[i] = 0
        }
    }

    fmt.Println(bubbleSort(nums))
}

func bubbleSort(array [5]int) [5]int {
    len := len(array)
    for i := 0; i < len-1; i++ {
        for j := 0; j < len-1; j++ {
            if array[j] > array[j+1] {
                array[j], array[j+1] = array[j+1], array[j]
            }
        }
    }
    return array
}
