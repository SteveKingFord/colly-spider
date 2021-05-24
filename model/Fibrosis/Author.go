package Fibrosis

import "gorm.io/gorm"

type FibrosisAuthor struct {
	gorm.Model
	Name      string `json:"name"`
	FibrosisArticleId uint   `json:"fibrosis_article_id"`
}
