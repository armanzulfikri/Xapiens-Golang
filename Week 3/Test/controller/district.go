package controller

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostCreateDistrict func
func (strDB *StrDB) PostCreateDistrict(c *gin.Context) {
	var (
		districts models.Districts
		result    gin.H
	)

	if err := c.Bind(&districts); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&districts)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":          districts.ID,
				"province_Id": districts.ProvinceID,
				"name":        districts.Name,
				"created_at":  districts.CreatedAt,
				"update_at":   districts.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

//GetAllDistrict func
func (strDB *StrDB) GetAllDistrict(c *gin.Context) {
	var (
		districts []models.Districts
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helper.Paging(&helper.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &districts)

	c.JSON(http.StatusOK, paginator)
}

//GetOneDistricts func
func (strDB *StrDB) GetOneDistricts(c *gin.Context) {
	var (
		districts []models.Districts
		result    gin.H
	)
	id := c.Query("ID")
	strDB.DB.Find(&districts, id)
	if length := len(districts); length <= 0 {
		result = ResponAPINil(districts, length)
	} else {
		result = ResponAPI(districts, length)
	}
	c.JSON(http.StatusOK, result)
}

//UpdateDistrict func
func (strDB *StrDB) UpdateDistrict(c *gin.Context) {
	var (
		districts models.Districts
		result    gin.H
	)
	id := c.Param("id")
	name := c.PostForm("name")
	if err := c.Bind(&districts); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).Find(&districts)
		districts.Name = name

		result = gin.H{
			"message": "success Update Data",
		}
		strDB.DB.Save(&districts)
		c.JSON(http.StatusOK, result)
	}
}

//DeleteDistrict func
func (strDB *StrDB) DeleteDistrict(c *gin.Context) {
	var (
		districts []models.Districts
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&districts)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
