package main

import (
	"context"
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
	//fmt.Printf("Created client: %f\n", c)

	// do unary here in below
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Printf("client: Starting to a unary gRPC\n")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ritesh",
			LastName:  "Harjani",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("client: error while calling Greet RPC: %v\n", err)
	}

	log.Printf("client: Response from Greet: %v\n", res.Result)
}
