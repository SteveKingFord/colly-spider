package Fibrosis

import (
	"gorm.io/gorm"
)

type FibrosisArticle struct {
	gorm.Model
	Href      string               `json:"href"`
	Type      string               `json:"type"`
	Title     string               `json:"title"`
	EupDate   string               `json:"eup_date"`
	FibrosisAuthors   []FibrosisAuthor   `json:"fibrosis_authors"`
	FibrosisAbstracts []FibrosisAbstract `json:"fibrosis_abstracts"`
}
