package server

type ChatServer interface {
	Listen(address string) error         // listen for connections from clients
	Broadcast(command interface{}) error // broadcast messages from the client to all clients
	Start()                              // provide services to clients
	Close()                              // shut down the server
}
