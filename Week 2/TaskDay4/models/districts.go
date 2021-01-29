package models

import "gorm.io/gorm"

//Province sample

type Districts struct {
	gorm.Model
	Name        string
	ProvincesID uint
}
