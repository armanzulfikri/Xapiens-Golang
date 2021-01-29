package controller

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

type PhoneBook struct {
	name   string `json:"name"`
	email  string `json:"email"`
	number string `json:"number"`
}

var listphone = []PhoneBook{}

func InputDatas(c *gin.Context) {
	var phoneBook PhoneBook
	var currentSave PhoneBook
	err := c.Bind(&phoneBook)
	if err != nil {
		fmt.Println("Something go wrong")
	}
	currentSave.name = phoneBook.name
	currentSave.email = phoneBook.email
	currentSave.number = phoneBook.number
	listphone = append(listphone, currentSave)
	fmt.Println(listphone)
	c.JSON(200, gin.H{
		"message": "Data Phone Book Successfully added",
		"name":    phoneBook.name,
		"email":   phoneBook.email,
		"number":  phoneBook.number,
	})

}

func UpdateNumber(c *gin.Context) {
	name := c.Param("id")
	number := c.PostForm("number")
	if !checkNomor(number) {
		fmt.Println("[Masukan Minimal 10 Digit dan Tidak Boleh Huruf!]")
	}

	for i := 0; i < len(listphone); i++ {
		if listphone[i].name == name {
			listphone[i].number = number
			c.JSON(200, gin.H{"message": "Number Successfully Updated"})
		} else {
			fmt.Println("Something Go Wrong")
		}
	}

}

func GetListBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All data Phone Books",
		"data":    listphone,
	})
}

func checkHuruf(name string) bool {
	isStringAlphabetic, _ := regexp.MatchString(`^[a-z A-Z]+$`, name)
	return isStringAlphabetic
}

func checkNomor(nomor string) bool {
	isStringNumeric, _ := regexp.MatchString(`^[\d]+$`, nomor)
	return isStringNumeric
}
