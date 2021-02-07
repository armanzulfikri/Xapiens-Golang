package models

import (
	"gorm.io/gorm"
)

//Persons model
type Persons struct {
	SubDistrictID uint
	gorm.Model
	Nip                        string
	FullName                   string
	FirstName                  string
	LastName                   string
	BirthDate                  string
	BirthPlace                 string
	Gender                     string
	Photo                      string
	ZonaLocation               string
	PersonHadleOfficeLocations []PersonHadleOfficeLocations `gorm:"ForeignKey:PersonID"`
}
