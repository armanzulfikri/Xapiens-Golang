package midlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

//LogTerminal func
func LogTerminal(c *gin.Context) {
	responseURL := c.Request.URL
	responseHTTP := c.Request.Host
	response := c.Request.Method

	log.Println(response, "-->", responseHTTP, responseURL)
}

// MiddlewareSentryLog func
func MiddlewareSentryLog(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(buf))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	messageSentry := "Post Province \n" + string(buf)
	Sentry(messageSentry)
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
	Sentry(messageSentry)
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

//Authentication func
func Authentication(c *gin.Context) {
	var Admin, Guest string
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env file")
	} else {
		log.Println("Sukses Load .env")
	}

	token := c.Request.Header["Authorized"]
	method := c.Request.Method

	if token[0] == Guest {
		if method != "GET" {
			result := gin.H{
				"message": "Unauthorized Actifity",
			}
			c.JSON(http.StatusUnauthorized, result)
			LogTerminal(c)
			MiddlewareSentryLog(c)
			c.Abort()
		} else {
			result := gin.H{
				"message": "Authorized Actifity",
			}
			c.JSON(http.StatusAccepted, result)
			LogTerminal(c)
			MiddlewareSentryLog(c)
		}
	} else if token[0] == Admin {
		result := gin.H{
			"message": "Authorized Actifity",
		}
		c.JSON(http.StatusAccepted, result)
		LogTerminal(c)
		MiddlewareSentryLog(c)
	}

}
