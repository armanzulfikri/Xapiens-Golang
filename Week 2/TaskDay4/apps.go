package main

import (
	// "fmt"

	"go-orm/config"
	"go-orm/controller"
	"go-orm/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// connection to DB Postgree
	dbPG := config.Connect()
	strDB := controller.StrDB{DB: dbPG}

	//migration to DB Postgree
	models.Migrations(dbPG)

	r := gin.Default()

	//route Province
	r.POST("/province", strDB.PostCreateProvince)
	r.GET("/province", strDB.GetReadProvince)
	r.PUT("/province/:id", strDB.UpdateProvince)
	r.DELETE("/province/delete/:id", strDB.DeleteProvince)

	//route District
	r.POST("/district", strDB.PostCreateDistrict)
	r.GET("/district", strDB.GetReadDistrict)
	r.PUT("/district/:id", strDB.UpdateDistrict)
	r.DELETE("/district/delete/:id", strDB.DeleteDistrict)

	r.Run()
}
