package controller

import (
	"CasePoint3/helper"
	"CasePoint3/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
	c.JSON(http.StatusOK, result)
}

//Register func
func (strDB *StrDB) Register(c *gin.Context) {
	var (
		user   models.Users
		result gin.H
	)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error -> ", err.Error())
	}
	if err := c.Bind(&user); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		user.Password = string(hash)
		strDB.DB.Create(&user)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":         user.ID,
				"email":      user.Email,
				"password":   user.Password,
				"role":       user.Role,
				"created_at": user.CreatedAt,
				"update_at":  user.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

//GetAllUser func
func (strDB *StrDB) GetAllUser(c *gin.Context) {
	var (
		users []models.Users
	)

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := helper.Paging(&helper.Param{
		DB:      strDB.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &users)

	c.JSON(http.StatusOK, paginator)
}

//GetOneUser func
func (strDB *StrDB) GetOneUser(c *gin.Context) {
	var (
		users  []models.Users
		result gin.H
	)
	id := c.Query("ID")
	strDB.DB.Find(&users, id)
	if length := len(users); length <= 0 {
		result = ResponAPINil(users, length)
	} else {
		result = ResponAPI(users, length)
	}
	c.JSON(http.StatusOK, result)
}

//UpdateUser func
func (strDB *StrDB) UpdateUser(c *gin.Context) {
	var (
		users  models.Users
		result gin.H
	)
	id := c.Param("id")
	email := c.PostForm("email")
	if err := c.Bind(&users); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Where("id = ?", id).Find(&users)
		users.Email = email

		result = gin.H{
			"message": "success Update Data",
		}
		strDB.DB.Save(&users)
		c.JSON(http.StatusOK, result)
	}
}

//DeleteUsers func
func (strDB *StrDB) DeleteUsers(c *gin.Context) {
	var (
		users []models.Users
	)
	id := c.Param("id")
	d := strDB.DB.Where("id = ?", id).Delete(&users)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
