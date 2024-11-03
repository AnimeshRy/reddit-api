package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch chan Message
}

func (s *Server) StartAndListen() {
	for {
		select {
		// block here until someone is sending a messag to the channel
		case msg := <-s.msgch:
			fmt.Printf("Received message from %s: %s\n", msg.From, msg.Payload)
		default:
		}
	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "server",
		Payload: payload,
	}
	msgch <- msg
}

func main() {
	s := &Server{
		msgch: make(chan Message),
	}
	go s.StartAndListen()
	var count = 4
	for i := 0; i < count; i++ {
		go func(i int) {
			time.Sleep(time.Duration(i) * 2 * time.Second)
			sendMessageToServer(s.msgch, fmt.Sprintf("Hello, world %d!", i+1))
		}(i)
	}

	select {}
}
