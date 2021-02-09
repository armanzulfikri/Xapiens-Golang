package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func main() {
	route := gin.Default()
	runCron()
	route.Run()

}
func runCron() {
	gocron.Every(9).Seconds().Do(SendMailInisial)
	gocron.Every(6).Seconds().Do(taskWihoutParams)
	gocron.Every(3).Seconds().Do(taskWithParams, 3, "tiga", true)
	<-gocron.Start()
}
func taskWihoutParams() {
	fmt.Println("task function without params")
}

func taskWithParams(a int, b string, c bool) {
	fmt.Println("Task function with params ", a, "|", b, "|", c, "|")
}

func SendMailInisial() {
	to := []string{"armanuzulfikri@yopmail.com"}
	cc := []string{"armanzulfikri@gmail.com"}

	subject := "Test mail"
	message := "hello ini email uji coba"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Mail Send!")
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
