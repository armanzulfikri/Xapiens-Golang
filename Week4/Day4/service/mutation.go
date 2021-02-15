package service

import (
	"golangGraphql/args"
	"golangGraphql/resolve"
	"golangGraphql/types"

	"github.com/graphql-go/graphql"
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
