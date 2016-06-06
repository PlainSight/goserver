package main

import (
	"net"
	"strconv"
)

type client struct {
	id int
	name string
	address *net.UDPAddr
	connection *net.UDPConn
}

var lastId int = 0
var clientMap = make(map[string]client)

func getClient(address *net.UDPAddr, name string, connection *net.UDPConn) client {

	key := address.IP.String() + ":" + strconv.Itoa(address.Port)

	cl, exists := clientMap[key]

	if !exists {
		cl = client {
			id: lastId,
			name: name,
			address: address,
			connection: connection,
		}

		clientMap[key] = cl

		lastId++
	}

	return cl
}

func (client *client) write(message string) {
	client.connection.WriteToUDP([]byte(client.name + ": " + message), client.address)
}