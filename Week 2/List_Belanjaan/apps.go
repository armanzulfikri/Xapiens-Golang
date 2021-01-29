package main

import (
	// "fmt"

	"Belanjaan_struct/config"
	"Belanjaan_struct/controller"
	"Belanjaan_struct/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// connection to DB Postgree
	dbPG := config.Connect()
	strDB := controller.StrDB{DB: dbPG}

	//migration to DB Postgree
	models.Migrations(dbPG)

	r := gin.Default()

	r.POST("/province", strDB.PostCreateProvince)
	r.GET("/province", strDB.GetReadProvince)
	// var shop controller.ItemBelanjaan
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"msg": "Welcome to Shopping Apps",
	// 	})
	// })

	// r.GET("/shop/", shop.GetList)
	// r.POST("/shop/create", shop.CreateItem)
	// r.PUT("/shop/update", shop.UpdateListShopping)

	r.Run()
}
