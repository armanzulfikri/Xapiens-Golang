package resolve

import (
	"golangGraphql/config"
	"golangGraphql/models"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

//GetDistrict func
func GetDistrict(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()
	type Districts struct {
		Name       string `json:"name"`
		ID         uint   `gorm:"primarykey" json:"id"`
		ProvinceID uint   `gorm:"foreignkey" json:"ProvinceID"`
	}

	var district []Districts
	dbPG.Find(&district)
	return district, nil
}

//CreateDistrict func
func CreateDistrict(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var districts models.Districts

	districts.ID = uint(rand.Intn(100000))
	districts.ProvinceID = uint(64)
	districts.Name = params.Args["name"].(string)

	db.Create(&districts)
	return districts, nil
}

//UpdateDistricts func
func UpdateDistricts(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)
	name, checkName := params.Args["name"].(string)
	provinceID := params.Args["province_id"].(uint)

	var districts []models.Districts
	for i, v := range districts {
		if uint(id) == v.ID {
			if checkName {
				districts[i].Name = name
				districts[i].ProvinceID = provinceID
			}
			db.Save(&districts)
		}
	}
	return districts, nil
}

//DeleteDistricts func
func DeleteDistricts(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)

	var districts models.Districts

	db.Delete(districts, id)

	return districts, nil
}
