// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------
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

