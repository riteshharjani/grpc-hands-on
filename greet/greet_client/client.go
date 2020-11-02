package main

import (
	"fmt"
	"log"

	"github.com/riteshharjani/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("hello from client")

	// 1st thing to do is to setup a connection to the server.
	// this is done using grpc.Dial()
	// by default grpc uses SSL, but for now we will use
	// insecure conn type
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("client: could not connect: %v", err)
	}

	defer cc.Close()

	// now we can create a client which will usee connection created above.
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f\n", c)
}
