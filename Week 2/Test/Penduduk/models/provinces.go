package models

import "gorm.io/gorm"

// Provinces model
type Provinces struct {
	Name      string      `json : "name"`
	Districts []Districts `gorm:"foreignKey:ProvinceID" json :"districts" `
	gorm.Model
}
