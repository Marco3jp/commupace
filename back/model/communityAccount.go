package model

import "github.com/jinzhu/gorm"

type CommunityAccount struct {
	gorm.Model
	ManagerAccountID uint
	DisplayID        string `gorm:"unique;not null"`
	DisplayName      string
	Icon             string
	Status           string
}
