package main

import (
	"context"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Bruce Wayne",
		Title: "My first case",
		Content: "Who is the Joker !!",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}