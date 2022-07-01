package routes

import (
	"adoptme/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/register", controller.RegisterUser)
	r.POST("/login", controller.LoginUser)
}
