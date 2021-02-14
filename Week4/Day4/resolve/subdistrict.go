package resolve

import (
	"golangGraphql/config"
	"golangGraphql/models"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

//GetSubDistrict func
func GetSubDistrict(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()
	type SubDistrict struct {
		Name       string `json:"name"`
		ID         uint   `gorm:"primarykey" json:"id"`
		DistrictID uint   `gorm:"foreignkey" json:"district_id"`
	}

	var subDistrict []SubDistrict
	dbPG.Find(&subDistrict)
	return subDistrict, nil
}

//CreateSubDistrict func
func CreateSubDistrict(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var subDistrict models.SubDistricts

	subDistrict.ID = uint(rand.Intn(100000))
	subDistrict.DistrictID = params.Args["district_id"].(uint)
	subDistrict.Name = params.Args["name"].(string)

	db.Create(&subDistrict)
	return subDistrict, nil
}

//UpdateSubDistrict func
func UpdateSubDistrict(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)
	name, checkName := params.Args["name"].(string)
	districtID := params.Args["district_id"].(uint)

	var subDistrict []models.SubDistricts
	for i, v := range subDistrict {
		if uint(id) == v.ID {
			if checkName {
				subDistrict[i].Name = name
				subDistrict[i].DistrictID = districtID
			}
			db.Save(&subDistrict)
		}
	}
	return subDistrict, nil
}

//DeleteSubDistrict func
func DeleteSubDistrict(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)

	var subDistrict models.SubDistricts

	db.Delete(subDistrict, id)

	return subDistrict, nil
}
