// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

import (
    "os"
    "fmt"
    "sort"
)
func main() {

    houses := map[string][]string{
        "gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
        "hufflepuf": {"wenlock", "scamander", "helga", "diggory"},
        "ravenclaw": {"bagnold", "wildsmith", "montmorency"},
        "slytherin": {"horace", "nigellus", "higgs", "scorpius"},
        "bobo": {"wizardry", "unwanted"},
    }

    // delete bobo from list of houses
    delete(houses, "bobo")

    args := os.Args[1:] 

    if len(args) != 1 {
        fmt.Println("Please type a Hogwarts house name.")
        return
    }

	house, students := args[0], houses[args[0]]
	if students == nil {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

    clone := append([]string(nil), students...)
    sort.Strings(clone)

    fmt.Printf("~~~ %s students ~~~\n\n", args[0])
    for _, wizard := range(clone) {
        fmt.Printf("\t+ %s\n", wizard)
    }
}
