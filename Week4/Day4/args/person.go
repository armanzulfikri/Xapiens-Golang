package args

import "github.com/graphql-go/graphql"

//CreatePersonArgs func
func CreatePersonArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"subdistrict_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"nip": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"full_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"first_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"last_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_place": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"gender": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"photo": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"zona_location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//UpdatePersonArgs func
func UpdatePersonArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"subdistrict_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"nip": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"full_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"first_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"last_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_place": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"gender": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"photo": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"zona_location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//DeletePersonArgs func
func DeletePersonArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
