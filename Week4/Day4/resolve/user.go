package resolve

import (
	"golangGraphql/config"
	"golangGraphql/models"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

//GetUser func
func GetUser(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()
	type OfficeLocation struct {
		Name     string `json:"name"`
		ID       uint   `gorm:"primarykey" json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	var user []models.Users
	dbPG.Find(&user)
	return user, nil
}

//RegisterUser func
func RegisterUser(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var user models.Users

	user.ID = uint(rand.Intn(100000))
	user.Email = params.Args["email"].(string)
	user.Password = params.Args["password"].(string)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error -> ", err.Error())
	}
	user.Password = string(hash)
	user.Role = params.Args["role"].(string)

	db.Create(&user)
	return user, nil
}

//LoginUser func
func LoginUser(params graphql.ResolveParams) (interface{}, error) {
	var (
		user   models.Users
		result interface{}
	)

	db := config.Connect()
	email := params.Args["email"].(string)
	password := params.Args["password"].(string)
	db.Where("email = ?", email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
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
			user.Email,
			user.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
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
	return result, nil
}

//comparePassword func
func comparePassword(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

const exp = "24h"

//createToken func
func createToken(user string) (string, error) {

	exp, err := time.ParseDuration(exp)
	if err != nil {

		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(exp).Unix(),
	})
	s, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return "", err
	}
	return s, nil
}
