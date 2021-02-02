package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Middleware log func
func Middleware(c *gin.Context) {
	reqMethod := c.Request.Method
	reqPath := c.Request.URL
	reqHttp := c.Request.Host
	log.Println(reqMethod, " -> ", reqHttp, reqPath)
}
