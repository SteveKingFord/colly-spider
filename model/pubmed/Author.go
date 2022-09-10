package pubmed

import "gorm.io/gorm"

type PubmedAuthor struct {
	gorm.Model
	Name            string `json:"name"`
	PubmedArticleId uint   `json:"pubmed_article_id"`
}
