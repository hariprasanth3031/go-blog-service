package db

import (
	"context"
	"github.com/hariprasanth3031/go-blog-service/config"
	"github.com/hariprasanth3031/go-blog-service/models"
	"gorm.io/gorm"
)

type BlogDb interface {
	CreatePost(ctx context.Context, input *models.Blog) (*models.Blog, error)
	GetPost(ctx context.Context, id uint64) (*models.Blog, error)
	UpdatePost(ctx context.Context, updates map[string]interface{}, id uint64) error
	DeletePost(ctx context.Context, id uint64) error
}

type blogDb struct {
	db *gorm.DB
}

func NewBlogRepo() BlogDb {
	return &blogDb{
		db: config.MariaDB,
	}
}

func (d *blogDb) CreatePost(ctx context.Context, input *models.Blog) (*models.Blog, error) {

	res := d.db.WithContext(ctx).Debug().
		Table("blog").
		Create(&input)

	if res.Error != nil {
		return input, res.Error
	}

	return input, nil

}

func (d *blogDb) GetPost(ctx context.Context, id uint64) (*models.Blog, error) {

	var output models.Blog

	res := d.db.WithContext(ctx).Debug().
		Table("blog").
		Where("id = ?", id).
		Find(&output)

	if res.Error != nil {
		return &output, res.Error
	}

	if res.RowsAffected == 0 {
		return &output, gorm.ErrRecordNotFound
	}

	return &output, nil

}

func (d *blogDb) UpdatePost(ctx context.Context, updates map[string]interface{}, id uint64) error {

	res := d.db.WithContext(ctx).Debug().
		Table("blog").
		Where("id = ?", id).
		Updates(updates)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (d *blogDb) DeletePost(ctx context.Context, id uint64) error {

	res := d.db.WithContext(ctx).Debug().
		Table("blog").
		Delete(&models.Blog{Id: id})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}
