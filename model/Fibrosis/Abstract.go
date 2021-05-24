package Fibrosis

import "gorm.io/gorm"

type FibrosisAbstract struct {
	gorm.Model
	FibrosisArticleId uint   `json:"fibrosis_article_id"`
	Content   string `json:"content"`
	Strong    string `json:"strong"`
}
