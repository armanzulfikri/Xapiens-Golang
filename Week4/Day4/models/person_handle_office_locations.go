package models

import "gorm.io/gorm"

//PersonHadleOfficeLocations models
type PersonHadleOfficeLocations struct {
	gorm.Model
	OfficeLocationID uint
	PersonID         uint
}
