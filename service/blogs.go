package service

import (
	"context"
	"github.com/hariprasanth3031/go-blog-service/db"
	"github.com/hariprasanth3031/go-blog-service/models"
)

func CreatePost(ctx context.Context, input models.Blog) (*models.Blog, error) {
	return db.CreatePost(ctx, &input)
}

func GetPost(ctx context.Context, id uint64) (*models.Blog, error) {
	return db.GetPost(ctx, id)
}

func UpdatePost(ctx context.Context, input models.Blog) error {

	updates := make(map[string]interface{}, 0)

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

	return db.UpdatePost(ctx, updates, input.Id)
}

func DeletePost(ctx context.Context, id uint64) error {
	return db.DeletePost(ctx, id)
}
