package main

import (
	"fmt"
)

type room struct {
	id   int
	name string
	password string
	comm chan message
	clients map[int]client
}

func createRoom(name string, id int, password string) room {
	messageChannel := make(chan message, 1000)

	r := room {
		id:   id,
		name: name,
		password: password,
		comm: messageChannel,
		clients: make(make[int]client)
	}

	return r
}

func (room *room) process() {
	for {
		m := <-room.comm

		for _, client := range room.clients {
			client.write
		}

		fmt.Println("room", room.name, "received", m.content)
	}
}

func (room *room) pass(message message) {
	room.comm <- message
}

func (room *room) register(client client) {
	
}