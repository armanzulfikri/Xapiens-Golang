package models

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

//SeederProv sample
func SeederProv(db *gorm.DB) {
	var provArray = [...]string{
		"Jawa Tengah",
		"Jawa Timur",
		"Jawa Barat",
	}
	var prov Provinces
	for _, v := range provArray {
		prov.Name = v
		prov.ID = 0
		db.Create(&prov)
	}
	fmt.Println("Seeder Prov created")
}

//SeederDistrict sample
func SeederDistrict(db *gorm.DB) {
	var districtArray = [...][2]string{
		{"1", "Banyumas"},
		{"1", "Solo"},
		{"1", "Semarang"},
		{"2", "Banyuwangi"},
		{"2", "Surabaya"},
		{"3", "Bandung"},
		{"3", "Jakarta"},
	}

	var district Districts
	for _, v1 := range districtArray {
		data, _ := strconv.ParseUint(v1[0], 10, 64)
		district.ProvinceID = uint(data)
		district.Name = v1[1]
		district.ID = 0
		db.Create(&district)

	}
	fmt.Println("Seeder District created")
}

//SeederSubDistrict sample
func SeederSubDistrict(db *gorm.DB) {
	var subDistrictArray = [...][2]string{
		{"1", "Kec di Banyumas"},
		{"2", "Kec di Solo"},
		{"3", "Kec di Semarang"},
		{"4", "Kec di Banyuwangi"},
		{"5", "Kec di Surabaya"},
		{"6", "Kec di Bandung"},
		{"7", "Kec di Jakarta"},
	}

	var subdistrict SubDistricts
	for _, v1 := range subDistrictArray {
		data, _ := strconv.ParseUint(v1[0], 10, 64)
		subdistrict.DistrictID = uint(data)
		subdistrict.Name = v1[1]
		subdistrict.ID = 0
		db.Create(&subdistrict)

	}
	fmt.Println("Seeder Sub District created")
}

// //sedder person
// func SeederPerson(db *gorm.DB) {
// 	var subDistrictArray = [...][2]string{
// 		{"1", "Kec di Banyumas"},
// 		{"2", "Kec di Solo"},
// 		{"3", "Kec di Semarang"},
// 		{"4", "Kec di Banyuwangi"},
// 		{"5", "Kec di Surabaya"},
// 		{"6", "Kec di Bandung"},
// 		{"7", "Kec di Jakarta"},
// 	}

// 	var subdistrict SubDistricts
// 	for _, v1 := range subDistrictArray {
// 		data, _ := strconv.ParseUint(v1[0], 10, 64)
// 		subdistrict.DistrictsID = uint(data)
// 		subdistrict.Name = v1[1]
// 		subdistrict.ID = 0
// 		db.Create(&subdistrict)

// 	}
// 	fmt.Println("Seeder Sub District created")
// }
