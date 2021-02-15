package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//Auth func
func Auth(c *gin.Context) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("SECRET TOKEN not found, message", err.Error())
		result := gin.H{
			"message": "kesalahan authentifikasi",
			"token":   nil,
		}
		c.JSON(http.StatusUnauthorized, result)
	}
	secret := os.Getenv("SECRET_TOKEN")
	tokenStringHeader := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenStringHeader, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Method not found or not Hs256", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err.Error())
		result := gin.H{
			"message": "Cannot Access that Operation",
			"eror":    err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	} else if token != nil {
		fmt.Println("token verified")
	}

	dataUser := token.Claims.(jwt.MapClaims)
	method := c.Request.Method

	if dataUser["role"] == "guest" {
		if method != "GET" {
			result := gin.H{
				"message": "Unauthorized Actifity",
			}
			log.Println("User :", dataUser["email"])
			LogTerminal(c)
			MidlewareSentry(c)
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		} else {
			result := gin.H{
				"message": "Authorized Actifity",
				"user":    dataUser["email"],
			}
			log.Println("User :", dataUser["email"])
			LogTerminal(c)
			MidlewareSentry(c)
			c.JSON(http.StatusAccepted, result)
		}
	} else if dataUser["role"] != "admin" {
		if method == "DELETE" {
			result := gin.H{
				"message": "Unauthorized Actifity",
			}
			log.Println("User :", dataUser["email"])
			LogTerminal(c)
			MidlewareSentry(c)
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		} else {
			result := gin.H{
				"message": "Authorized Actifity",
				"user":    dataUser["email"],
			}
			log.Println("User :", dataUser["email"])
			LogTerminal(c)
			MidlewareSentry(c)
			c.JSON(http.StatusAccepted, result)
		}
	} else {
		result := gin.H{
			"message": "Authorized Actifity",
		}
		log.Println("User :", dataUser["email"])
		LogTerminal(c)
		MidlewareSentry(c)
		c.JSON(http.StatusAccepted, result)
	}

}

//MidlewareSentry func
func MidlewareSentry(c *gin.Context) {
	responseURL := c.Request.URL.Path
	responseHTTP := c.Request.Host
	response := c.Request.Method

	buf, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(buf))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	messageSentry := string(response) + "-->" + responseHTTP + responseURL + "\n" + string(buf)
	LogTerminal(c)
	Sentry(messageSentry)
}

//LogTerminal func
func LogTerminal(c *gin.Context) {
	responseHTTP := c.Request.Host
	response := c.Request.Method

	buf, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	log.Println(response, "-->", responseHTTP, "\n", string(buf))
}

//Sentry func
func Sentry(data string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://77a76c1791c74db79298ac0145e974b4@o515017.ingest.sentry.io/5619087",
		Environment: "",
		Release:     "my-project-name@1.0.0",
		Debug:       true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage(data)
}
