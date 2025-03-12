package main

import (
	"context"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/calc/proto"
)

func doSum(c pb.CalcServiceClient) {
	log.Println("doSum is invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 3,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}