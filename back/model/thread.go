package model

import "github.com/jinzhu/gorm"

type Thread struct {
	gorm.Model
	CommunityID  uint
	ThreadName   string
	threadNumber uint
}
