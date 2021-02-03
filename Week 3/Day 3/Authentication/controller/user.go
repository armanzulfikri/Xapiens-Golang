package controller

import (
	"log"
	"net/http"
	"pendudukApp/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (strDB *StrDB) LoginUser(c *gin.Context) {
	var (
		user   models.Users
		userDB models.Users
		result gin.H
	)

	if err := c.Bind(&user); err != nil {
		log.Println("Data tidak ada, error mesage :", err.Error())
	}

	strDB.DB.Where("email = ?", user.Email).First(&userDB)

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		log.Println("Email", user.Email, "Password Salah")
		result = gin.H{
			"message": "email atau password anda salah",
		}
	} else {
		log.Println("Email", user.Email, "password", user.Password)
		result = gin.H{
			"message": "anda berhasil login",
		}
	}
	c.JSON(http.StatusOK, result)
}
