package main

import (
	"server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":2020")

	// start the server
	s.Start()
}
