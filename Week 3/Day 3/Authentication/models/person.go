package models

import (
	"gorm.io/gorm"
)

type Persons struct {
	SubDistrictsID uint
	gorm.Model
	Nip          string
	FullName     string
	LastName     string
	BirthDate    string
	BirthPlace   string
	Gender       string
	Photo        string
	ZonaLocation string
}
