package models

import "gorm.io/gorm"

// OfficesLocation model
type OfficesLocation struct {
	gorm.Model
	Name                       string
	SubDistrictID              uint
	PersonHadleOfficeLocations []PersonHadleOfficeLocations `gorm:"ForeignKey:OfficeLocationID"`
}
