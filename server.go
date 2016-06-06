package main

import (
	//"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(0)
	}
}

func main() {
	fmt.Printf("starting server\n")

	serverAddr, err := net.ResolveUDPAddr("udp", ":3000")

	checkError(err)

	serverConnection, err := net.ListenUDP("udp", serverAddr)

	checkError(err)
	defer serverConnection.Close()

	buffer := make([]byte, 1024)

	//jsonDecoder := json.NewDecoder()

	var relays [20]relay

	defaultRoom := createRoom("default", 0, "")
	relayCount := 0
	relays[relayCount] = &defaultRoom
	relayCount++

	go defaultRoom.process()

	for i := 0 ;; i++ {
		messageLength, address, error := serverConnection.ReadFromUDP(buffer)

		if error != nil {
			fmt.Printf("Error: %s\n", error)
		}

		b := buffer[0:messageLength]
		m := string(b)

		fmt.Printf("Received Data: %s %s %v\n", m, address.IP.String(), address.Port)

		client := getClient(address, strings.TrimSpace(m))

		client.write(serverConnection, m)

		smessage := message{
			id:      i,
			content: m,
		}

		for r := 0; r < relayCount; r++ {
			relays[r].pass(smessage)
		}
	}

}

type message struct {
	id      int
	content string
}