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
func (strDB *StrDB) PostDataFromCommentsAPI() {
	var comments models.Comments

	res, err := http.Get("https://jsonplaceholder.typicode.com/comments/1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)

	if error := json.Unmarshal([]byte(stringBody), &comments); error != nil {
		fmt.Println("error", error.Error())
	}

	strDB.DB.Create(&comments)
}

//Get Data Comments
func (strDB *StrDB) GetComments() {
	var comments models.Comments

	result := strDB.DB.Find(&comments)
	fmt.Println(result)
}
