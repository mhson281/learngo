// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------
package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
)

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	fmt.Printf("Minh's game store has %d games.\n", len(games))

    gameIds := make(map[int]game)
    
    for _, g := range games {
       gameIds[g.id] = g
    }

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id #   : Get game by ID number
  > quit   : quits

`)

		if !in.Scan() {
			break
		}

		fmt.Println()

        cmd := strings.Fields(in.Text())
        if len(cmd) < 1 {
            fmt.Println("Please enter a command")
            return
        }

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
        case "id":
            if len(cmd) != 2 {
                fmt.Println("Please enter id and a number")
                return
            }
            
            id, err := strconv.Atoi(cmd[1])
            if err != nil {
                fmt.Println("Please enter a valid number")
                continue
            }

            g, ok := gameIds[id]

            if !ok {
                fmt.Println("Id does not exist")
                continue
            }

			fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "(" + g.genre + ")", g.price)
		}

	}
}
