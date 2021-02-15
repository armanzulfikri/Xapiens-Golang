package types

import (
	"github.com/graphql-go/graphql"
)

//ProvinceType func
func ProvinceType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//DistrictType func
func DistrictType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"province_id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}

//SubDistrictType func
func SubDistrictType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "SubDistrict",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"district_id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}

//PersonType func
func PersonType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Person",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"subdistrict_id": &graphql.Field{
					Type: graphql.Int,
				},
				"nip": &graphql.Field{
					Type: graphql.String,
				},
				"full_name": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
				"birth_date": &graphql.Field{
					Type: graphql.String,
				},
				"birth_place": &graphql.Field{
					Type: graphql.String,
				},
				"gender": &graphql.Field{
					Type: graphql.String,
				},
				"photo": &graphql.Field{
					Type: graphql.String,
				},
				"zona_location": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//OfficeLocationType func
func OfficeLocationType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "OfficeLocation",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"subdistrict_id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}

//PersonHandleOffice func
func PersonHandleOffice() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PersonHandleOffice",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"office_location_id": &graphql.Field{
					Type: graphql.Int,
				},
				"person_id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)

}

//User func
func User() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"role": &graphql.Field{
					Type: graphql.String,
				},
				"token": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
