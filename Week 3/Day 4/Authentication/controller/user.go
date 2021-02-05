package controller

import (
	"log"
	"net/http"
	"os"
	"pendudukApp/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

//LoginUser samples
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
		type authCustomClaims struct {
			Email string `json:"email"`
			Role  string `json:"role"`
			jwt.StandardClaims
		}
		Claims := &authCustomClaims{
			userDB.Email,
			userDB.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
		if err := godotenv.Load(".env"); err != nil {
			log.Println("SECRET TOKEN not found, message", err.Error())
			result = gin.H{
				"message": "gagal create token",
				"token":   nil,
			}

		} else {
			sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), Claims)

			if token, err := sign.SignedString([]byte(os.Getenv("SECRET_TOKEN"))); err != nil {
				log.Println("Gagal Create Token, message", err.Error())
				result = gin.H{
					"message": "anda tidak berhasil login",
					"token":   nil,
				}
			} else {
				log.Println("Email", user.Email, "telah login")
				result = gin.H{
					"message": "anda berhasil login",
					"token":   token,
				}
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
