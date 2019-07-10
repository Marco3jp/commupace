package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	CommunityAccountID uint
	ThreadID           uint
	PostText           string
	PostNumber         uint
	PostType           string
	PostPath           string
}
