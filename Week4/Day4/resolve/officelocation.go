package resolve

import (
	"golangGraphql/config"
	"golangGraphql/models"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

//GetOfficeLocation func
func GetOfficeLocation(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()
	type OfficeLocation struct {
		Name          string `json:"name"`
		ID            uint   `gorm:"primarykey" json:"id"`
		SubDistrictID uint   `gorm:"foreignkey:SubDistrictID" json:"subdistrict"`
	}

	var officeLocation []OfficeLocation
	dbPG.Find(&officeLocation)
	return officeLocation, nil
}

//CreateOfficeLocation func
func CreateOfficeLocation(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var officeLocation models.OfficesLocation

	officeLocation.ID = uint(rand.Intn(100000))
	officeLocation.SubDistrictID = params.Args["subdistrict_id"].(uint)
	officeLocation.Name = params.Args["name"].(string)

	db.Create(&officeLocation)
	return officeLocation, nil
}

//UpdateOfficeLocation func
func UpdateOfficeLocation(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)
	name, checkName := params.Args["name"].(string)
	subDistrictID := params.Args["subdistrict_id"].(uint)

	var officeLocation []models.OfficesLocation
	for i, v := range officeLocation {
		if uint(id) == v.ID {
			if checkName {
				officeLocation[i].Name = name
				officeLocation[i].SubDistrictID = subDistrictID
			}
			db.Save(&officeLocation)
		}
	}
	return officeLocation, nil
}

//DeleteOfficeLocation func
func DeleteOfficeLocation(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)

	var officeLocation models.OfficesLocation

	db.Delete(officeLocation, id)

	return officeLocation, nil
}
