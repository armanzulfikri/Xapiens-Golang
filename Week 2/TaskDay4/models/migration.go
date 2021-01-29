package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var checkTableProvinces, checkTableDistricts, checkConstraint bool

	checkTableProvinces = db.Migrator().HasTable(&Provinces{})
	if !checkTableProvinces {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("Create Table Provinces")
	}

	checkTableDistricts = db.Migrator().HasTable(&Districts{})
	if !checkTableDistricts {
		db.Migrator().CreateTable(&Districts{})
		fmt.Println("Create Table Districts")
	}

	checkConstraint = db.Migrator().HasConstraint(&Provinces{}, "District")
	if !checkConstraint {
		db.Migrator().CreateConstraint(&Provinces{}, "District")
	}
}
