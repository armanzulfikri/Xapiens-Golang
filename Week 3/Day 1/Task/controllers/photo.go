package controllers

import (
	"Scraping/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//PostData From API Photos
func (strDB *StrDB) PostDataFromPhotosAPI() {
	var photos models.Photos

	res, err := http.Get("https://jsonplaceholder.typicode.com/photos/1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)

	if error := json.Unmarshal([]byte(stringBody), &photos); error != nil {
		fmt.Println("error", error.Error())
	}

	strDB.DB.Create(&photos)
}

//Get Data Photos
func (strDB *StrDB) GetPhotos() {
	var photos models.Photos

	result := strDB.DB.Find(&photos)
	fmt.Println(result)
}
