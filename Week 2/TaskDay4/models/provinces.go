package models

import "gorm.io/gorm"

// Provinces model
type Provinces struct {
	Name string
	gorm.Model
	District []Districts `gorm:"foreignKey:ProvinceID"`
}
