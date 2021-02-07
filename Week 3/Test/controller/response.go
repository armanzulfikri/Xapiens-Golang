package controller

import "github.com/gin-gonic/gin"

//ResponAPI func
func ResponAPI(str interface{}, length int) gin.H {
	return gin.H{
		"message": "Sukses",
		"data":    str,
		"count":   length,
	}
}

//ResponAPINil func
func ResponAPINil(str interface{}, length int) gin.H {
	return gin.H{
		"message": "Sukses",
		"data":    nil,
		"count":   length,
	}
}
