package pubmed

import (
	"gorm.io/gorm"
)

type PubmedArticle struct {
	gorm.Model
	Href            string           `json:"href"`
	Type            string           `json:"type"`
	Title           string           `json:"title"`
	EupDate         string           `json:"eup_date"`
	PubmedAuthors   []PubmedAuthor   `json:"pubmed_author"`
	PubmedAbstracts []PubmedAbstract `json:"pubmed_abstracts"`
}
