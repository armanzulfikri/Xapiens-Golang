package controller

import (
	"CasePoint3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SumGender func
func (strDB *StrDB) SumGender(c *gin.Context) {
	var person models.Persons
	query := "SELECT COUNT(gender) FROM persons"
	// pria := strDB.DB.Model(&person).Where("gender = ?", "M").Count(&count)
	wanita := strDB.DB.Model(&person).Select(query)
	result :=
		gin.H{
			"meesage": "Data Find",

			"Wanita": wanita,
		}
	c.JSON(http.StatusOK, result)
}
