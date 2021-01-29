package controller

import (
	"fmt"
	"go-orm/models"
	"net/http"

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
func (strDB *StrDB) GetReadDistrict(c *gin.Context) {
	var (
		districts []models.Districts
		result    gin.H
	)

	strDB.DB.Find(&districts)
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
