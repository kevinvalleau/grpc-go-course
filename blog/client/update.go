package main

import (
	"context"
	"log"

	pb "github.com/kevinvalleau/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Kevin",
		Title:    "A new title",
		Content:  "Content different from the first blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error happened while updating %v\n", err)
	}

	log.Println("Blog was updated")

}
