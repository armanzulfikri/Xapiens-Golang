package main

import (
	// "fmt"

	"pendudukApp/config"
	"pendudukApp/controller"
	"pendudukApp/models"

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
	r.GET("/province/find/", strDB.GetOneProvince)
	r.PUT("/province/:id", strDB.UpdateProvince)
	r.DELETE("/province/delete/:id", strDB.DeleteProvince)

	//route District
	r.POST("/district", strDB.PostCreateDistrict)
	r.GET("/district", strDB.GetAllDistrict)
	r.GET("/district/:id", strDB.GetOneDistricts)
	r.PUT("/district/:id", strDB.UpdateDistrict)
	r.DELETE("/district/delete/:id", strDB.DeleteDistrict)

	//route Subdistrict
	r.POST("/subdistrict", strDB.PostCreateSubDistrict)
	r.GET("/subdistrict", strDB.GetAllSubDistrict)
	r.GET("/subdistrict/:id", strDB.GetOneSubDistricts)
	r.PUT("/subdistrict/:id", strDB.UpdateSubDistrict)
	r.DELETE("/subdistrict/delete/:id", strDB.DeleteSubDistrict)

	//route Person
	r.POST("/person", strDB.PostCreatePerson)
	r.GET("/person", strDB.GetAllPerson)
	r.GET("/person/:id", strDB.GetOnePerson)
	r.PUT("/person/:id", strDB.UpdatePerson)
	r.DELETE("/person/delete/:id", strDB.DeletePerson)
	r.Run()
}
