package models

type Blog struct {
	Id              uint64 `gorm:"column:id;type:bigint(20) unsigned;not null;primaryKey"`
	Title           string `gorm:"column:title;type:varchar(200);not null;"`
	Content         string `gorm:"column:content;type:varchar(200);not null;"`
	Author          string `gorm:"column:author;type:varchar(200);not null;"`
	PublicationDate uint64 `gorm:"column:publication_date;type:bigint(20);not null;"`
	Tags            string `gorm:"column:tags;type:varchar(200);not null;"`
}
