package model

import "github.com/jinzhu/gorm"

type CommunityUser struct {
	gorm.Model
	CommunityAccountID uint
	CommunityID        uint
}
