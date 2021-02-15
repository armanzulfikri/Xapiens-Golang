package models

import "gorm.io/gorm"

// OfficesLocation model
type OfficesLocation struct {
	gorm.Model
	Name                       string
	SubDistrictID              uint                         `json:"subdistrict_id"`
	PersonHadleOfficeLocations []PersonHadleOfficeLocations `gorm:"ForeignKey:OfficeLocationID"`
}
