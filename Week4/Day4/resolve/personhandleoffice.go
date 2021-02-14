package resolve

import (
	"golangGraphql/config"
	"golangGraphql/models"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

//GetPersonHandleOffice func
func GetPersonHandleOffice(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()
	type PersonHadleOfficeLocations struct {
		ID               uint `gorm:"primarykey" json:"id"`
		OfficeLocationID uint `gorm:"foreignkey:OfficelocationID" json:"office_location_id"`
		PersonID         uint `gorm:"foreignkey:PersonID" json:"person_id"`
	}

	var personHadleOfficeLocations []PersonHadleOfficeLocations
	dbPG.Find(&personHadleOfficeLocations)
	return personHadleOfficeLocations, nil
}

//CreatePersonHandleOffice func
func CreatePersonHandleOffice(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var personHadleOfficeLocations models.PersonHadleOfficeLocations

	personHadleOfficeLocations.ID = uint(rand.Intn(100000))
	personHadleOfficeLocations.PersonID = params.Args["person_id"].(uint)
	personHadleOfficeLocations.OfficeLocationID = params.Args["office_location_id"].(uint)

	db.Create(&personHadleOfficeLocations)
	return personHadleOfficeLocations, nil
}

//UpdatePersonHandleOffice func
func UpdatePersonHandleOffice(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)
	officeLocationID := params.Args["office_location_id"].(uint)
	personID := params.Args["person_id"].(uint)

	var personHadleOfficeLocations []models.PersonHadleOfficeLocations
	for i, v := range personHadleOfficeLocations {
		if uint(id) == v.ID {
			personHadleOfficeLocations[i].PersonID = personID
			personHadleOfficeLocations[i].OfficeLocationID = officeLocationID
			db.Save(&personHadleOfficeLocations)
		}
	}
	return personHadleOfficeLocations, nil
}

//DeletePersonHandleOffice func
func DeletePersonHandleOffice(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id, _ := params.Args["id"].(int)

	var personHadleOfficeLocations models.PersonHadleOfficeLocations

	db.Delete(personHadleOfficeLocations, id)

	return personHadleOfficeLocations, nil
}
