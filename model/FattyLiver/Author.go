package FattyLiver

import "gorm.io/gorm"

type FattyLiverAuthor struct {
	gorm.Model
	Name      string `json:"name"`
	FattyLiverArticleId uint   `json:"fatty_liver_article_id"`
}
