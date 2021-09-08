package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (*Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received mesage body from client: %v", message.Body)
	return &Message{Body: "Hello " + message.Body + "!"}, nil
}
