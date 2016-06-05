package main

import (
	"net"
	"strconv"
)

type client struct {
	id int
	name string
	address *net.UDPAddr
}

var lastId int = 0
var clientMap = make(map[string]client)

func getClient(address *net.UDPAddr, name string) client {

	key := address.IP.String() + ":" + strconv.Itoa(address.Port)

	cl, exists := clientMap[key]

	if !exists {
		cl = client {
			id: lastId,
			name: name,
			address: address,
		}

		clientMap[key] = cl

		lastId++
	}

	return cl
}

func (client *client) writeToClient(connection *net.UDPConn, message string) {
	connection.WriteToUDP([]byte(client.name + ": " + message), client.address)
}