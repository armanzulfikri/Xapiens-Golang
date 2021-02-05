package models

import "gorm.io/gorm"

//user struct
type Users struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
