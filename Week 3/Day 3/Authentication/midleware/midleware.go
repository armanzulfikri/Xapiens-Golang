package midleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

//LogTerminal
func LogTerminal(c *gin.Context) {
	responseURL := c.Request.URL
	responseHTTP := c.Request.Host
	response := c.Request.Method

	log.Println(response, "-->", responseHTTP, responseURL)
}

// MiddlewareSentryLog
func MiddlewareSentryLog(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(buf))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	messageSentry := "Post Province \n" + string(buf)
	Sentry(messageSentry)
}

//SentryMonitoring
func SentryMonitoring() {
	err := sentry.Init(sentry.ClientOptions{
		TracesSampleRate: 0.2,
		TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			return sentry.SampledTrue
		}),
	})
	if err != nil {
		log.Fatalf("sentry.Init : %s", err)
	}
}

//MidlewareSentry
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

//MidlewareSatu
func MidlewareSatu(c *gin.Context) {
	fmt.Println("ini midleware satu")
	c.Next()
	fmt.Println("ini kembali ke midleware satu")
}

//MidlewareDua
func MidlewareDua(c *gin.Context) {
	fmt.Println("ini midleware dua")
	c.Abort()
}

//Sentry
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

//Authorized
func Authorized(c *gin.Context) {
	var Admin, Guest string
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env file")
	} else {
		Admin = os.Getenv("TOKEN_ADMIN")
		Guest = os.Getenv("TOKEN_GUEST")
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
