package main

import (
	"CasePoint3/config"
	"CasePoint3/controller"
	"CasePoint3/midlewares"
	"CasePoint3/models"
	"CasePoint3/seeders"

	"github.com/gin-gonic/gin"
)

func main() {
	//postgre Connect
	pgDB := config.Connect()
	strDB := controller.StrDB{DB: pgDB}

	// migration
	models.Migrations(pgDB)

	//Seeding From API Raja Ongkir
	seeders.SeedProvince(pgDB)
	seeders.SeedDisctrict(pgDB)

	//Seeding Data
	seeders.SeedUser(pgDB)
	seeders.SeedSubDistrict(pgDB)
	seeders.SeedPerson(pgDB)
	seeders.SeedOffice(pgDB)
	seeders.SeedOfficePersonLocation(pgDB)

	r := gin.Default()

	//route Province
	r.POST("/province", midlewares.Corn, strDB.PostCreateProvince)
	r.GET("/province", midlewares.Auth, strDB.GetListProvinceDistrictSubDistrictRename)
	r.GET("/province/find", midlewares.Auth, strDB.GetOneProvince)
	r.PUT("/province/:id", midlewares.Corn, strDB.UpdateProvince)
	r.DELETE("/province/delete/:id", midlewares.Corn, strDB.DeleteProvince)

	//route District
	r.POST("/district", midlewares.Auth, strDB.PostCreateDistrict)
	r.GET("/district", midlewares.Auth, strDB.GetAllDistrict)
	r.GET("/district/find", midlewares.Auth, strDB.GetOneDistricts)
	r.PUT("/district/:id", midlewares.Auth, strDB.UpdateDistrict)
	r.DELETE("/district/delete/:id", midlewares.Auth, strDB.DeleteDistrict)

	//route Subdistrict
	r.POST("/subdistrict", midlewares.Auth, strDB.PostCreateSubDistrict)
	r.GET("/subdistrict", midlewares.Auth, strDB.GetAllSubDistrict)
	r.GET("/subdistrict/find", midlewares.Auth, strDB.GetOneSubDistricts)
	r.PUT("/subdistrict/:id", midlewares.Auth, strDB.UpdateSubDistrict)
	r.DELETE("/subdistrict/delete/:id", midlewares.Auth, strDB.DeleteSubDistrict)

	//route Person
	r.POST("/person", midlewares.Auth, strDB.PostCreatePerson)
	r.GET("/person", midlewares.Auth, strDB.GetAllPerson)
	r.GET("/person/find", midlewares.Auth, strDB.GetOnePerson)
	r.PUT("/person/:id", midlewares.Auth, strDB.UpdatePerson)
	r.DELETE("/person/delete/:id", midlewares.Auth, strDB.DeletePerson)
	r.PATCH("/person/upload/:id", midlewares.Auth, strDB.UpdateProfile)

	//route OfficeLocation
	r.POST("/office", midlewares.Auth, strDB.PostCreateOffice)
	r.GET("/office", midlewares.Auth, strDB.GetAllOffices)
	r.GET("/office/find", midlewares.Auth, strDB.GetOneOffices)
	r.PUT("/office/:id", midlewares.Auth, strDB.UpdateOffices)
	r.DELETE("/office/delete/:id", midlewares.Auth, strDB.DeleteOffices)

	//route PersonHandelOffice
	r.POST("/handleoffice", midlewares.Auth, strDB.PostCreatePersonHandleOffice)
	r.GET("/handleoffice", midlewares.Auth, strDB.GetAllPersonHandleOffices)
	r.GET("/handleoffice/find", midlewares.Auth, strDB.GetOnePersonHandleOffices)
	r.DELETE("/handleoffice/delete/:id", midlewares.Auth, strDB.DeletePersonHandleOffice)

	//user Router
	r.POST("/user/login", strDB.LoginUser)
	r.POST("/user/Register", strDB.Register)
	r.PUT("/user/:id", midlewares.Auth, strDB.UpdateOffices)
	r.DELETE("/user/delete/:id", midlewares.Auth, strDB.DeleteOffices)

	//Report Router

	r.GET("/report", strDB.SumGender)
	r.GET("/report/case2", strDB.Case2)
	r.Run()
}
