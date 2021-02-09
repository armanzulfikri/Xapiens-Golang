package controller

import (
	"CasePoint3/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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

//Case2 func
func (strDB *StrDB) Case2(c *gin.Context) {
	type response struct {
		ID       uint     `json:"id"`
		FullName string   `json:"full_name"`
		Total    int64    `json:"total"`
		CityName []string `json:"city_name"`
	}
	var (
		resp     []response
		id       uint
		FullName string
		CityName string
		total    int64
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
		rows, _ := strDB.DB.Table("persons").
			Select("persons.id as id, persons.full_name as FullName, count(*) as Total").
			Joins("join person_hadle_office_locations opl on opl.person_id = persons.id").
			Joins("join offices_locations ofc on opl.office_location_id = ofc.id").
			Joins("join sub_districts sd on sd.id = ofc.sub_district_id").
			Joins("join districts d on d.id = sd.district_id").
			Group("persons.id, persons.full_name").
			Rows()
		for rows.Next() {
			rows.Scan(&id, &FullName, &total)
			// fmt.Println(id, " | ", FullName, " | ", total)
			var location []string
			innerRows, _ := strDB.DB.Table("persons").
				Select("d.name as CityName").
				Where("persons.id = ?", id).
				Joins("join person_hadle_office_locations opl on opl.person_id = persons.id").
				Joins("join offices_locations ofc on opl.office_location_id = ofc.id").
				Joins(" join sub_districts sd on sd.id = ofc.sub_district_id").
				Joins("join districts d on d.id = sd.district_id").
				Rows()
			for innerRows.Next() {
				innerRows.Scan(&CityName)
				location = append(location, CityName)
			}
			resp = append(resp, response{id, FullName, total, location})
			jm, _ := json.Marshal(resp)
			_, _ = conn.Do("SET", "report:1", string(jm))
			_, _ = conn.Do("EXPIRE", "report:1", "60")
		}

	}

	c.JSON(http.StatusAccepted, string(reply))
}
