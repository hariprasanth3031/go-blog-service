package service

import (
	"context"
	"errors"

	"github.com/hariprasanth3031/go-blog-service/db"
	"github.com/hariprasanth3031/go-blog-service/models"
)

type BlogService interface {
	CreatePost(ctx context.Context, input models.Blog) (*models.Blog, error)
	GetPost(ctx context.Context, id uint64) (*models.Blog, error)
	UpdatePost(ctx context.Context, input models.Blog) error
	DeletePost(ctx context.Context, id uint64) error
}

type blogService struct {
	blogDb db.BlogDb
}

func NewBlogService(blogDb db.BlogDb) BlogService {
	return &blogService{
		blogDb: blogDb,
	}
}

func (s *blogService) CreatePost(ctx context.Context, input models.Blog) (*models.Blog, error) {
	return s.blogDb.CreatePost(ctx, &input)
}

func (s *blogService) GetPost(ctx context.Context, id uint64) (*models.Blog, error) {
	return s.blogDb.GetPost(ctx, id)
}

func (s *blogService) UpdatePost(ctx context.Context, input models.Blog) error {

	updates := make(map[string]interface{}, 0)

	if input.Id < 0 {
		return errors.New("id invalid")
	}

	if input.Author != "" {
		updates["author"] = input.Author
	}

	if input.Title != "" {
		updates["title"] = input.Title
	}

	if input.Content != "" {
		updates["content"] = input.Content
	}

	if input.Tags != "" {
		updates["tags"] = input.Tags
	}

	return s.blogDb.UpdatePost(ctx, updates, input.Id)
}

func (s *blogService) DeletePost(ctx context.Context, id uint64) error {

	//var err error
	err := s.blogDb.DeletePost(ctx, id)

	return err
}
