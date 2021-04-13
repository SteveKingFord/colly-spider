package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name      string `json:"name"`
	ArticleId uint   `json:"articleId"`
}
