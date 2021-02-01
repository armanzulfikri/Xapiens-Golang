package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect with gorm
func Connect() *gorm.DB {
	var userDB, pwDB, portDB, hostDB, nameDB string
	userDB = "postgres"
	pwDB = "root"
	portDB = "5432"
	hostDB = "localhost"
	nameDB = "ScrapingAPI"

	conn := " host=" + hostDB +
		" user=" + userDB +
		" password=" + pwDB +
		" dbname=" + nameDB +
		" port=" + portDB +
		" sslmode=disable TimeZone=Asia/Shanghai"

	db, errConn := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if errConn != nil {
		panic("failed to connect to the database")
	} else {
		fmt.Println("successful connection")
	}
	return db
}
