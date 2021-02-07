package controller

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PostCreatePersonHandleOffice func
func (strDB *StrDB) PostCreatePersonHandleOffice(c *gin.Context) {
	var (
		personsOffices models.PersonHadleOfficeLocations
		result         gin.H
	)

	if err := c.Bind(&personsOffices); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&personsOffices)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":                 personsOffices.ID,
				"office_location_id": personsOffices.OfficeLocationID,
				"person_id":          personsOffices.PersonID,
				"created_at":         personsOffices.CreatedAt,
				"update_at":          personsOffices.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

//GetAllPersonHandleOffices func
func (strDB *StrDB) GetAllPersonHandleOffices(c *gin.Context) {
	var (
		personsOffices []models.PersonHadleOfficeLocations
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helper.Paging(&helper.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &personsOffices)

	c.JSON(http.StatusOK, paginator)
}

// GetOnePersonHandleOffices func
func (strDB *StrDB) GetOnePersonHandleOffices(c *gin.Context) {
	var (
		personsOffices []models.PersonHadleOfficeLocations
		result         gin.H
	)
	id := c.Query("ID")
	strDB.DB.Find(&personsOffices, id)
	if length := len(personsOffices); length <= 0 {
		result = ResponAPINil(personsOffices, length)
	} else {
		result = ResponAPI(personsOffices, length)
	}
	c.JSON(http.StatusOK, result)
}

//DeletePersonHandleOffice func
func (strDB *StrDB) DeletePersonHandleOffice(c *gin.Context) {
	var (
		PersonHadleOfficeLocations []models.PersonHadleOfficeLocations
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&PersonHadleOfficeLocations)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
