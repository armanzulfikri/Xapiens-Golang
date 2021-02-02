package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var checkTableProvinces, checkTableDistricts, checkTableSubDistricts, checkTablePerson bool

	db.Migrator().DropTable(&Provinces{})
	db.Migrator().DropTable(&SubDistricts{})
	db.Migrator().DropTable(&Districts{})
	db.Migrator().DropTable(&Persons{})

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

	checkTablePerson = db.Migrator().HasTable(&Persons{})
	if !checkTablePerson {
		db.Migrator().CreateTable(&Persons{})
		fmt.Println("Create Table Person")
	}

	checkTableSubDistricts = db.Migrator().HasTable(&SubDistricts{})
	if !checkTableSubDistricts {
		db.Migrator().CreateTable(&SubDistricts{})
		fmt.Println("Create Table SubsDistricts")
	}
}
