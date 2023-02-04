package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	ID             int
	AuthorID       string
	AuthorUsername string
	Text           string
}
