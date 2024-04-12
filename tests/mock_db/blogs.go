package mock_db

import (
	"context"

	"github.com/hariprasanth3031/go-blog-service/models"
	"github.com/stretchr/testify/mock"
)

type MockBlogsDb struct {
	mock.Mock
}

func (m *MockBlogsDb) CreatePost(ctx context.Context, input *models.Blog) (*models.Blog, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	if blog, ok := args.Get(0).(*models.Blog); ok {
		return blog, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBlogsDb) GetPost(ctx context.Context, id uint64) (*models.Blog, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	if blog, ok := args.Get(0).(*models.Blog); ok {
		return blog, args.Error(1)
	}
	return nil, args.Error(1)

}

func (m *MockBlogsDb) UpdatePost(ctx context.Context, updates map[string]interface{}, id uint64) error {
	args := m.Called(ctx, updates, id)
	return args.Error(0)
}

func (m *MockBlogsDb) DeletePost(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
