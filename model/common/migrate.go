package common

import (
	"colly-spider/model"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.FattyLiverArticle{},
		&model.FattyLiverAuthor{},
		&model.FattyLiverAbstract{},
	)
}
