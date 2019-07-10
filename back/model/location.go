package model

import "github.com/jinzhu/gorm"

type Location struct {
	gorm.Model
	Latitude  float64 `gorm:"index:lat"`
	Longitude float64 `gorm:"index:long"`
}
