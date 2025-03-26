package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer connection.Close()

	client := pb.NewBlogServiceClient(connection)

	id := createBlog(client)

	ReadBlog(client, id)
	ReadBlog(client, "InvlidId")
	updateBlog(client, id)
	listBlog(client) // list once before delete
	deleteBlog(client, id)
	listBlog(client) // list again after delete
}