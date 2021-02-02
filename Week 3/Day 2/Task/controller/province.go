package controller

import (
	"fmt"
	"net/http"
	"pendudukApp/helpers"
	"pendudukApp/models"
	"strconv"

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
	)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	strDB.DB.Preload("Districts").Find(&province)
	paginator := helpers.Paging(&helpers.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &province)
	c.JSON(http.StatusOK, paginator)
}

func (strDB *StrDB) GetOneProvince(c *gin.Context) {
	var (
		province []models.Provinces
		result   gin.H
	)
	id := c.Query("id")
	strDB.DB.
		Preload("Districts").
		Preload("Districts.SubDistricts").
		Find(&province, id)
	if length := len(province); length <= 0 {
		result = ResponAPINil(province, length)
	} else {
		result = ResponAPI(province, length)
	}
	c.JSON(http.StatusOK, result)
}

// GetListProvinceDistrictSubDistrictRename comment
func (strDB *StrDB) GetListProvinceDistrictSubDistrictRename(c *gin.Context) {
	type SubDistricts struct {
		Name       string `json:"name"`
		ID         uint   `gorm:"primarykey" json:"id"`
		DistrictID uint   `json:"district_id"`
	}
	type Districts struct {
		Name         string         `json:"name"`
		ID           uint           `gorm:"primarykey" json:"id"`
		ProvinceID   uint           `json:"province_id"`
		SubDistricts []SubDistricts `gorm:"foreignKey:DistrictID" json:"sub_districts"`
	}
	type Provinces struct {
		Name      string      `json:"name"`
		Districts []Districts `gorm:"foreignKey:ProvinceID" json:"districts"`
		ID        uint        `gorm:"primarykey" json:"id"`
	}

	var (
		province []Provinces
		result   gin.H
	)

	strDB.DB.
		// Joins("JOIN districts on districts.province_id=provinces.id").
		// Joins("JOIN sub_districts on sub_districts.district_id=districts.id").
		Preload("Districts").
		Preload("Districts.SubDistricts").
		Find(&province)
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
		province models.Provinces
		result   gin.H
	)
	id := c.Param("id")
	name := c.PostForm("name")
	if err := c.Bind(&province); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).Find(&province)
		province.Name = name

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

//Delete Province
func (strDB *StrDB) Midleware(c *gin.Context) {

	header := c.Request.Header.Get("Authorization")
	fmt.Println(header)

	if header == "token" {
		fmt.Println("token berhasil")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
			"error":   "token Salah",
		})
		c.Abort()
	}

}
