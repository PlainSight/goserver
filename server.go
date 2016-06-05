package main

import (
	//"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(0)
	}
}

func createRoom(name string, id int, password string) room {
	messageChannel := make(chan message, 1000)

	r := room{
		id:   id,
		name: name,
		comm: messageChannel,
	}

	return r
}

func main() {
	fmt.Printf("starting server\n")

	ServerAddr, err := net.ResolveUDPAddr("udp", ":3000")

	checkError(err)

	ServerConnection, err := net.ListenUDP("udp", ServerAddr)

	checkError(err)
	defer ServerConnection.Close()

	buffer := make([]byte, 1024)

	//jsonDecoder := json.NewDecoder()

	var rooms [20]room

	dr := createRoom("default", 0, "")
	roomCount := 0
	rooms[roomCount] = dr
	roomCount++

	go dr.process()

	i := 0

	for {
		messageLength, address, error := ServerConnection.ReadFromUDP(buffer)

		m := string(buffer[0:messageLength])

		fmt.Printf("Received Data: %s %s %v\n", m, address.IP.String(), address.Port)

		ServerConnection.WriteToUDP(buffer[0:messageLength], address)

		if error != nil {
			fmt.Printf("Error: %s\n", error)
		}

		smessage := message{
			id:      i,
			content: m,
		}

		for r := 0; r < roomCount; r++ {
			rooms[r].comm <- smessage
		}

		i++
	}

}

func readNetwork() {

}

type message struct {
	id      int
	content string
}

type room struct {
	id   int
	name string
	comm chan message
}

func (room *room) process() {
	fmt.Println("processing room", room.name)

	for {
		time.Sleep(2 * time.Second)

		m := <-room.comm

		fmt.Println("room", room.name, "received", m.content)

	}
}
