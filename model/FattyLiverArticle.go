package model

import "gorm.io/gorm"

type FattyLiverArticle struct {
	gorm.Model
	Href      string    `json:"href"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	FattyLiverAuthors   []FattyLiverAuthor  `json:"fatty_liver_author"`
	FattyLiverAbstracts []FattyLiverAbstract `json:"fatty_liver_abstracts"`
}
