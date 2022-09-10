package pubmed

import "gorm.io/gorm"

type PubmedAbstract struct {
	gorm.Model
	PubmedArticleId uint   `json:"pubmed_article_id"`
	Content         string `json:"content" gorm:"type:text"`
	Strong          string `json:"strong"`
}
