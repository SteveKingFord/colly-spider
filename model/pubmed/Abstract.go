package pubmed

import "gorm.io/gorm"

type PubmedAbstract struct {
	gorm.Model
	PubmedArticleId uint   `json:"pubmed_article_id"`
	Content         string `json:"content"`
	Strong          string `json:"strong"`
}
