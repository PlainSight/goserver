package main

import (
	"fmt"
	"time"
)

type game struct {
	id   int
	name string
	comm chan message
}

const tickLength = 50

func createGame(name string, id int, password string) game {
	messageChannel := make(chan message, 1000)

	r := game {
		id:   id,
		name: name,
		comm: messageChannel,
	}

	return r
}

func (game *game) process() {
	fmt.Println("processing room", game.name)

	for {
		time.Sleep(tickLength * time.Millisecond)

		for {
			m := <-game.comm

			fmt.Println("game", game.name, "received", m.content)
		}
	}
}

func (game *game) pass(message message) {
	game.comm <- message
}