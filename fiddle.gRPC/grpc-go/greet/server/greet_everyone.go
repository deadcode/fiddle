package main

import (
	"io"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Recvied greeting from: %s\n", req.FirstName)
		
		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending response to client: %v\n", err)
		}
	}
}