package main


import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"phonebook/controller"
)



func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Welcome to Phone book books apps",
		})
	})
	
	r.GET("/phone",controller.GetListBook)
	r.POST("/phone/create",controller.InputDatas)
	r.PUT("/phone/update",controller.UpdateNumber)

	r.Run()
}