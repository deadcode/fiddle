package main

import (
	"log"
	"net"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/calc/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalcServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterCalcServiceServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}