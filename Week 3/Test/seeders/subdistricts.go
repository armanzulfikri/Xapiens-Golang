package seeders

import (
	"CasePoint3/models"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

//SeedSubDistrict func
func SeedSubDistrict(db *gorm.DB) {
	var SeedSubDistrictArray = [...][2]string{
		{"42", "GentengKulon"},
		{"42", "Muncar"},
		{"42", "Songgon"},
		{"42", "Genteng Wetan"},
		{"42", "Banyuwangi Kota"},
		{"48", "Nongsa"},
		{"21", "Bandar Mataram"},
		{"22", "Antapani"},
		{"23", "Bandung Kidul"},
		{"22", "Bandung Kulon"},
		{"23", "Arcamanik"},
		{"22", "Buahbatu"},
	}

	var subDistrict models.SubDistricts
	for _, v := range SeedSubDistrictArray {
		data, _ := strconv.ParseUint(v[0], 10, 32)
		subDistrict.DistrictID = uint(data)
		subDistrict.Name = v[1]
		subDistrict.ID = 0
		db.Create(&subDistrict)

	}
	fmt.Println("Seeder SubDistrict created Sucessfully")
}
