// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------
import (
    "fmt"
    "os"
    "bufio"
)


type item struct {
    id int
    name string
    price int
}

type game struct {
    item 
    genre string
}

func main() {
    games := []game{
        {item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
        {item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
        {item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
    }

    fmt.Printf("Minh's game store has %d games", len(games))

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
  > list   : lists all the games
  > quit   : quits
    `)

    fmt.Println()

    if !in.Scan() {
        break
    }

    switch in.Text() {
        case "quit":
            fmt.Println("See you next time!")
            return
        case "list":
            for _, game := range games {
                fmt.Printf("%d\t%-15s\t%-20d\t%s\n", game.item.id, game.name, game.price, game.genre)
            }

        }
    }
}
