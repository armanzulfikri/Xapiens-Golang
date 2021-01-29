package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var userDatabase, passDatabase, portDatabase, hostDatabase, nameDatabase string

	userDatabase = "postgres"
	passDatabase = "root"
	portDatabase = "5432"
	hostDatabase = "localhost"
	nameDatabase = "penduduk"

	conn :=
		" host=" + hostDatabase +
			" user=" + userDatabase +
			" password=" + passDatabase +
			" dbname=" + nameDatabase +
			" port=" + portDatabase +
			" sslmode=disable Timezone=Asia/Shanghai"
	db, errConn := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if errConn != nil {
		panic("failed to connect to database")
	} else {
		fmt.Println("Koneksi Sukses")
	}
	return db
}
