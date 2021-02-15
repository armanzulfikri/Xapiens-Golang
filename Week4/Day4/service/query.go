package service

import (
	"golangGraphql/resolve"
	"golangGraphql/types"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"list_province": &graphql.Field{
				Type:        graphql.NewList(types.ProvinceType()),
				Description: "Get Province List",
				Resolve:     resolve.GetProvince,
			},

			"list_district": &graphql.Field{
				Type:        graphql.NewList(types.DistrictType()),
				Description: "Get District List",
				Resolve:     resolve.GetDistrict,
			},

			"get_province_district": &graphql.Field{
				Type:        graphql.NewList(types.ProvinceType()),
				Description: "Get province and district list",
				Resolve:     resolve.GetProvinceDistrict,
			},
			"list_subdistrict": &graphql.Field{
				Type:        graphql.NewList(types.SubDistrictType()),
				Description: "Get SubDistrict List",
				Resolve:     resolve.GetSubDistrict,
			},
			"list_person": &graphql.Field{
				Type:        graphql.NewList(types.PersonType()),
				Description: "Get Person List",
				Resolve:     resolve.GetPerson,
			},
			"list_office_location": &graphql.Field{
				Type:        graphql.NewList(types.OfficeLocationType()),
				Description: "Get Office Location List",
				Resolve:     resolve.GetOfficeLocation,
			},
			"list_person_handle_office": &graphql.Field{
				Type:        graphql.NewList(types.PersonHandleOffice()),
				Description: "Get Person Handle Location List",
				Resolve:     resolve.GetPersonHandleOffice,
			},
		},
	},
)
