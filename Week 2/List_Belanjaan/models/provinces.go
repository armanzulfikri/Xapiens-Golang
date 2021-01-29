package models

import "gorm.io/gorm"

//Province sample

type Provinces struct {
	Name string
	gorm.Model
}
