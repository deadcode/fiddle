package main

import (
	"context"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet is invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Master Bruce",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}