package controller

import (
	"CasePoint3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SumGender func
func (strDB *StrDB) SumGender(c *gin.Context) {
	var person models.Persons
	var countM int64
	var countF int64

	gender := c.Query("gender")

	if gender == "M" {
		strDB.DB.Model(&person).Where("gender = ?", "M").Count(&countM)
		result :=
			gin.H{
				"meesage": "Data Laki Laki ",
				"count":   countM,
			}
		c.JSON(http.StatusFound, result)
	} else if gender == "F" {
		strDB.DB.Model(&person).Where("gender = ?", "F").Count(&countF)
		result :=
			gin.H{
				"meesage": "Data Perempuan",
				"count":   countF,
			}
		c.JSON(http.StatusFound, result)
	} else {
		result := gin.H{
			"meesage": "Please Search M or F",
		}
		c.JSON(http.StatusNotFound, result)
	}

}
