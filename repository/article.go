package repository

import (
	"colly-spider/model"
	"gorm.io/gorm"
)

type ArticleRespository struct {
	DB *gorm.DB
}

func (r *ArticleRespository) Create(article *model.Article) error {
	return r.DB.Create(article).Error
}

func (r *ArticleRespository) GetList(pageIndex int, pageSize int) ([]model.Article, error) {
	if pageIndex == 0 {
		pageIndex =1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageIndex-1)*pageSize
	var articles []model.Article
	err := r.DB.Preload("Authors").Preload("Abstracts").Offset(offset).Limit(pageSize).Find(&articles).Error
	return articles, err
}

func (r *ArticleRespository) Total()  (*int64,error){
	var count int64
	err :=	r.DB.Model(&model.Article{}).Count(&count).Error
	return &count ,err
}