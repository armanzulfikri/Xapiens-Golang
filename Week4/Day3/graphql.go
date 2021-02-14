package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

//Products model
type Products struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Price float64 `json:"price"`
}

//Stores model
type Stores struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Info     string `json:"info"`
	Location string `json:"location"`
}

//Provinces model
type Provinces struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"update_at"`
	DeleteAt  string `json:"delete_at"`
}

var products = []Products{
	{
		ID:    1,
		Name:  "coki coki",
		Info:  "coklat + kacang mete",
		Price: 1000,
	},
	{
		ID:    2,
		Name:  "beng beng",
		Info:  "coklat crunchy",
		Price: 1000,
	},
}

var stores = []Stores{
	{
		ID:       1,
		Name:     "Toko Sembarangan",
		Info:     "Tokoknya mau tutup",
		Location: "yogya",
	},
	{
		ID:       2,
		Name:     "Toko Campur Aduk",
		Info:     "Tokoknya jualan semen doang",
		Location: "yogya",
	},
}




var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
var storeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Store",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get Product List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return products, nil
				},
			},
			"storelist": &graphql.Field{
				Type:        graphql.NewList(storeType),
				Description: "Get Store List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return stores, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createproduct": &graphql.Field{
				Type:        productType,
				Description: "Create new product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},

				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					product := Products{
						ID:    int64(rand.Intn(100000)),
						Name:  params.Args["name"].(string),
						Info:  params.Args["info"].(string),
						Price: params.Args["price"].(float64),
					}
					products = append(products, product)
					return product, nil
				},
			},
			"createstore": &graphql.Field{
				Type:        storeType,
				Description: "Create new store",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},

				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					store := Stores{
						ID:       int64(rand.Intn(100000)),
						Name:     params.Args["name"].(string),
						Info:     params.Args["info"].(string),
						Location: params.Args["location"].(string),
					}
					stores = append(stores, store)
					return store, nil
				},
			},
			"updateproduct": &graphql.Field{
				Type:        productType,
				Description: "Update Product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},

				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, _ := params.Args["id"].(int)
					name, checkName := params.Args["name"].(string)
					info, _ := params.Args["info"].(string)
					price, _ := params.Args["price"].(float64)
					product := Products{}
					for i, v := range products {
						if int64(id) == v.ID {
							if checkName {
								products[i].Name = name
							}
							products[i].Info = info
							products[i].Price = price
							product = products[i]
							break
						}
					}
					return product, nil
				},
			},
			"updatestore": &graphql.Field{
				Type:        storeType,
				Description: "Update Store",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},

				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, _ := params.Args["id"].(int)
					name, checkName := params.Args["name"].(string)
					info, _ := params.Args["info"].(string)
					location, _ := params.Args["location"].(string)
					store := Stores{}
					for i, v := range stores {
						if int64(id) == v.ID {
							if checkName {
								stores[i].Name = name
							}
							stores[i].Info = info
							stores[i].Location = location
							store = stores[i]
							break
						}
					}
					return store, nil
				},
			},
			"deleteproduct": &graphql.Field{
				Type:        productType,
				Description: "Delete Product",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},

				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, _ := params.Args["id"].(int)
					product := Products{}
					for i, v := range products {
						if int64(id) == v.ID {
							product = products[i]
							products = append(products[:i], products[i+1:]...)
						}
					}
					return product, nil
				},
			},
			"deletestore": &graphql.Field{
				Type:        storeType,
				Description: "Delete Store",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},

				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, _ := params.Args["id"].(int)
					store := Stores{}
					for i, v := range stores {
						if int64(id) == v.ID {
							store = stores[i]
							stores = append(stores[:i], stores[i+1:]...)
						}
					}
					return store, nil
				},
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
		// json.NewEncoder(c.Writer).Encode(result)
	})
	r.Run()
}
