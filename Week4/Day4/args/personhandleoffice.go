package args

import "github.com/graphql-go/graphql"

//CreatePersonHandleOfficeArgs func
func CreatePersonHandleOfficeArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"office_location_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"person_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

//UpdatePersonHandleOfficeArgs func
func UpdatePersonHandleOfficeArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"office_location_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"person_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

//DeletePersonHandleOfficeArgs func
func DeletePersonHandleOfficeArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
