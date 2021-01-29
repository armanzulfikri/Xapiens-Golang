package models

import "gorm.io/gorm"

//Province sample

type Provinces struct {
	Name string
	gorm.Model
	Districts []Districts `gorm:"foreignKey:ProvinceID"`
}
