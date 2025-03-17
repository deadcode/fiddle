package main

import (
	"context"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/calc/proto"
)

func doAvg(c pb.CalcServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream")
	}

	numbers := []int32{3, 5, 9, 54, 39}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving avg: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}