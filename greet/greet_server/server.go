package main

import (
	"fmt"
	"log"
	"net"

	"github.com/riteshharjani/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

// this will be our server object
type server struct{}

func main() {
	fmt.Println("hello world")

	// write a listner
	// 50051 is the default port for gRPC
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// well we need to register a service
	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
