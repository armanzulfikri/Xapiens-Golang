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

// Get District route struct method
func (strDB *StrDB) GetAllDistrict(c *gin.Context) {
	var (
		districts []models.Districts
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helpers.Paging(&helpers.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
		Join:    "left join provinces on provinces.id = districts.province_id",
		Query:   "districts.id, districts.name, provinces.name as province",
	}, &districts)

	c.JSON(http.StatusOK, paginator)
}

func (strDB *StrDB) GetOneDistricts(c *gin.Context) {
	var (
		districts []models.Districts
		result    gin.H
	)
	id := c.Query("id")
	strDB.DB.First(&districts, id)
	if length := len(districts); length <= 0 {
		result = ResponAPINil(districts, length)
	} else {
		result = ResponAPI(districts, length)
	}
	c.JSON(http.StatusOK, result)
}

//Update District
func (strDB *StrDB) UpdateDistrict(c *gin.Context) {
	var (
		districts []models.Districts
		result    gin.H
	)
	id := c.Param("id")
	if err := c.Bind(&districts); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).First(&districts)
		result = gin.H{
			"message": "success Update Data",
		}
		strDB.DB.Save(&districts)
		c.JSON(http.StatusOK, result)
	}
}

//Delete District
func (strDB *StrDB) DeleteDistrict(c *gin.Context) {
	var (
		districts []models.Districts
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&districts)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
