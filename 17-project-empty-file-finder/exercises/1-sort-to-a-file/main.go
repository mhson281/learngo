// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

import (
    "os"
    "io/ioutil"
    "fmt"
    "sort"
)

func main() {
    fruits := os.Args[1:]
    if len(fruits) == 0 {
        fmt.Println("Send me some items and I will sort them")
        return
    }
    sort.Strings(fruits)
    
    var fruitsInByte []byte
    
    for _, fruit := range(fruits) {
        fruitsInByte = append(fruitsInByte, fruit...)
        fruitsInByte = append(fruitsInByte, '\n')
    }
    
    err := ioutil.WriteFile("sorted.txt", fruitsInByte, 0644)

    if err != nil{
        fmt.Println(err)
        return
    }
}
