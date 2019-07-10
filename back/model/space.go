package model

import "github.com/jinzhu/gorm"

type Space struct {
	gorm.Model
	SpaceName  string
	LocationID uint
}
