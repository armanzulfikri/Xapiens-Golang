package models

import "gorm.io/gorm"

type SubDistricts struct {
	Name string
	gorm.Model
	DistrictID uint
}
