package controller

import (
	"fmt"
	"go-orm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostCreateProvince route struct method
func (strDB *StrDB) PostCreateProvince(c *gin.Context) {
	var (
		province models.Provinces
		result   gin.H
	)

	if err := c.Bind(&province); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&province)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":         province.ID,
				"name":       province.Name,
				"created_at": province.CreatedAt,
				"update_at":  province.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// Get Province
func (strDB *StrDB) GetReadProvince(c *gin.Context) {
	var (
		province []models.Provinces
		result   gin.H
	)

	strDB.DB.Find(&province)
	if length := len(province); length <= 0 {
		result = ResponAPINil(province, length)
	} else {
		result = ResponAPI(province, length)
	}

	c.JSON(http.StatusOK, result)
}

//update Province
func (strDB *StrDB) UpdateProvince(c *gin.Context) {
	var (
		province []models.Provinces
		result   gin.H
	)
	id := c.Param("id")
	if err := c.Bind(&province); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).First(&province)
		result = gin.H{
			"message": "success Update Data",
		}
		strDB.DB.Save(&province)
		c.JSON(http.StatusOK, result)
	}
}

//Delete Province
func (strDB *StrDB) DeleteProvince(c *gin.Context) {
	var (
		province []models.Provinces
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&province)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
