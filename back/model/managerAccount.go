package model

import "github.com/jinzhu/gorm"

type ManagerAccount struct {
	gorm.Model
	ManagerAccountID string
	Password         string
	Email            string
}
