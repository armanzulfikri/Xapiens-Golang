package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	r.GET("/ping/:id", func(c *gin.Context){
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message":"pong"+ id,
		})
	})

	r.GET("/ping/", func(c *gin.Context){
		id := c.Query("id")
		role := c.DefaultQuery("role","0")
		c.JSON(200, gin.H{
			"message":"pong",
			"id" : id,
			"role" : role,
		})
	})

	r.POST("/user/", func(c *gin.Context){
		name := c.PostForm("name")
		role := c.PostForm("role")
		c.JSON(200, gin.H{
			"message":"pong",
			"name" : name,
			"role" : role,
			"data" : map[string]string{}{
				"phone":"4662",
				"city": "jogja",
				"pengalaman" : []string{"guru","engineer","qa"}
			},
			"pengalaman": map[string][]string{
				"pengalaman" : {"guru","enginer","qa"},
				
			}package main

			import (
				"github.com/gin-gonic/gin"
				"fmt"
			)
			
			func main() {
			
				r := gin.Default()
				r.GET("/ping", func(c *gin.Context){
					c.JSON(200, gin.H{
						"message":"pong",
					})
				})
			
				r.GET("/ping/:id", func(c *gin.Context){
					id := c.Param("id")
					c.JSON(200, gin.H{
						"message":"pong"+ id,
					})
				})
			
				r.GET("/ping/", func(c *gin.Context){
					id := c.Query("id")
					role := c.DefaultQuery("role","0")
					c.JSON(200, gin.H{
						"message":"pong",
						"id" : id,
						"role" : role,
					})
				})
			
				r.POST("/user/", func(c *gin.Context){
					name := c.PostForm("name")
					role := c.PostForm("role")
					c.JSON(200, gin.H{
						"message":"pong",
						"name" : name,
						"role" : role,
						"data" : map[string]string{}{
							"phone":"4662",
							"city": "jogja",
							"pengalaman" : []string{"guru","engineer","qa"}
						},
						"pengalaman": map[string][]string{
							"pengalaman" : {"guru","enginer","qa"},
							
						}
				
				type user struct{
					Username string `json:"username"`
					Password string `json:"password"`
			
				}
				r.POST("/user", func(c *gin.Context){
					var datauser []userSimpan
					var user user
					var currentUser userSimpan
					err := c.Bind(&user)
					if err != nil {
						fmt.Println("telah terjadi error")
					}
					currentUser.Username = user.Username
					currentUser.Password = user.Password
					datauser = append(datauser,currentUser)
					fmt.Println(datauser)
					c.JSON(200, gin.H{
						"message":"pong",
						"username": user.Username,
						"password" : user.Password,
			
					})
				})
			
				v1 := r.Group ("/v1")
				{
					v1.GET("/ping", func(c *gin.Context){
					c.JSON(200, gin.H{		
						"message": "pong",
					 })
					})
				}
			
				v2 := r.Group ("/v2")
				{
					 }
					 v2.GET("/ping", func(c *gin.Context){
					c.JSON(200, gin.H{
						"message": "pong",
					})
					
				})
				
				r.Run()
			}
			package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	r.GET("/ping/:id", func(c *gin.Context){
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message":"pong"+ id,
		})
	})

	r.GET("/ping/", func(c *gin.Context){
		id := c.Query("id")
		role := c.DefaultQuery("role","0")
		c.JSON(200, gin.H{
			"message":"pong",
			"id" : id,
			"role" : role,
		})
	})

	r.POST("/user/", func(c *gin.Context){
		name := c.PostForm("name")
		role := c.PostForm("role")
		c.JSON(200, gin.H{
			"message":"pong",
			"name" : name,
			"role" : role,
			"data" : map[string]string{}{
				"phone":"4662",
				"city": "jogja",
				"pengalaman" : []string{"guru","engineer","qa"}
			},
			"pengalaman": map[string][]string{
				"pengalaman" : {"guru","enginer","qa"},
				
			}
	
	type user struct{
		Username string `json:"username"`
		Password string `json:"password"`

	}
	r.POST("/user", func(c *gin.Context){
		var datauser []userSimpan
		var user user
		var currentUser userSimpan
		err := c.Bind(&user)
		if err != nil {
			fmt.Println("telah terjadi error")
		}
		currentUser.Username = user.Username
		currentUser.Password = user.Password
		datauser = append(datauser,currentUser)
		fmt.Println(datauser)
		c.JSON(200, gin.H{
			"message":"pong",
			"username": user.Username,
			"password" : user.Password,

		})
	})

	v1 := r.Group ("/v1")
	{
		v1.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{		
			"message": "pong",
		 })
		})
	}

	v2 := r.Group ("/v2")
	{
		 }
		 v2.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
		
	})
	
	r.Run()
}
package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	r.GET("/ping/:id", func(c *gin.Context){
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message":"pong"+ id,
		})
	})

	r.GET("/ping/", func(c *gin.Context){
		id := c.Query("id")
		role := c.DefaultQuery("role","0")
		c.JSON(200, gin.H{
			"message":"pong",
			"id" : id,
			"role" : role,
		})
	})

	r.POST("/user/", func(c *gin.Context){
		name := c.PostForm("name")
		role := c.PostForm("role")
		c.JSON(200, gin.H{
			"message":"pong",
			"name" : name,
			"role" : role,
			"data" : map[string]string{}{
				"phone":"4662",
				"city": "jogja",
				"pengalaman" : []string{"guru","engineer","qa"}
			},
			"pengalaman": map[string][]string{
				"pengalaman" : {"guru","enginer","qa"},
				
			}
	
	type user struct{
		Username string `json:"username"`
		Password string `json:"password"`

	}
	r.POST("/user", func(c *gin.Context){
		var datauser []userSimpan
		var user user
		var currentUser userSimpan
		err := c.Bind(&user)
		if err != nil {
			fmt.Println("telah terjadi error")
		}
		currentUser.Username = user.Username
		currentUser.Password = user.Password
		datauser = append(datauser,currentUser)
		fmt.Println(datauser)
		c.JSON(200, gin.H{
			"message":"pong",
			"username": user.Username,
			"password" : user.Password,

		})
	})

	v1 := r.Group ("/v1")
	{
		v1.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{		
			"message": "pong",
		 })
		})
	}

	v2 := r.Group ("/v2")
	{
		 }
		 v2.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
		
	})
	
	r.Run()
}
package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	r.GET("/ping/:id", func(c *gin.Context){
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message":"pong"+ id,
		})
	})

	r.GET("/ping/", func(c *gin.Context){
		id := c.Query("id")
		role := c.DefaultQuery("role","0")
		c.JSON(200, gin.H{
			"message":"pong",
			"id" : id,
			"role" : role,
		})
	})

	r.POST("/user/", func(c *gin.Context){
		name := c.PostForm("name")
		role := c.PostForm("role")
		c.JSON(200, gin.H{
			"message":"pong",
			"name" : name,
			"role" : role,
			"data" : map[string]string{}{
				"phone":"4662",
				"city": "jogja",
				"pengalaman" : []string{"guru","engineer","qa"}
			},
			"pengalaman": map[string][]string{
				"pengalaman" : {"guru","enginer","qa"},
				
			}
	
	type user struct{
		Username string `json:"username"`
		Password string `json:"password"`

	}
	r.POST("/user", func(c *gin.Context){
		var datauser []userSimpan
		var user user
		var currentUser userSimpan
		err := c.Bind(&user)
		if err != nil {
			fmt.Println("telah terjadi error")
		}
		currentUser.Username = user.Username
		currentUser.Password = user.Password
		datauser = append(datauser,currentUser)
		fmt.Println(datauser)
		c.JSON(200, gin.H{
			"message":"pong",
			"username": user.Username,
			"password" : user.Password,

		})
	})

	v1 := r.Group ("/v1")
	{
		v1.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{		
			"message": "pong",
		 })
		})
	}

	v2 := r.Group ("/v2")
	{
		 }
		 v2.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
		
	})
	
	r.Run()
}

