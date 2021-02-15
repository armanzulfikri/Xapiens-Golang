package main

import (
	"fmt"
	"golangGraphql/config"
	"golangGraphql/models"
	"golangGraphql/seeders"
	"golangGraphql/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//main func
func main() {
	pgDB := config.Connect()

	// migration
	models.Migrations(pgDB)

	//Seeding From API Raja Ongkir
	seeders.SeedProvince(pgDB)
	seeders.SeedDisctrict(pgDB)
	seeders.SeedSubDistrict(pgDB)
	seeders.SeedPerson(pgDB)
	seeders.SeedOffice(pgDB)
	seeders.SeedOfficePersonLocation(pgDB)

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}
		var query Query

		c.Bind(&query)

		result := service.ExecuteQuery(query.Query, service.Schema)
		fmt.Println(result)
		c.JSON(http.StatusAccepted, result)
	})
	r.Run()
}
