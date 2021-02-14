package main

import (
	"fmt"
	"golangGraphql/args"
	"golangGraphql/config"
	"golangGraphql/models"
	"golangGraphql/resolve"
	"golangGraphql/seeders"
	"golangGraphql/types"

	"net/http"

	"github.com/gin-gonic/gin"
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
var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//======================== province mutation =======================
			"createprovince": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Create new province",
				Args:        args.CreateProvinceArgs(),

				Resolve: resolve.CreateProvince,
			},
			"updateprovince": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Update Product",
				Args:        args.UpdateProvinceArgs(),

				Resolve: resolve.UpdateProvince,
			},
			"deleteprovince": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Delete Province",
				Args:        args.DeleteProvinceArgs(),
				Resolve:     resolve.DeleteProvince,
			},
			//========================= district mutation ============================
			"createdistrict": &graphql.Field{
				Type:        types.DistrictType(),
				Description: "Create new district",
				Args:        args.CreateDistrictArgs(),

				Resolve: resolve.CreateDistrict,
			},
			"updatedistrict": &graphql.Field{
				Type:        types.DistrictType(),
				Description: "Update District",
				Args:        args.UpdateDistrictsArgs(),

				Resolve: resolve.UpdateDistricts,
			},
			"deletedistrict": &graphql.Field{
				Type:        types.DistrictType(),
				Description: "Delete District",
				Args:        args.DeleteDistrictsArgs(),
				Resolve:     resolve.DeleteDistricts,
			},
			//========================= sub district mutation ====================
			"createsubdistrict": &graphql.Field{
				Type:        types.SubDistrictType(),
				Description: "Create new Subdistrict",
				Args:        args.CreateSubDistrictArgs(),

				Resolve: resolve.CreateSubDistrict,
			},
			"updatesubdistrict": &graphql.Field{
				Type:        types.SubDistrictType(),
				Description: "Update Subdistrict",
				Args:        args.UpdateSubDistrictArgs(),

				Resolve: resolve.UpdateSubDistrict,
			},
			"deletesubdistrict": &graphql.Field{
				Type:        types.SubDistrictType(),
				Description: "Delete Subdistrict",
				Args:        args.DeleteSubDistrictArgs(),
				Resolve:     resolve.DeleteSubDistrict,
			},
			//========================== person mutation ===========================
			"createperson": &graphql.Field{
				Type:        types.PersonType(),
				Description: "Create new Person",
				Args:        args.CreatePersonArgs(),

				Resolve: resolve.CreatePerson,
			},
			"updateperson": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Update Person",
				Args:        args.UpdatePersonArgs(),

				Resolve: resolve.CreatePerson,
			},
			"deleteperson": &graphql.Field{
				Type:        types.PersonType(),
				Description: "Delete Person",
				Args:        args.DeletePersonArgs(),
				Resolve:     resolve.DeletePerson,
			},
			//======================== office location mutation =================
			"createofficelocation": &graphql.Field{
				Type:        types.OfficeLocationType(),
				Description: "Create new Office Location",
				Args:        args.CreateOfficeLocation(),

				Resolve: resolve.CreateOfficeLocation,
			},
			"updateofficelocation": &graphql.Field{
				Type:        types.OfficeLocationType(),
				Description: "Update Office Location",
				Args:        args.UpdateOfficeLocationArgs(),

				Resolve: resolve.UpdateOfficeLocation,
			},
			"deleteofficelocation": &graphql.Field{
				Type:        types.OfficeLocationType(),
				Description: "Delete Office Location",
				Args:        args.DeleteOfficeLocationArgs(),
				Resolve:     resolve.DeleteOfficeLocation,
			},
			//======================== person handle office mutation =================
			"createpersonhandleoffice": &graphql.Field{
				Type:        types.PersonHandleOffice(),
				Description: "Create new Person Handle Office",
				Args:        args.CreatePersonHandleOfficeArgs(),

				Resolve: resolve.CreatePersonHandleOffice,
			},
			"updatepersonhandleoffice": &graphql.Field{
				Type:        types.PersonHandleOffice(),
				Description: "Update Person Handle Office",
				Args:        args.UpdateOfficeLocationArgs(),

				Resolve: resolve.UpdateOfficeLocation,
			},
			"deletepersonhandleoffice": &graphql.Field{
				Type:        types.PersonHandleOffice(),
				Description: "Delete Person Handle Office",
				Args:        args.DeletePersonHandleOfficeArgs(),
				Resolve:     resolve.DeletePersonHandleOffice,
			},
			//================== User mutation ===========================
			"register": &graphql.Field{
				Type:        types.User(),
				Description: "Register User",
				Args:        args.CreateUserArgs(),

				Resolve: resolve.RegisterUser,
			},
			"login": &graphql.Field{
				Type:        types.User(),
				Description: "Login User",
				Args:        args.LoginUserArgs(),

				Resolve: resolve.LoginUser,
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Println(result.Errors)
	}
	return result
}

//main func
func main() {
	pgDB := config.Connect()

	// migration
	models.Migrations(pgDB)

	//Seeding From API Raja Ongkir
	seeders.SeedProvince(pgDB)
	seeders.SeedDisctrict(pgDB)
	seeders.SeedSubDistrict(pgDB)
	seeders.SeedPerson(pgDB)
	seeders.SeedOffice(pgDB)
	seeders.SeedOfficePersonLocation(pgDB)

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}
		var query Query

		c.Bind(&query)

		result := executeQuery(query.Query, schema)
		fmt.Println(result)
		c.JSON(http.StatusAccepted, result)
	})
	r.Run()
}
