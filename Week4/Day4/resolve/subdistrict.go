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
	subDistrict.DistrictID = uint(params.Args["district_id"].(int))
	subDistrict.Name = params.Args["name"].(string)

	db.Create(&subDistrict)
	return subDistrict, nil
}

//UpdateSubDistrict func
func UpdateSubDistrict(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := params.Args["id"].(int)
	name := params.Args["name"].(string)
	districtID := uint(params.Args["district_id"].(int))

	var subDistrict models.SubDistricts

	db.Model(&subDistrict).Where("id = ?", id).Updates(models.SubDistricts{Name: name, DistrictID: districtID})
	return subDistrict, nil
}

//DeleteSubDistrict func
func DeleteSubDistrict(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var subDistrict models.SubDistricts

	db.Delete(&subDistrict, id)

	return subDistrict, nil
}
