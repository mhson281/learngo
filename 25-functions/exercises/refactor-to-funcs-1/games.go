package main

import "fmt"

func newGame(id, price int, name string, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

func addGame(games []game, g game) []game {
	return append(games, g)
}

func load() (games []game) {
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	games = addGame(games, newGame(3, 20, "minecraft", "sandbox"))

	return
}

func indexByID(games []game) (byId map[int]game) {
	byId = make(map[int]game)
	for _, g := range games {
		byId[g.id] = g
	}

	return
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}
