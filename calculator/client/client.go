package main

import (
	"context"
	"fmt"
	"log"

	"github.com/riteshharjani/grpc-hands-on/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client: Hello there!")

	// we need to dial a connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client: failed to Dial", err)
	}

	//fmt.Println("client: connection", cc)
	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)

	doUnary(c)
}

func doUnary(c calcpb.CalcServiceClient) {

	fmt.Println("client: Client calling")

	resp, err := c.Calculate(context.Background(), &calcpb.CalculatorRequest{
		Calc: &calcpb.Calculator{
			Num1: 123,
			Num2: 456,
		},
	})

	if err != nil {
		log.Fatalln("client: Calculate req failed", err)
	}
	log.Println("client: resp received", resp)
}
