package seeders

import (
	"CasePoint3/models"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

// SeedOfficePersonLocation func
func SeedOfficePersonLocation(db *gorm.DB) {
	var ofcPersonArr = [...][2]string{
		{"1", "1"},
		{"2", "2"},
		{"3", "3"},
		{"1", "4"},
		{"2", "5"},
		{"3", "6"},
		{"1", "7"},
		{"2", "8"},
		{"3", "9"},
		{"1", "10"},
		{"2", "11"},
		{"3", "12"},
		{"1", "1"},
		{"2", "2"},
		{"3", "3"},
		{"1", "4"},
		{"2", "5"},
		{"3", "6"},
		{"1", "7"},
		{"2", "8"},
	}

	var officePerson models.PersonHadleOfficeLocations
	for _, v := range ofcPersonArr {
		office, _ := strconv.ParseUint(v[0], 10, 32)
		officePerson.OfficeLocationID = uint(office)
		person, _ := strconv.ParseUint(v[1], 10, 32)
		officePerson.PersonID = uint(person)
		officePerson.ID = 0
		db.Create(&officePerson)

	}
	fmt.Println("Seeder PersonHandelOffice created Sucessfully")
}
