package resolve

import (
	"golangGraphql/config"
	"golangGraphql/models"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

//GetPerson func
func GetPerson(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	type PersonHadleOfficeLocations struct {
		ID               uint `gorm:"primarykey" json:"id"`
		OfficeLocationID uint `gorm:"foreignkey:OfficelocationID" json:"office_location_id"`
		PersonID         uint `gorm:"foreignkey:PersonID" json:"person_id"`
	}

	type Person struct {
		ID                         uint                         `gorm:"primarykey" json:"id"`
		SubDistrictID              uint                         `gorm:"foreignkey" json:"subdistrict_id"`
		Nip                        string                       `json:"nip"`
		FullName                   string                       `json:"full_name"`
		FirstName                  string                       `json:"first_name"`
		LastName                   string                       `json:"last_name"`
		BirthDate                  string                       `json:"birth_date"`
		BirthPlace                 string                       `json:"birth_place"`
		Gender                     string                       `json:"gender"`
		Photo                      string                       `json:"photo"`
		ZonaLocation               string                       `json:"zona_location"`
		PersonHadleOfficeLocations []PersonHadleOfficeLocations `gorm:"ForeignKey:PersonID" json:"person_handle_office_location"`
	}

	var person []Person
	dbPG.Preload("PersonHadleOfficeLocations").Find(&person)
	return person, nil
}

//CreatePerson func
func CreatePerson(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var person models.Persons

	person.ID = uint(rand.Intn(100000))
	person.SubDistrictID = params.Args["subdistrict_id"].(uint)
	person.Nip = params.Args["nip"].(string)
	person.FullName = params.Args["full_name"].(string)
	person.FirstName = params.Args["first_name"].(string)
	person.LastName = params.Args["last_name"].(string)
	person.BirthDate = params.Args["birth_date"].(string)
	person.BirthPlace = params.Args["birth_place"].(string)
	person.Gender = params.Args["gender"].(string)
	person.Photo = params.Args["photo"].(string)
	person.ZonaLocation = params.Args["zona_location"].(string)

	db.Create(&person)
	return person, nil
}

//UpdatePerson func
func UpdatePerson(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)
	subDistrictID := params.Args["subdistrict_id"].(uint)
	nip := params.Args["nip"].(string)
	fullName, checkName := params.Args["full_name"].(string)
	firstName := params.Args["first_name"].(string)
	lastName := params.Args["last_name"].(string)
	birthDate := params.Args["birth_date"].(string)
	birthPlace := params.Args["birth_place"].(string)
	gender := params.Args["gender"].(string)
	photo := params.Args["photo"].(string)
	zonaLocation := params.Args["zona_location"].(string)

	var person []models.Persons
	for i, v := range person {
		if uint(id) == v.ID {
			if checkName {
				person[i].SubDistrictID = subDistrictID
				person[i].Nip = nip
				person[i].FullName = fullName
				person[i].FirstName = firstName
				person[i].LastName = lastName
				person[i].BirthDate = birthDate
				person[i].BirthPlace = birthPlace
				person[i].Gender = gender
				person[i].Photo = photo
				person[i].ZonaLocation = zonaLocation
			}
			db.Save(&person)
		}
	}
	return person, nil
}

//DeletePerson func
func DeletePerson(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)

	var person models.Persons

	db.Delete(person, id)

	return person, nil
}
