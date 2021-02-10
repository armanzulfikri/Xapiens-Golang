package midlewares

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

//Corn func
func Corn(c *gin.Context) {
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

	if dataUser["role"] == "entry" {
		if method != "DELETE" {
			result := gin.H{
				"message": "Authorized Actifity",
			}
			go runCron()
			log.Println("User :", dataUser["email"])
			c.JSON(http.StatusAccepted, result)
		} else {
			result := gin.H{
				"message": "Unauthorized Actifity",
				"user":    dataUser["email"],
			}
			log.Println("User :", dataUser["email"])
			go runCron()
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		}
	}

}

//runCron func
func runCron() {
	gocron.Every(10).Minute().Do(SendMailInisial)
	<-gocron.Start()
}

func sendMail(to []string, cc []string, subject, message string) error {
	err := godotenv.Load(".env")
	if err == nil {
		body := "From :" + os.Getenv("MAIL_EMAIL") + "\n" +
			"To :" + strings.Join(to, ",") + "\n" +
			"Cc :" + strings.Join(cc, ",") + "\n" +
			"Subject : " + subject + "\n\n" +
			message

		auth := smtp.PlainAuth("", os.Getenv("MAIL_EMAIL"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_SMTP_HOST"))
		smtpAddr := os.Getenv("MAIL_SMTP_HOST") + ":" + os.Getenv("MAIL_SMTP_PORT")
		err := smtp.SendMail(smtpAddr, auth, os.Getenv("MAIL_EMAIL"), append(to, cc...), []byte(body))
		if err == nil {
			return nil
		}
		return err
	} else {
		return err
	}
}

//SendMailInisial func
func SendMailInisial() {

	to := []string{"adminadmin@yopmail.com"}
	cc := []string{"admin@gmail.com"}

	subject := "Entry user operation"

	message := "Entry user Has been"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Mail Send!")
}
