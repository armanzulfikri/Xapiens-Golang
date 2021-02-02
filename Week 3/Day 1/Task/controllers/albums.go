package controllers

import (
	"Scraping/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//PostData From API Commnets
func (strDB *StrDB) PostDataFromAlbumAPI() {
	var albums models.Albums

	res, err := http.Get("https://jsonplaceholder.typicode.com/albums/1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)

	if error := json.Unmarshal([]byte(stringBody), &albums); error != nil {
		fmt.Println("error", error.Error())
	}

	strDB.DB.Create(&albums)
	fmt.Println(albums.UserId)
	fmt.Println(albums.ID)
	fmt.Println(albums.Title)
}

//Get Data Comments
func (strDB *StrDB) GetAlbums() {
	var albums models.Albums

	result := strDB.DB.Find(&albums)
	fmt.Println(result)
}
