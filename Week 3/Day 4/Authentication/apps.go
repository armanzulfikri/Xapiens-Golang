package main

import (
	// "fmt"

	"pendudukApp/config"
	"pendudukApp/controller"

	"pendudukApp/midleware"

	"pendudukApp/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// connection to DB Postgree
	dbPG := config.Connect()
	strDB := controller.StrDB{DB: dbPG}

	//migration to DB Postgree
	models.Migrations(dbPG)

	//seeding data
	models.SeederProv(dbPG)
	models.SeederDistrict(dbPG)
	models.SeederSubDistrict(dbPG)
	models.SedderUser(dbPG)

	r := gin.Default()

	//route Province
	r.POST("/province", midleware.Auth, strDB.PostCreateProvince)
	r.GET("/province", midleware.Auth, strDB.GetListProvinceDistrictSubDistrictRename)
	r.GET("/province/find", midleware.Auth, strDB.GetOneProvince)
	r.PUT("/province/:id", midleware.Auth, strDB.UpdateProvince)
	r.DELETE("/province/delete/:id", midleware.Auth, strDB.DeleteProvince)

	//route District
	r.POST("/district", midleware.Auth, strDB.PostCreateDistrict)
	r.GET("/district", midleware.Auth, strDB.GetAllDistrict)
	r.GET("/district/find", midleware.Auth, strDB.GetOneDistricts)
	r.PUT("/district/:id", midleware.Auth, strDB.UpdateDistrict)
	r.DELETE("/district/delete/:id", midleware.Auth, strDB.DeleteDistrict)

	//route Subdistrict
	r.POST("/subdistrict", midleware.Auth, strDB.PostCreateSubDistrict)
	r.GET("/subdistrict", midleware.Auth, strDB.GetAllSubDistrict)
	r.GET("/subdistrict/find", midleware.Auth, strDB.GetOneSubDistricts)
	r.PUT("/subdistrict/:id", midleware.Auth, strDB.UpdateSubDistrict)
	r.DELETE("/subdistrict/delete/:id", midleware.Auth, strDB.DeleteSubDistrict)

	//route Person
	r.POST("/person", midleware.Auth, strDB.PostCreatePerson)
	r.GET("/person", midleware.Auth, strDB.GetAllPerson)
	r.GET("/person/find", midleware.Auth, strDB.GetOnePerson)
	r.PUT("/person/:id", midleware.Auth, strDB.UpdatePerson)
	r.DELETE("/person/delete/:id", midleware.Auth, strDB.DeletePerson)
	r.PATCH("/person/upload/:id", midleware.Auth, strDB.UpdateProfile)

	//user Router
	r.POST("/user/login", strDB.LoginUser)

	r.Run()
}
