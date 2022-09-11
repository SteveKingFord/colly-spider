package repository

import (
	"fmt"

	"github.com/skingford/colly-spider/model/pubmed"

	"gorm.io/gorm"
)

type PubmedArticleRepository struct {
	DB *gorm.DB
}

func (r *PubmedArticleRepository) Create(article *pubmed.PubmedArticle) error {
	result := r.DB.Where("href = ?", article.Href).First(article)

	fmt.Println("result.article:", result.RowsAffected)

	if result.RowsAffected != 0 {
		return result.Error
	}

	return r.DB.Create(article).Error
}

func (r *PubmedArticleRepository) GetList(pageIndex int, pageSize int) ([]pubmed.PubmedArticle, error) {
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageIndex - 1) * pageSize
	var articles []pubmed.PubmedArticle
	err := r.DB.Preload("PubmedAuthors").Preload("PubmedAbstracts").Offset(offset).Limit(pageSize).Find(&articles).Error
	return articles, err
}

func (r *PubmedArticleRepository) Total() (*int64, error) {
	var count int64
	err := r.DB.Model(&pubmed.PubmedArticle{}).Count(&count).Error
	return &count, err
}
