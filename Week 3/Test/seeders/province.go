package seeders

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/gorm"
)

// SeedProvince func
func SeedProvince(db *gorm.DB) {
	//Scraping Data From Raja Ongkir
	baseURL := "https://api.rajaongkir.com/starter/province"
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
		response models.RajaongkirProvince
		province models.Provinces
	)

	if err := json.Unmarshal(resBody, &response); err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	for i := 0; i < len(response.RajaOngkir.ProvinceResults); i++ {
		province.Name = response.RajaOngkir.ProvinceResults[i].Province
		province.ID = 0
		db.Create(&province)
	}
	fmt.Println("Seeder Province created Sucessfully")
}
