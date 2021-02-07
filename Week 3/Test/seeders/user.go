package seeders

import (
	"CasePoint3/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUser func
func SeedUser(db *gorm.DB) {
	var userArray = [...][3]string{
		{"arman@xapiens.id", "guest", "guest"},
		{"admin@xapiens.id", "admin", "admin"},
		{"entry@xapiens.id", "entry", "entry"},
	}

	var user models.Users
	for _, v := range userArray {
		user.Email = v[0]
		user.Password = v[1]
		user.Role = v[2]
		user.ID = 0

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error -> ", err.Error())
		}
		user.Password = string(hash)

		db.Create(&user)

	}
	fmt.Println("Seeder Users created Sucessfully")
}
