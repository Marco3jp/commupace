package model

import "github.com/jinzhu/gorm"

type Community struct {
	gorm.Model
	CommunityId          string `gorm:"UNIQUE"`
	SpaceID              uint
	CommunityName        string
	CommunityDescription string
	CommunityIcon        string
	IsSpaceDirect        bool
}
