package main

import (
	"adoptme/app/model"
	"adoptme/config"
	"adoptme/routes"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mail()
	config.Connect()
	config.Migrate(&model.User{})

	router := gin.Default()

	routes.Router(router)

	router.Run()
}

func mail() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", "bimapw4@gmail.com")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "<b>have a nice day</b>")

	port, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		port,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
