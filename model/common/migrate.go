package common

import (
	"colly-spider/model/FattyLiver"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&FattyLiver.FattyLiverArticle{},
		&FattyLiver.FattyLiverAuthor{},
		&FattyLiver.FattyLiverAbstract{},
	)
}
