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
	fmt.Println(stringBody)

	if error := json.Unmarshal([]byte(stringBody), &comments); error != nil {
		fmt.Println("error", error.Error())
	}

	strDB.DB.Create(&comments)

}

//Get Data Comments
func (strDB *StrDB) GetComments() {
	var comments models.Comments

	strDB.DB.Find(&comments)
	fmt.Println(comments.PostId)
	fmt.Println(comments.ID)
	fmt.Println(comments.Name)
	fmt.Println(comments.Email)
	fmt.Println(comments.Body)

}
