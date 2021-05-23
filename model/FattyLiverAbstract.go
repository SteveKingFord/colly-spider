package model

import "gorm.io/gorm"

type FattyLiverAbstract struct {
	gorm.Model
	FattyLiverArticleId uint   `json:"fatty_liver_article_id"`
	Content   string `json:"content"`
	Strong    string `json:"strong"`
}
