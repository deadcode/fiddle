package main

import (
	"io"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/calc/proto"
	"google.golang.org/grpc"
)

func (s * Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	log.Println("Max function was invoked")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error in receviging from client: %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number

			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}