package controller

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

//PostCreatePerson func
func (strDB *StrDB) PostCreatePerson(c *gin.Context) {
	var (
		persons models.Persons
		result  gin.H
	)

	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10,
	)
	pool.MaxActive = 10

	conn := pool.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", "report:1"))
	if err == nil {
		if err := c.Bind(&persons); err != nil {
			fmt.Println("No Data or something wrong happen!!!")
		} else {
			file, _ := c.FormFile("photo")
			if file.Size > 800000 && file.Header["Content-Type"][0] != "image/jpg" || file.Header["Content-Type"][0] != "image/png" {
				result := gin.H{
					"Error": "must upload photo under 200kb and Format PNG or JPG",
				}
				c.JSON(http.StatusBadRequest, result)
				return
			} else {
				result := gin.H{
					"Message": "Upload Photo Confirmed",
				}
				c.JSON(http.StatusBadRequest, result)
			}
			path := "images/" + file.Filename

			if err := c.SaveUploadedFile(file, path); err != nil {
				fmt.Println("Terjadi Error ", err.Error())
			}
			persons.Photo = strconv.FormatUint(uint64(persons.ID), 10) + strings.ToLower(file.Filename) + " " + time.Now().String()
			strDB.DB.Create(&persons)
			result = gin.H{
				"message": "success",
				"data": map[string]interface{}{
					"ID":              persons.ID,
					"subdistricts_Id": persons.SubDistrictID,
					"nip":             persons.Nip,
					"full_name":       persons.FullName,
					"first_name":      persons.FirstName,
					"last_name":       persons.LastName,
					"birth_date":      persons.BirthDate,
					"birth_place":     persons.BirthPlace,
					"gender":          persons.Gender,
					"photo":           persons.Photo,
					"zona_location":   persons.ZonaLocation,
					"created_at":      persons.CreatedAt,
					"update_at":       persons.UpdatedAt,
				},
			}
		}
		jm, _ := json.Marshal(&persons)
		_, _ = conn.Do("SET", "report:2", string(jm))
		_, _ = conn.Do("EXPIRE", "report:2", "60")
		fmt.Println(reply)
	}
	c.JSON(http.StatusOK, result)

}

//GetAllPerson func
func (strDB *StrDB) GetAllPerson(c *gin.Context) {
	var (
		persons []models.Persons
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helper.Paging(&helper.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &persons)

	c.JSON(http.StatusOK, paginator)
}

//GetOnePerson func
func (strDB *StrDB) GetOnePerson(c *gin.Context) {
	var (
		persons []models.Persons
		result  gin.H
	)
	id := c.Query("id")
	strDB.DB.First(&persons, id)
	if length := len(persons); length <= 0 {
		result = ResponAPINil(persons, length)
	} else {
		result = ResponAPI(persons, length)
	}
	c.JSON(http.StatusOK, result)
}

//UpdatePerson func
func (strDB *StrDB) UpdatePerson(c *gin.Context) {
	var (
		persons models.Persons
		result  gin.H
	)
	id := c.Param("id")
	nip := c.PostForm("nip")
	fullName := c.PostForm("fullName")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	birthDate := c.PostForm("birthDate")
	birthPlace := c.PostForm("birthPlace")
	gender := c.PostForm("gender")
	zonaLocation := c.PostForm("zonaLocation")

	if err := c.Bind(&persons); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).Find(&persons)
		persons.Nip = nip
		persons.FullName = fullName
		persons.FirstName = firstName
		persons.LastName = lastName
		persons.BirthDate = birthDate
		persons.BirthPlace = birthPlace
		persons.Gender = gender
		persons.ZonaLocation = zonaLocation

		result = gin.H{
			"message": "success Update Data",
		}
		strDB.DB.Save(&persons)
		c.JSON(http.StatusOK, result)
	}
}

//DeletePerson func
func (strDB *StrDB) DeletePerson(c *gin.Context) {
	var (
		persons []models.Persons
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&persons)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

//UpdateProfile func
func (strDB *StrDB) UpdateProfile(c *gin.Context) {
	var (
		persons models.Persons
		result  gin.H
	)

	id := c.Param("id")
	file, _ := c.FormFile("photo")
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10,
	)
	pool.MaxActive = 10

	conn := pool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", "report:3"))
	if err == nil {
		if file.Size > 800000 && file.Header["Content-Type"][0] != "image/jpg" || file.Header["Content-Type"][0] != "image/png" {
			result := gin.H{
				"Error": "must upload photo under 200kb and Format PNG or JPG",
			}
			c.JSON(http.StatusBadRequest, result)
			return
		} else {
			result := gin.H{
				"Message": "Upload Photo Confirmed",
			}
			c.JSON(http.StatusBadRequest, result)
		}
		path := "images/" + file.Filename

		if err := c.SaveUploadedFile(file, path); err != nil {
			fmt.Println("Terjadi Error ", err.Error())
		}

		strDB.DB.Where("id = ?", id).Find(&persons)
		persons.Photo = id + "_" + strings.ToLower(file.Filename) + "_" + time.Now().String()
		result = gin.H{
			"message": "success Uploadfoto",
		}

		strDB.DB.Save(&persons)
		jm, _ := json.Marshal(&persons)
		_, _ = conn.Do("SET", "report:3", string(jm))
		_, _ = conn.Do("EXPIRE", "report:3", "60")
		fmt.Println(reply)
		c.JSON(http.StatusOK, result)
	}
}
