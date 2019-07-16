package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	CommunityAccountID uint
	ThreadID           uint
	CommunityID        uint
	PostText           string
	PostNumber         uint
	PostType           string
	PostPath           string
}

type PostData struct {
	Post
	CommunityAccount
}
