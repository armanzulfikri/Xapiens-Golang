package seeders

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

//RajaOngkirCity models
type RajaOngkirCity struct {
	RajaOngkir RajaOngkirCityResponse `json:"rajaongkir"`
}

//RajaOngkirCityResponse model
type RajaOngkirCityResponse struct {
	CityResults []CityResults `json:"results"`
}

//CityResults model
type CityResults struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

//SeedDisctrict func
func SeedDisctrict(db *gorm.DB) {
	baseURL := "https://api.rajaongkir.com/starter/city"
	apiKey := helper.GetEnv("RAJAONGKIR_APIKEY")
	res, err := http.Get(baseURL + "?key=" + apiKey)
	if err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	var (
		response RajaOngkirCity
		dist     models.Districts
	)

	if err := json.Unmarshal(resBody, &response); err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	for i := 0; i < 50; i++ {
		uProvID, _ := strconv.ParseUint(response.RajaOngkir.CityResults[i].ProvinceID, 10, 32)
		dist.Name = response.RajaOngkir.CityResults[i].CityName
		dist.ProvinceID = uint(uProvID)
		dist.ID = 0
		db.Create(&dist)
	}
	fmt.Println("Seeder District created Sucessfully")

}
