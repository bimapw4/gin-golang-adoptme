package main

import (
	"adoptme/app/model"
	"adoptme/config"
	"adoptme/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Connect()
	config.Migrate(&model.User{})

	router := gin.Default()

	routes.Router(router)

	router.Run()
}
