package main

import (
	"fmt"
)

type room struct {
	id   int
	name string
	comm chan message
}

func createRoom(name string, id int, password string) room {
	messageChannel := make(chan message, 1000)

	r := room {
		id:   id,
		name: name,
		comm: messageChannel,
	}

	return r
}

func (room *room) process() {
	fmt.Println("processing room", room.name)

	for {
		m := <-room.comm
		fmt.Println("room", room.name, "received", m.content)
	}
}

func (room *room) pass(message message) {
	room.comm <- message
}