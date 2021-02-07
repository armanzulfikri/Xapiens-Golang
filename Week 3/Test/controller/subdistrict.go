package controller

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PostCreateSubDistrict func
func (strDB *StrDB) PostCreateSubDistrict(c *gin.Context) {
	var (
		subdistricts models.SubDistricts
		result       gin.H
	)

	if err := c.Bind(&subdistricts); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&subdistricts)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":          subdistricts.ID,
				"district_Id": subdistricts.DistrictID,
				"name":        subdistricts.Name,
				"created_at":  subdistricts.CreatedAt,
				"update_at":   subdistricts.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

//GetAllSubDistrict func
func (strDB *StrDB) GetAllSubDistrict(c *gin.Context) {
	var (
		subdistricts []models.SubDistricts
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helper.Paging(&helper.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
		Join:    "left join districts on districts.id = sub_districts.districts_id left join provinces on provinces.id = districts.province_id",
		Query:   "sub_districts.id, sub_districts.name, districts.name as district, provinces.name as province",
	}, &subdistricts)

	c.JSON(http.StatusOK, paginator)
}

//GetOneSubDistricts func
func (strDB *StrDB) GetOneSubDistricts(c *gin.Context) {
	var (
		subdistricts []models.SubDistricts
		result       gin.H
	)
	id := c.Query("id")
	strDB.DB.First(&subdistricts, id)
	if length := len(subdistricts); length <= 0 {
		result = ResponAPINil(subdistricts, length)
	} else {
		result = ResponAPI(subdistricts, length)
	}
	c.JSON(http.StatusOK, result)
}

//UpdateSubDistrict func
func (strDB *StrDB) UpdateSubDistrict(c *gin.Context) {
	var (
		subdistricts models.SubDistricts
		result       gin.H
	)
	id := c.Param("id")
	name := c.PostForm("name")
	if err := c.Bind(&subdistricts); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).Find(&subdistricts)
		subdistricts.Name = name

		strDB.DB.Save(&subdistricts)
		c.JSON(http.StatusOK, result)
	}
}

//DeleteSubDistrict func
func (strDB *StrDB) DeleteSubDistrict(c *gin.Context) {
	var (
		subdistricts []models.SubDistricts
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&subdistricts)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