func (c *gin.Context) {

		c.JSON(200, gin H{
			"message" : "pong"
		})

}
func (c *gin.Context) {

		c.JSON(200, gin H{
			"message" : "pong"
		})

}
func (c *gin.Context) {

		c.JSON(200, gin H{
			"message" : "pong"
		})

}
			func (c *gin.Context) {
			
					c.JSON(200, gin H{
						"message" : "pong"
					})
			
			}
	
	type user struct{
		Username string `json:"username"`
		Password string `json:"password"`

	}
	r.POST("/user", func(c *gin.Context){
		var datauser []userSimpan
		var user user
		var currentUser userSimpan
		err := c.Bind(&user)
		if err != nil {
			fmt.Println("telah terjadi error")
		}
		currentUser.Username = user.Username
		currentUser.Password = user.Password
		datauser = append(datauser,currentUser)
		fmt.Println(datauser)
		c.JSON(200, gin.H{
			"message":"pong",
			"username": user.Username,
			"password" : user.Password,

		})
	})

	v1 := r.Group ("/v1")
	{
		v1.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{		
			"message": "pong",
		 })
		})
	}

	v2 := r.Group ("/v2")
	{
		 }
		 v2.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
		
	})
	
	r.Run()
}

func (c *gin.Context) {

		c.JSON(200, gin H{
			"message" : "pong"
		})

}