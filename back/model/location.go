package model

import "github.com/jinzhu/gorm"

type Location struct {
	gorm.Model
	Coordinates
}

type Coordinates struct {
	Latitude  float64 `gorm:"index:lat"`
	Longitude float64 `gorm:"index:long"`
}
