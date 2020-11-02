package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/riteshharjani/grpc-hands-on/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calcpb.CalculatorRequest) (*calcpb.CalculatorResponse, error) {
	fmt.Println("server: Received calc RPC: ", req)
	num1 := req.Calc.Num1
	num2 := req.Calc.Num2
	//num1 := req.GetCalc().GetNum1()
	//num2 := req.GetCalc().GetNum2()
	res := calcpb.CalculatorResponse{
		Resp: int64(num1 + num2),
	}
	return &res, nil
}

func main() {
	fmt.Println("server: Hello World")

	// write a listner
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Server: Listner failed", err)
	}

	// now we need a server
	serv := grpc.NewServer()

	calcpb.RegisterCalcServiceServer(serv, &server{})

	if err := serv.Serve(lis); err != nil {
		log.Fatal("server: Accept failed", err)
	}
}
