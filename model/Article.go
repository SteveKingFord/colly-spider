package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Href      string    `json:"href"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Authors   []Author  `json:"authors"`
	Abstracts []Content `json:"abstracts"`
}
