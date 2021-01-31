package controller

import (
	"fmt"
	"net/http"
	"pendudukApp/helpers"
	"pendudukApp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostCreateDistrict route struct method
func (strDB *StrDB) PostCreatePerson(c *gin.Context) {
	var (
		persons models.Persons
		result  gin.H
	)

	if err := c.Bind(&persons); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&persons)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":              persons.ID,
				"subdistricts_Id": persons.SubDistrictsID,
				"nip":             persons.Nip,
				"full_name":       persons.FullName,
				"last_name":       persons.LastName,
				"birth_date":      persons.BirthDate,
				"birth_place":     persons.BirthPlace,
				"gender":          persons.Gender,
				"zona_location":   persons.ZonaLocation,
				"created_at":      persons.CreatedAt,
				"update_at":       persons.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// Get District route struct method
func (strDB *StrDB) GetAllPerson(c *gin.Context) {
	var (
		persons []models.Persons
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helpers.Paging(&helpers.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
		Join:    "left join subdistricts on subdistricts.id = person.subdistricts_id",
		Query:   "person.id, person.name, subdistricts.name as subdistricts",
	}, &persons)

	c.JSON(http.StatusOK, paginator)
}

func (strDB *StrDB) GetOnePerson(c *gin.Context) {
	var (
		persons []models.Persons
		result  gin.H
	)
	id := c.Query("id")
	strDB.DB.First(&persons, id)
	if length := len(persons); length <= 0 {
		result = ResponAPINil(persons, length)
	} else {
		result = ResponAPI(persons, length)
	}
	c.JSON(http.StatusOK, result)
}

//Update District
func (strDB *StrDB) UpdatePerson(c *gin.Context) {
	var (
		persons []models.Persons
		result  gin.H
	)
	id := c.Param("id")
	if err := c.Bind(&persons); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).First(&persons)
		result = gin.H{
			"message": "success Update Data",
		}
		strDB.DB.Save(&persons)
		c.JSON(http.StatusOK, result)
	}
}

//Delete District
func (strDB *StrDB) DeletePerson(c *gin.Context) {
	var (
		persons []models.Persons
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&persons)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
