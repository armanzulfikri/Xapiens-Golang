package controller

import "github.com/gin-gonic/gin"

//respon API
func ResponAPI(str interface{}, length int) gin.H {
	return gin.H{
		"message": "Sukses",
		"data":    str,
		"count":   length,
	}
}

//responseAPI Nill
func ResponAPINil(str interface{}, length int) gin.H {
	return gin.H{
		"message": "Sukses",
		"data":    nil,
		"count":   length,
	}
}
