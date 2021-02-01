package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var (
		checkComments, checkAlbums, checkPhotos bool
	)

	db.Migrator().DropTable(&Comments{})
	db.Migrator().DropTable(&Albums{})
	db.Migrator().DropTable(&Photos{})

	checkComments = db.Migrator().HasTable(&Comments{})
	if !checkComments {
		db.Migrator().CreateTable(&Comments{})
		fmt.Println("Comments table created")
	}
	checkAlbums = db.Migrator().HasTable(&Albums{})
	if !checkAlbums {
		db.Migrator().CreateTable(&Albums{})
		fmt.Println("Albums table created")
	}
	checkPhotos = db.Migrator().HasTable(&Photos{})
	if !checkPhotos {
		db.Migrator().CreateTable(&Photos{})
		fmt.Println("Photos table created")
	}

}
