package seeders

import (
	"fmt"
	"golangGraphql/models"
	"strconv"

	"gorm.io/gorm"
)

// SeedOffice func
func SeedOffice(db *gorm.DB) {
	var officeArray = [...][2]string{
		{"1", "PT Cendrwasih"},
		{"2", "PrivyID"},
		{"3", "PT Xapiens Teknologi Indonesia"},
	}

	var office models.OfficesLocation
	for _, v := range officeArray {
		data, _ := strconv.ParseUint(v[0], 10, 32)
		office.SubDistrictID = uint(data)
		office.Name = v[1]
		office.ID = 0
		db.Create(&office)

	}
	fmt.Println("Seeder OfficeLocation created Sucefully")
}
