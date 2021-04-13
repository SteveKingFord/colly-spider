package repository

import (
	"colly-spider/model"
	"gorm.io/gorm"
)

type  ArticleRespository struct {
	DB *gorm.DB
}

func (r *ArticleRespository) Create(article *model.Article) error  {
	return r.DB.Create(article).Error
}