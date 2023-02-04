package models

import "gorm.io/gorm"

type Following struct {
	gorm.Model
	UserID            string
	Username          string
	FollowingID       string
	FollowingUsername string
}
