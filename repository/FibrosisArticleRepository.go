package repository

import (
	"colly-spider/model/Fibrosis"
	"gorm.io/gorm"
)

type FibrosisArticleRepository struct {
	DB *gorm.DB
}

func (r *FibrosisArticleRepository) Create(article *Fibrosis.FibrosisArticle) error {
	return r.DB.Create(article).Error
}

func (r *FibrosisArticleRepository) GetList(pageIndex int, pageSize int) ([]Fibrosis.FibrosisArticle, error) {
	if pageIndex == 0 {
		pageIndex =1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageIndex-1)*pageSize
	var articles []Fibrosis.FibrosisArticle
	err := r.DB.Preload("FibrosisAuthors").Preload("FibrosisAbstracts").Offset(offset).Limit(pageSize).Find(&articles).Error
	return articles, err
}

func (r *FibrosisArticleRepository) Total()  (*int64,error){
	var count int64
	err :=	r.DB.Model(&Fibrosis.FibrosisArticle{}).Count(&count).Error
	return &count ,err
}