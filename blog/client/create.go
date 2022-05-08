package main

import (
	"context"
	"log"

	pb "github.com/kevinvalleau/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("CreateBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Kevin",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
