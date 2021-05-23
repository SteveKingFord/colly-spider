package repository

import (
	"colly-spider/model"
	"gorm.io/gorm"
)

type FattyLiverArticleRepository struct {
	DB *gorm.DB
}

func (r *FattyLiverArticleRepository) Create(article *model.FattyLiverArticle) error {
	return r.DB.Create(article).Error
}

func (r *FattyLiverArticleRepository) GetList(pageIndex int, pageSize int) ([]model.FattyLiverArticle, error) {
	if pageIndex == 0 {
		pageIndex =1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageIndex-1)*pageSize
	var articles []model.FattyLiverArticle
	err := r.DB.Preload("Authors").Preload("Abstracts").Offset(offset).Limit(pageSize).Find(&articles).Error
	return articles, err
}

func (r *FattyLiverArticleRepository) Total()  (*int64,error){
	var count int64
	err :=	r.DB.Model(&model.FattyLiverArticle{}).Count(&count).Error
	return &count ,err
}