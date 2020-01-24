package main

import (
	"context"
	"fmt"
	"gRPC_exampleStream/server/counter"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	addr = "192.168.88.246:10000"
)

type server struct {
	counter.UnimplementedCounterServer
	clients map[*client]bool
	register chan *client
	unregister chan *client
	broadcast chan *counter.Count
}

type client struct {
	server *server
	send chan *counter.Count
}

func newServer() *server {
	s := &server{
		clients: make(map[*client]bool),
		register: make(chan *client),
		unregister: make(chan *client),
		broadcast: make(chan *counter.Count)}
	go s.run()
	return s
}

func (s *server) run() {
	s.startCounter()
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}
		case message := <-s.broadcast:
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	counter.RegisterCounterServer(s, newServer())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

func (s *server) GetCount(ctx context.Context, command *counter.Command) (*counter.Count, error) {
	var count counter.Count

	newClient := &client{server:s, send:make(chan *counter.Count)}
	newClient.server.register <- newClient
	fmt.Printf("Client connected: %t\n", s.clients[newClient])

	switch command.CommandName {
	case "start":
		s.startCounter()
		count := *<-newClient.send
		fmt.Printf("Count start: %d\n", count.Value)
		s.unregister <-newClient
		fmt.Printf("Client disconnected!\n")
	case "reset":
		s.resetCounter()
		count := *<-newClient.send
		fmt.Printf("Count reset: %d\n", count.Value)
		s.unregister <-newClient
		fmt.Printf("Client disconnected!\n")
	case "stop":
		s.stopCounter()
		count := *<-newClient.send
		fmt.Printf("Count stop: %d\n", count.Value)
		s.unregister <-newClient
		fmt.Printf("Client disconnected!\n")
	}
	return &count, nil
}

func (s *server) StreamFromCount(streamField *counter.StreamField, stream counter.Counter_StreamFromCountServer) error {
	newClient := &client{server:s, send:make(chan *counter.Count)}
	newClient.server.register <- newClient
	fmt.Printf("Client connected: %t\n", s.clients[newClient])

	switch streamField.FieldName{
	case "value":
		for {
			if err := stream.Send(<-newClient.send); err != nil {
				s.unregister <- newClient
				fmt.Printf("Client disconnected!\n")
				return err
			}
		}
	case "timestamp":
		for {
			if err := stream.Send(<-newClient.send); err != nil {
				s.unregister <- newClient
				fmt.Printf("Client disconnected!\n")
				return err
			}
		}
	}
	return nil
}

func (s *server) StatisticsAboutCount(stream counter.Counter_StatisticsAboutCountServer) error {

	return nil
}

func (s *server) CountHandler(stream counter.Counter_CountHandlerServer) error {
	for {
		in, _ := stream.Recv()
		/*if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
*/
		fmt.Printf("Message: %s\n", in)

		if err := stream.Send(in); err != nil {
			return err
		}
	}
	return nil
}