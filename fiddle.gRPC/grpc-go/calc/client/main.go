package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/calc/proto"
)

var addr string = "localhost:50051"

func main() {
	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer connection.Close()

	client := pb.NewCalcServiceClient(connection)

	//doSum(client)
	//doPrimes(client)
	//doAvg(client)
	//doMax(client)
	doSqrt(client, 9)
}