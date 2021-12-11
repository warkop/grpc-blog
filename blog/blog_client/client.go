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
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	blogId := createBlogRes.GetBlog().GetId()

	// read blog
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "61b2119e12ae320038b78d12",
	})
	if err2 != nil {
		fmt.Printf("error happened while reading: %v", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("error happened while reading: %v", readBlogErr)
	}

	fmt.Println("blog was read: ", readBlogRes)

	// update blog
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "john",
		Title:    "my book(edited)",
		Content:  "konten orang-orang",
	}

	updateRes, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if err != nil {
		fmt.Printf("error happened while updating: %v", err)
	}

	fmt.Println("blog was updated: ", updateRes)

	// delete blog
	deleteRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogId})
	if err != nil {
		fmt.Printf("error happened while deleting: %v", err)
	}

	fmt.Println("blog was deleted: ", deleteRes)
}
