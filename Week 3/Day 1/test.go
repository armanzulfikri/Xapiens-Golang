package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type DataPost struct {
		UserID    uint   `json:"userId"`
		ID        uint   `json: "id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	fmt.Println(body)
	sb := string(body)
	fmt.Println(sb)
	dataPosts := DataPost{}
	if error := json.Unmarshal([]byte(sb), &dataPosts); error != nil {
		fmt.Println("error ", error.Error())
	}

	fmt.Println(dataPosts.Title)

}
