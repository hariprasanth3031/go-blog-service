package db

import (
	"context"
	"github.com/hariprasanth3031/go-blog-service/config"
	"github.com/hariprasanth3031/go-blog-service/models"
	"gorm.io/gorm"
)

func CreatePost(ctx context.Context, input *models.Blog) (*models.Blog, error) {

	res := config.MariaDB.WithContext(ctx).Debug().
		Table("blog").
		Create(&input)

	if res.Error != nil {
		return input, res.Error
	}

	return input, nil

}

func GetPost(ctx context.Context, id uint64) (*models.Blog, error) {

	var output models.Blog

	res := config.MariaDB.WithContext(ctx).Debug().
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

func UpdatePost(ctx context.Context, updates map[string]interface{}, id uint64) error {

	res := config.MariaDB.WithContext(ctx).Debug().
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

func DeletePost(ctx context.Context, id uint64) error {

	res := config.MariaDB.WithContext(ctx).Debug().
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
