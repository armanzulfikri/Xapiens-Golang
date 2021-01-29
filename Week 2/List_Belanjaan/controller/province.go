package controller

import (
	"Belanjaan_struct/models"
	"fmt"
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
			"coba":    province.Name,
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
