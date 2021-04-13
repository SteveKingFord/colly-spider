package common

import (
	"colly-spider/model"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Article{},
		&model.Author{},
		&model.Content{},
	)
}
