package main

import (
	"context"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---UpdateBlog was invoked---")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Alfred",
		Title: "Yes Master Bruce",
		Content: "He is a slippery one!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
	