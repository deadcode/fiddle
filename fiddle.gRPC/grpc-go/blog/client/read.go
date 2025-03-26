package main

import (
	"context"
	"log"

	pb "github.com/deadcode/fiddle/fiddle.gRPC/grpc-go/blog/proto"
)

func ReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}