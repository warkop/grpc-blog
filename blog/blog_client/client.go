package main

import (
	"blog/blog/blogpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create blog
	blog := &blogpb.Blog{
		AuthorId: "john",
		Title:    "my book",
		Content:  "konten orang",
	}
	createBlog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("createblog: %v", createBlog)
}
