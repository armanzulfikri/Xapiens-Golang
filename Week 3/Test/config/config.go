package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var userDatabase, passDatabase, portDatabase,
		hostDatabase, nameDatabase, sslDatabase, timezoneDatabase string

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env file")
	} else {
		userDatabase = os.Getenv("DATABASE_USER")
		passDatabase = os.Getenv("DATABASE_PASS")
		portDatabase = os.Getenv("DATABASE_PORT")
		hostDatabase = os.Getenv("DATABASE_HOST")
		nameDatabase = os.Getenv("DATABASE_NAME")
		sslDatabase = os.Getenv("DATABASE_SSLMODE")
		timezoneDatabase = os.Getenv("DATABASE_TIMEZONE")

	}
	conn :=
		" host=" + hostDatabase +
			" user=" + userDatabase +
			" password=" + passDatabase +
			" dbname=" + nameDatabase +
			" port=" + portDatabase +
			" sslmode=" + sslDatabase +
			" Timezone=" + timezoneDatabase

	db, errConn := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if errConn != nil {
		panic("failed to connect to database")
	} else {
		fmt.Println("Koneksi Sukses")
	}
	return db
}
