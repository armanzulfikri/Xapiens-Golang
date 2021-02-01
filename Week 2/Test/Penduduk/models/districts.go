package models

import "gorm.io/gorm"

// Districts model
type Districts struct {
	Name string
	gorm.Model
	ProvinceID   uint
	SubDistricts []SubDistricts `gorm:"foreignKey:DistrictID`
}
