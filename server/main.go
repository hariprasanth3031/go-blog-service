package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/hariprasanth3031/go-blog-service/config"
	"github.com/hariprasanth3031/go-blog-service/models"
	pb "github.com/hariprasanth3031/go-blog-service/protos"
	"github.com/hariprasanth3031/go-blog-service/service"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type blogServer struct {
	pb.UnimplementedBlogsServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterBlogsServer(s, &blogServer{})
	config.InitializeEnv()
	config.InitializeDB()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *blogServer) Create(ctx context.Context, input *pb.CreatePostRequest) (*pb.PostResponse, error) {

	var tags []byte

	tags, _ = json.Marshal(input.Post.Tags)

	inputReq := models.Blog{
		Title:           input.Post.Title,
		Content:         input.Post.Content,
		Author:          input.Post.Author,
		PublicationDate: input.Post.PublicationDate,
		Tags:            string(tags),
	}

	fmt.Println("input", inputReq.Title)

	post, err := service.CreatePost(ctx, inputReq)

	if err != nil {
		return &pb.PostResponse{
			Post:   nil,
			Status: "Error while creating the post",
		}, nil
	}

	var tag []string

	json.Unmarshal([]byte(post.Tags), &tag)

	return &pb.PostResponse{
		Post: &pb.PostDetails{
			Id:              post.Id,
			Title:           post.Title,
			Content:         post.Content,
			Author:          post.Author,
			PublicationDate: post.PublicationDate,
			Tags:            tag,
		},
		Status: "Created Successfully",
	}, nil
}

func (s *blogServer) Get(ctx context.Context, input *pb.PostId) (*pb.PostResponse, error) {

	blog, err := service.GetPost(ctx, uint64(input.PostId))

	if err != nil {
		return &pb.PostResponse{
			Post:   nil,
			Status: "Error while updating the post",
		}, nil
	}

	var tag []string

	json.Unmarshal([]byte(blog.Tags), &tag)

	return &pb.PostResponse{
		Post: &pb.PostDetails{
			Id:              blog.Id,
			Title:           blog.Title,
			Content:         blog.Content,
			Author:          blog.Author,
			PublicationDate: blog.PublicationDate,
			Tags:            tag,
		},
		Status: "Fetched Successfully",
	}, nil
}

func (s *blogServer) Update(ctx context.Context, input *pb.UpdatePost) (*pb.PostResponse, error) {

	var tags []byte

	tags, _ = json.Marshal(input.UpdateInput.Tags)

	inputReq := models.Blog{
		Title:           input.UpdateInput.Title,
		Content:         input.UpdateInput.Content,
		Author:          input.UpdateInput.Author,
		PublicationDate: input.UpdateInput.PublicationDate,
		Tags:            string(tags),
	}

	err := service.UpdatePost(ctx, inputReq)

	if err != nil {
		return &pb.PostResponse{
			Post:   nil,
			Status: "Error while updating the post",
		}, nil
	}

	return &pb.PostResponse{
		Post: &pb.PostDetails{
			Id:              input.UpdateInput.Id,
			Title:           input.UpdateInput.Title,
			Content:         input.UpdateInput.Content,
			Author:          input.UpdateInput.Author,
			PublicationDate: input.UpdateInput.PublicationDate,
			Tags:            input.UpdateInput.Tags,
		},
		Status: "Created Successfully",
	}, nil

}

func (s *blogServer) Delete(ctx context.Context, input *pb.PostId) (*pb.DeletePost, error) {

	err := service.DeletePost(ctx, uint64(input.PostId))

	if err != nil {
		return &pb.DeletePost{
			Status: err.Error(),
		}, nil
	}

	return &pb.DeletePost{
		Status: "Successfully deleted",
	}, nil
}
