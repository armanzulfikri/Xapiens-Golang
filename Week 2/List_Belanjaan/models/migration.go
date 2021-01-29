package models

import (
	"fmt"

	"gorm.io/gorm"
)

//Migration
func Migrations(db *gorm.DB) {
	var check bool

	check = db.Migrator().HasTable(&Provinces{})
	if !check {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("create table provinces")
	}
}
