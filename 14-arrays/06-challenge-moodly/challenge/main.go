// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------
import ("math/rand"; "os"; "fmt"; "time")

func getMood(name string) string {
    rand.Seed(time.Now().UnixNano())
    moodArray := [...]string{"good 👍", "happy 😀", "awesome 😎", "terrible 😩", "sad 😞", "bad 👎"}

    n := rand.Intn(len(moodArray))

    return name + " feels " + moodArray[n]
}

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        fmt.Println("Please enter a name")
        return
    }

    fmt.Println(getMood(args[0]))

}
