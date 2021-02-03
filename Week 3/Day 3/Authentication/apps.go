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
	r.POST("/province", midleware.Authentication, strDB.PostCreateProvince)
	r.GET("/province", midleware.Authentication, strDB.GetListProvinceDistrictSubDistrictRename)
	r.GET("/province/find", midleware.Authentication, strDB.GetOneProvince)
	r.PUT("/province/:id", midleware.Authentication, strDB.UpdateProvince)
	r.DELETE("/province/delete/:id", midleware.Authentication, strDB.DeleteProvince)

	//route District
	r.POST("/district", midleware.Authentication, strDB.PostCreateDistrict)
	r.GET("/district", midleware.Authentication, strDB.GetAllDistrict)
	r.GET("/district/find", midleware.Authentication, strDB.GetOneDistricts)
	r.PUT("/district/:id", midleware.Authentication, strDB.UpdateDistrict)
	r.DELETE("/district/delete/:id", midleware.Authentication, strDB.DeleteDistrict)

	//route Subdistrict
	r.POST("/subdistrict", midleware.Authentication, strDB.PostCreateSubDistrict)
	r.GET("/subdistrict", midleware.Authentication, strDB.GetAllSubDistrict)
	r.GET("/subdistrict/find", midleware.Authentication, strDB.GetOneSubDistricts)
	r.PUT("/subdistrict/:id", midleware.Authentication, strDB.UpdateSubDistrict)
	r.DELETE("/subdistrict/delete/:id", midleware.Authentication, strDB.DeleteSubDistrict)

	//route Person
	r.POST("/person", midleware.Authentication, strDB.PostCreatePerson)
	r.GET("/person", midleware.Authentication, strDB.GetAllPerson)
	r.GET("/person/find", midleware.Authentication, strDB.GetOnePerson)
	r.PUT("/person/:id", midleware.Authentication, strDB.UpdatePerson)
	r.DELETE("/person/delete/:id", midleware.Authentication, strDB.DeletePerson)
	r.PATCH("/person/upload/:id", midleware.Authentication, strDB.UpdateProfile)

	//user Router
	r.POST("/user/login", strDB.LoginUser)

	r.Run()
}
