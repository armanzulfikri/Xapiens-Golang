package resolve

import (
	"golangGraphql/config"
	"golangGraphql/models"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

//GetProvince func
func GetProvince(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()
	type Provinces struct {
		Name string `json:"name"`
		ID   uint   `gorm:"primarykey" json:"id"`
	}
	var (
		provinces []Provinces
	)
	dbPG.Preload("Districts").Find(&provinces)
	return provinces, nil
}

//GetProvinceDistrict func
func GetProvinceDistrict(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()
	type Districts struct {
		ID         uint   `gorm:"primarykey" json:"id"`
		Name       string `json:"name"`
		ProvinceID uint   `json:"province_id`
	}
	type Provinces struct {
		Name      string      `json:"name"`
		Districts []Districts `gorm:"foreignKey:ProvinceID" json:"districts"`
		ID        uint        `gorm:"primarykey" json:"id"`
	}
	var (
		provinces []Provinces
	)
	dbPG.Preload("Districts").Find(&provinces)

	return provinces, nil
}

//CreateProvince func
func CreateProvince(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var province models.Provinces

	province.ID = uint(rand.Intn(100000))
	province.Name = params.Args["name"].(string)

	db.Create(&province)
	return province, nil
}

//UpdateProvince func
func UpdateProvince(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	name := params.Args["name"].(string)

	var provinces models.Provinces
	db.Model(&provinces).Where("id =  ?", id).Update("name", name)

	return provinces, nil
}

//DeleteProvince func
func DeleteProvince(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var provinces models.Provinces

	db.Where("id = ?", id).Delete(&provinces, id)

	return provinces, nil
}
