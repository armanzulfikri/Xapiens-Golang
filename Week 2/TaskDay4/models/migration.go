package models

import (
	"fmt"

	"gorm.io/gorm"
)

//Migration
func Migrations(db *gorm.DB) {
	var checkProvince, checkDistricts, checkRelation bool

	checkProvince = db.Migrator().HasTable(&Provinces{})
	if !checkProvince {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("create table provinces")
	}

	checkDistricts = db.Migrator().HasTable(&Districts{})
	if !checkDistricts {
		db.Migrator().CreateTable(&Districts{})
		fmt.Println("create table districts")
	}

	checkRelation = db.Migrator().HasConstraint(&Provinces{}, "District")
	if !checkRelation {
		db.Migrator().CreateConstraint(&Provinces{}, "District")
	}
}
