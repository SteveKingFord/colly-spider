package model

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	ArticleId uint   `json:"articleId"`
	Content   string `json:"content"`
	Strong    string `json:"strong"`
}
