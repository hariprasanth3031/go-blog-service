package main

import (
	"context"
	"encoding/json"
	"errors"
	"net"

	"github.com/hariprasanth3031/go-blog-service/config"
	"github.com/hariprasanth3031/go-blog-service/constants"
	"github.com/hariprasanth3031/go-blog-service/models"
	pb "github.com/hariprasanth3031/go-blog-service/protos"
	"github.com/hariprasanth3031/go-blog-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	port = ":8080"
)

type blogServer struct {
	pb.UnimplementedBlogsServer
}

func main() {

	config.InitializeLogger()
	config.InitializeEnv()
	config.InitializeDB()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		config.Logger.Error("failed to listen | Error -", err)
	}
	s := grpc.NewServer()

	pb.RegisterBlogsServer(s, &blogServer{})

	config.Logger.Debug("server listening at -", lis.Addr())
	if err := s.Serve(lis); err != nil {
		config.Logger.Error("failed to serve | Error -", err)
	}
}

func (s *blogServer) Create(ctx context.Context, input *pb.CreatePostRequest) (*pb.PostResponse, error) {

	var (
		inputTags  []byte
		outputTags []string
		err        error
		post       *models.Blog
	)

	inputTags, err = json.Marshal(input.Post.Tags)

	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrCreatePost, ": %v", err)
	}

	inputReq := models.Blog{
		Title:           input.Post.Title,
		Content:         input.Post.Content,
		Author:          input.Post.Author,
		PublicationDate: input.Post.PublicationDate,
		Tags:            string(inputTags),
	}

	post, err = service.CreatePost(ctx, inputReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrCreatePost, ": %v", err)
	}

	err = json.Unmarshal([]byte(post.Tags), &outputTags)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrCreatePost, ": %v", err)
	}

	return &pb.PostResponse{
		Post: &pb.PostDetails{
			Id:              post.Id,
			Title:           post.Title,
			Content:         post.Content,
			Author:          post.Author,
			PublicationDate: post.PublicationDate,
			Tags:            outputTags,
		},
		Status: constants.RecordCreated,
	}, nil
}

func (s *blogServer) Get(ctx context.Context, input *pb.PostId) (*pb.PostResponse, error) {

	var (
		tag  []string
		err  error
		post *models.Blog
	)

	post, err = service.GetPost(ctx, uint64(input.PostId))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, constants.ErrRecordNotFound)
		}
		return nil, status.Errorf(codes.Internal, constants.ErrFetchPost, ": %v", err)
	}

	err = json.Unmarshal([]byte(post.Tags), &tag)

	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrFetchPost, ": %v", err)
	}

	return &pb.PostResponse{
		Post: &pb.PostDetails{
			Id:              post.Id,
			Title:           post.Title,
			Content:         post.Content,
			Author:          post.Author,
			PublicationDate: post.PublicationDate,
			Tags:            tag,
		},
		Status: constants.RecordFetched,
	}, nil
}

func (s *blogServer) Update(ctx context.Context, input *pb.UpdatePost) (*pb.PostResponse, error) {

	var (
		tags []byte
		err  error
	)

	tags, err = json.Marshal(input.UpdateInput.Tags)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrUpdatePost, ": %v", err)
	}

	inputReq := models.Blog{
		Id:              input.UpdateInput.Id,
		Title:           input.UpdateInput.Title,
		Content:         input.UpdateInput.Content,
		Author:          input.UpdateInput.Author,
		PublicationDate: input.UpdateInput.PublicationDate,
		Tags:            string(tags),
	}

	err = service.UpdatePost(ctx, inputReq)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, constants.ErrRecordNotFound)
		}
		return nil, status.Errorf(codes.Internal, constants.ErrUpdatePost, ": %v", err)
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
		Status: constants.RecordUpdated,
	}, nil

}

func (s *blogServer) Delete(ctx context.Context, input *pb.PostId) (*pb.DeletePost, error) {

	var err error

	err = service.DeletePost(ctx, uint64(input.PostId))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, constants.ErrRecordNotFound)
		}
		return nil, status.Errorf(codes.Internal, constants.ErrDeletePost, ": %v", err)
	}

	return &pb.DeletePost{
		Status: constants.RecordDeleted,
	}, nil
}
