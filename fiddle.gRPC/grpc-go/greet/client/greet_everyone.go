package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveyone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest {
		{FirstName: "Bruce"},
		{FirstName: "Orangie"},
		{FirstName: "Yajat"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %s\n", req.FirstName)

			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving from server: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}