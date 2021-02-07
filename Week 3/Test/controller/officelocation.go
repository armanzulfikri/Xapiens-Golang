package controller

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostCreateOffice route struct method
func (strDB *StrDB) PostCreateOffice(c *gin.Context) {
	var (
		offices models.OfficesLocation
		result  gin.H
	)

	if err := c.Bind(&offices); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&offices)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":               offices.ID,
				"sub_districts_id": offices.SubDistrictID,
				"name":             offices.Name,
				"created_at":       offices.CreatedAt,
				"update_at":        offices.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

//GetAllOffices func
func (strDB *StrDB) GetAllOffices(c *gin.Context) {
	var (
		offices []models.OfficesLocation
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helper.Paging(&helper.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &offices)

	c.JSON(http.StatusOK, paginator)
}

// GetOneOffices func
func (strDB *StrDB) GetOneOffices(c *gin.Context) {
	var (
		offices []models.OfficesLocation
		result  gin.H
	)
	id := c.Query("ID")
	strDB.DB.Find(&offices, id)
	if length := len(offices); length <= 0 {
		result = ResponAPINil(offices, length)
	} else {
		result = ResponAPI(offices, length)
	}
	c.JSON(http.StatusOK, result)
}

//UpdateOffices func
func (strDB *StrDB) UpdateOffices(c *gin.Context) {
	var (
		offices models.OfficesLocation
		result  gin.H
	)
	id := c.Param("id")
	name := c.PostForm("name")
	if err := c.Bind(&offices); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).Find(&offices)
		offices.Name = name

		result = gin.H{
			"message": "success Update Data",
		}
		strDB.DB.Save(&offices)
		c.JSON(http.StatusOK, result)
	}
}

//DeleteOffices func
func (strDB *StrDB) DeleteOffices(c *gin.Context) {
	var (
		office []models.OfficesLocation
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&office)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
