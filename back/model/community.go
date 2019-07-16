package model

import "github.com/jinzhu/gorm"

type Community struct {
	gorm.Model
	CommunityId          string `gorm:"type:varchar(40);unique;not null"`
	SpaceID              uint
	CommunityName        string
	CommunityDescription string
	CommunityIcon        string
	IsSpaceDirect        bool
}
