package common

import (
	"github.com/skingford/colly-spider/model/pubmed"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&pubmed.PubmedArticle{},
		&pubmed.PubmedAuthor{},
		&pubmed.PubmedAbstract{},
	)
}
