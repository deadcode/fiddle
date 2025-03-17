package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("Client stream ended")
			stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})

			break
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Recieving: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}

	return nil
}